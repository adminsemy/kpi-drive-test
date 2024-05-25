package cli

import "flag"

const (
	BufferSize = "buf"
	RespCount  = "resp"
)

// Структура для параметров командной строки
type Cli struct {
	CommandsInt map[string]int64
}

func New() *Cli {
	return &Cli{
		CommandsInt: make(map[string]int64),
	}
}

// Заполняем параметры из командной строки
func (c *Cli) Run() {
	// Размер буфера
	bufSize := flag.Int64("buf", 1000, "buffer's size")
	// Количество запросов к API
	respCount := flag.Int64("resp", 10, "response's count")

	flag.Parse()

	c.CommandsInt[BufferSize] = *bufSize
	c.CommandsInt[RespCount] = *respCount
}
