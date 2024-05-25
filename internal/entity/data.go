package entity

import "time"

// Структура для эмуляции формы данных
type Data struct {
	PeriodStart         time.Time
	PeriodEnd           time.Time
	PeriodKey           string
	IndicatorToMoId     int
	IndicatorToMoFactId int
	Value               float64
	FactTime            time.Time
	IsPlan              bool
	AuthUserId          int64
	Comment             string
}
