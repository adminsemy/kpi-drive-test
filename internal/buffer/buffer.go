package buffer

import (
	"container/list"
	"sync"

	"github.com/adminsemy/kpi-drive-test/internal/entity"
	"github.com/adminsemy/kpi-drive-test/internal/myErrors"
)

// Буфер для хранения данных с максимальным размером.
// Содержит в себе размер данных maxSize и список с данными list.
type Buffer struct {
	maxSize int64
	list    *list.List
	sync.Mutex
}

// maxSize - максимальный размер буфера
func New(maxSize int64) *Buffer {
	return &Buffer{
		maxSize: maxSize,
		list:    list.New(),
	}
}

// Добавляем данные в буфер
func (b *Buffer) Add(item entity.Data) error {
	b.Lock()
	defer b.Unlock()
	if int64(b.list.Len()) == b.maxSize {
		return myErrors.ErrMaxSize
	}
	b.list.PushBack(item)
	return nil
}

// Извлекаем данные из буфера
func (b *Buffer) Get() (entity.Data, error) {
	b.Lock()
	defer b.Unlock()
	if b.list.Len() == 0 {
		return entity.Data{}, myErrors.ErrEmptyBuffer
	}
	return b.list.Remove(b.list.Front()).(entity.Data), nil
}
