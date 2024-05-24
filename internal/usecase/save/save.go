package save

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/adminsemy/kpi-drive-test/internal/entity"
	"github.com/adminsemy/kpi-drive-test/internal/myErrors"
)

type Buffer interface {
	Add(item entity.Data) error
	Get() (entity.Data, error)
}

type Client interface {
	Save(item entity.Data) error
}

type Save struct {
	ctx    context.Context
	buffer Buffer
	chSave chan entity.Data
	client Client
}

func New(ctx context.Context, buffer Buffer, client Client) *Save {
	s := &Save{
		ctx:    ctx,
		buffer: buffer,
		chSave: make(chan entity.Data),
		client: client,
	}
	s.run()

	return s
}

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
	}
}

// Запускаем цикл для извлечения данных из буфера
// и сохранения их в нужное место. Если данных нет
// то ждем секунду и повторяем цикл
func (s *Save) run() {
	go func() {
		for {
			select {
			case <-s.ctx.Done():
				return
			default:
				item, err := s.buffer.Get()
				// Если ошибка пустой буфер, то ждем секунду и повторяем цикл
				if errors.Is(myErrors.ErrMaxSize, err) {
					time.Sleep(time.Second)
					continue
				}
				// Если что-то другое - печатаем и выходим
				if err != nil {
					slog.Error("can't get data from buffer", err)
					return
				}
				err = s.client.Save(item)
				if err != nil {
					slog.Error("can't save data", err)
					return
				}
			}
		}
	}()
}
