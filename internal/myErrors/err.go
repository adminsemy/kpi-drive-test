package myErrors

import "fmt"

var (
	ErrMaxSize     = fmt.Errorf("buffer is full")
	ErrEmptyBuffer = fmt.Errorf("buffer is empty")
)
