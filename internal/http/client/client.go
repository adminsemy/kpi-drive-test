package client

import (
	"bytes"
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/adminsemy/kpi-drive-test/internal/entity"
)

// URL для запроса сохранения данных
const saveDataPostForm = "https://development.kpi-drive.ru/_api/facts/save_fact"

// Клиент для http запросов к API
type Client struct {
	ctx    context.Context
	token  string
	client *http.Client
}

// ctx - контекст для отмены
// token - токен для авторизации
func New(ctx context.Context, token string) *Client {
	return &Client{
		ctx:    ctx,
		token:  token,
		client: &http.Client{},
	}
}

// Сохраняем данные по пути saveDataPostForm
// Если не удалось - возвращаем ошибку
func (c *Client) Save(item entity.Data) error {
	values := getValues(item)
	request, err := http.NewRequestWithContext(c.ctx, http.MethodPost, saveDataPostForm, bytes.NewBufferString(values.Encode()))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Authorization", "Bearer "+c.token)
	response, err := c.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return nil
}

// Преобразуем данные в url.Values
// для передачи в запрос
func getValues(item entity.Data) url.Values {
	plan := "0"
	if item.IsPlan {
		plan = "1"
	}
	values := make(url.Values)
	values.Add("period_start", item.PeriodStart.Format(time.DateOnly))
	values.Add("period_end", item.PeriodEnd.Format(time.DateOnly))
	values.Add("period_key", item.PeriodKey)
	values.Add("indicator_to_mo_id", strconv.Itoa(int(item.IndicatorToMoId)))
	values.Add("indicator_to_mo_fact", strconv.Itoa(int(item.IndicatorToMoFactId)))
	values.Add("value", strconv.Itoa(int(item.Value)))
	values.Add("fact_time", item.FactTime.Format(time.DateOnly))
	values.Add("is_plan", plan)
	values.Add("auth_user_id", strconv.Itoa(int(item.AuthUserId)))

	return values
}
