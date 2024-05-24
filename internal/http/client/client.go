package client

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/adminsemy/kpi-drive-test/internal/entity"
)

const saveDataPostForm = "https://development.kpi-drive.ru/_api/facts/save_fact"

type Client struct {
	client *http.Client
}

func New(client *http.Client) *Client {
	return &Client{
		client: client,
	}
}

func (c *Client) Save(item entity.Data) error {
	values := getValues(item)
	response, err := c.client.PostForm(saveDataPostForm, values)
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}

func getValues(item entity.Data) url.Values {
	plan := "0"
	if item.IsPlan {
		plan = "1"
	}
	values := url.Values{
		"period_start":         {item.PeriodStart.Format(time.DateOnly)},
		"period_end":           {item.PeriodEnd.Format(time.DateOnly)},
		"period_key":           {item.PeriodKey},
		"indicator_to_mo_id":   {strconv.Itoa(int(item.IndicatorToMoId))},
		"indicator_to_mo_fact": {strconv.Itoa(int(item.IndicatorToMoFactId))},
		"value":                {strconv.Itoa(int(item.Value))},
		"fact_time":            {item.FactTime.Format(time.DateOnly)},
		"is_plan":              {plan},
		"auth_user_id":         {strconv.Itoa(int(item.AuthUserId))},
	}

	return values
}
