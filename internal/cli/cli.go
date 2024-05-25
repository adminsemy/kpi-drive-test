package cli

import "flag"

const (
	BufferSize = "buf"
	RespCount  = "resp"
)

type Cli struct {
	CommandsInt map[string]int64
}

func New() *Cli {
	return &Cli{
		CommandsInt: make(map[string]int64),
	}
}

func (c *Cli) Run() {
	bufSize := flag.Int64("buf", 1000, "buffer's size")
	respCount := flag.Int64("resp", 10, "response's count")

	flag.Parse()

	c.CommandsInt[BufferSize] = *bufSize
	c.CommandsInt[RespCount] = *respCount
}
