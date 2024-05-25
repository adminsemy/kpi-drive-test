package myErrors

import "fmt"

// Описание всех пользовательских ошибок
var (
	ErrMaxSize     = fmt.Errorf("buffer is full")
	ErrEmptyBuffer = fmt.Errorf("buffer is empty")
)
