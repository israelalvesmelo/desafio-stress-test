package infra

import (
	"net/http"

	"github.com/israelalvesmelo/desafio-stress-test/internal/domain/entity"
)

type RequestGateway struct {
}

func NewRequestGateway() *RequestGateway {
	return &RequestGateway{}
}

func (r *RequestGateway) SendRequest(url string, c entity.Concurrency) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.Error <- err
		return
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		if response != nil {
			c.Status <- response.StatusCode
			return
		}
		c.Error <- err
		return
	}

	c.Status <- response.StatusCode

	defer func(response *http.Response) {
		err := response.Body.Close()
		if err != nil {
			c.Error <- err
			return
		}
	}(response)
}
