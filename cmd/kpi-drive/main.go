package main

import (
	"context"
	"fmt"
	"time"

	"github.com/adminsemy/kpi-drive-test/internal/buffer"
	"github.com/adminsemy/kpi-drive-test/internal/cli"
	"github.com/adminsemy/kpi-drive-test/internal/entity"
	"github.com/adminsemy/kpi-drive-test/internal/http/client"
	"github.com/adminsemy/kpi-drive-test/internal/usecase/save"
)

// Запуск программы
func main() {
	c := cli.New()
	c.Run()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	buffer := buffer.New(c.CommandsInt[cli.BufferSize])
	client := client.New("48ab34464a5573519725deb5865cc74c")
	chDone := make(chan struct{})
	save := save.New(ctx, buffer, client, chDone)
	start, _ := time.Parse(time.DateOnly, "2024-05-01")
	end, _ := time.Parse(time.DateOnly, "2024-05-31")
	factTime, _ := time.Parse(time.DateOnly, "2024-05-31")
	ent := entity.Data{
		PeriodStart:         start,
		PeriodEnd:           end,
		PeriodKey:           "month",
		IndicatorToMoId:     227373,
		IndicatorToMoFactId: 0,
		FactTime:            factTime,
		IsPlan:              false,
		AuthUserId:          40,
		Comment:             "buffer Last_name",
	}
	for i := int64(0); i < c.CommandsInt[cli.RespCount]; i++ {
		ent.Value = float64(i)
		save.Add(ent)
	}
	close(chDone)
	<-save.Done()
	fmt.Println("done")
}
