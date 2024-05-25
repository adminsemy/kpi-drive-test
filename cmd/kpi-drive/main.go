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
	// Считываем параметры командной строки
	c := cli.New()
	c.Run()
	// Создаем контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Создаем буфер для хранения данных,
	// которые нужно сохранить
	buffer := buffer.New(c.CommandsInt[cli.BufferSize])
	// Создаем клиент для http запросов
	client := client.New(ctx, "48ab34464a5573519725deb5865cc74c")
	// Канал для сигнала о том, что запросов больше нет
	// Своего рода заглушка, что бы программа не завершилась раньше
	chDone := make(chan struct{})
	// Юзкейс для сохранения данных (описывается логика сохранения)
	save := save.New(ctx, buffer, client, chDone)

	// Заполняем данные для запроса
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

	// Запускаем цикл для сохранения данных
	for i := int64(0); i < c.CommandsInt[cli.RespCount]; i++ {
		ent.Value = float64(i)
		save.Add(ent)
	}
	// Сообщаем, что даных больше нет
	close(chDone)
	// Ждем сигнала о завершении сохранения
	<-save.Done()
	fmt.Println("done")
}
