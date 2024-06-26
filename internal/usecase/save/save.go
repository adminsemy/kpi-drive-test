package save

import (
	"context"
	"errors"
	"log/slog"

	"github.com/adminsemy/kpi-drive-test/internal/entity"
	"github.com/adminsemy/kpi-drive-test/internal/myErrors"
)

// Интерфейс для буфера данных
type Buffer interface {
	Add(item entity.Data) error
	Get() (entity.Data, error)
}

// Интерфейс для клиента
type Client interface {
	Save(item entity.Data) error
}

// Процесс сохранения данных в нужное место
// Считывает данные из буфера и сохраняет их в нужное место
type Save struct {
	ctx    context.Context
	buffer Buffer
	chDone chan struct{}
	chExit chan struct{}
	client Client
}

// ctx - контекст для отмены
// buffer - буфер для хранения данных
// client - клиент для http запросов
// chDone - канал для сигнала о том,что данных больше нет
func New(ctx context.Context, buffer Buffer, client Client, chDone chan struct{}) *Save {
	s := &Save{
		ctx:    ctx,
		buffer: buffer,
		chDone: chDone,
		chExit: make(chan struct{}),
		client: client,
	}
	s.run()

	return s
}

// Добавляем данные в буфер
// Если не удалось - возвращаем ошибку
// Если буфер заполнен - остальные данные будут утеряны
func (s *Save) Add(entity entity.Data) {
	select {
	case <-s.ctx.Done():
		return
	default:
		err := s.buffer.Add(entity)
		if err != nil {
			slog.Error("can't add data to buffer", err)
			return
		}
		slog.Info("add data to buffer", "value:", entity.Value)
	}
}

// Сигнал о том,что все сохранения завершены
func (s *Save) Done() <-chan struct{} {
	return s.chExit
}

// Запускаем цикл для извлечения данных из буфера
// и сохранения их в нужное место. Если данных нет
// то повторяем цикл
func (s *Save) run() {
	go func() {
		for {
			select {
			case <-s.ctx.Done():
				return
			default:
				item, err := s.buffer.Get()
				// Если ошибка пустой буфер, то ждем секунду и повторяем цикл
				if errors.Is(myErrors.ErrEmptyBuffer, err) {
					select {
					case <-s.ctx.Done():
						return
					case <-s.chDone:
						close(s.chExit)
						slog.Info("exit save")
						return
					default:
						continue
					}
				}
				// Если что-то другое - печатаем и выходим
				if err != nil {
					slog.Error("can't get data from buffer", err)
					return
				}
				slog.Info("get data from buffer", "value:", item.Value)
				err = s.client.Save(item)
				if err != nil {
					slog.Error("can't save data", err)
					return
				}
				slog.Info("save data", "value:", item.Value)
			}
		}
	}()
}
