package main

import (
	"fmt"

	"github.com/adminsemy/kpi-drive-test/internal/cli"
)

// Запуск программы
func main() {
	c := cli.New()
	c.Run()
	fmt.Println(c.CommandsInt[cli.BufferSize])
}
