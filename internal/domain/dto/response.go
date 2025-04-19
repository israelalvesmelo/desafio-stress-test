package dto

import (
	"github.com/israelalvesmelo/desafio-stress-test/internal/domain/entity"
)

type Response struct {
	Requests        int         `json:"requests"`
	Concurrency     int         `json:"concurrency"`
	TotalDuration   string      `json:"total_duration"`
	AverageDuration string      `json:"average_duration"`
	Status          map[int]int `json:"status"`
	Errors          []string    `json:"errors"`
	ErrorsCount     int         `json:"errors_count"`
}

func NewResponseByDomain(r *entity.Response) *Response {
	return &Response{
		Requests:        r.Requests(),
		Concurrency:     r.Concurrency(),
		TotalDuration:   r.TotalDuration().String(),
		AverageDuration: r.AverageDuration().String(),
		Status:          r.StatusMap(),
		Errors:          r.ErrorsList(),
		ErrorsCount:     r.ErrorsCount(),
	}
}
