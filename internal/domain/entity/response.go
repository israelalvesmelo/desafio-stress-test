package entity

import (
	"sync"
	"time"
)

type Response struct {
	requests      int
	concurrency   int
	statusMap     map[int]int
	totalDuration time.Duration
	errors        []error
	SyMu          sync.Mutex
}

func NewResponse(concurrency int) *Response {
	return &Response{
		concurrency: concurrency,
		statusMap:   make(map[int]int),
	}
}

func (r *Response) IncrementStatusMap(statusCode int) {
	r.SyMu.Lock()
	defer r.SyMu.Unlock()
	r.statusMap[statusCode]++
}

func (r *Response) CalculateTotalDuration(start, end time.Time) {
	r.totalDuration = end.Sub(start)
}

func (r *Response) IncrementRequest() {
	r.SyMu.Lock()
	defer r.SyMu.Unlock()
	r.requests++
}

func (r *Response) AddErrors(err error) {
	r.SyMu.Lock()
	defer r.SyMu.Unlock()
	r.errors = append(r.errors, err)
}

func (r *Response) TotalDuration() time.Duration {
	return r.totalDuration
}

func (r *Response) AverageDuration() time.Duration {
	return r.totalDuration / time.Duration(r.requests)
}

func (r *Response) Concurrency() int {
	return r.concurrency
}

func (r *Response) Requests() int {
	return r.requests
}

func (r *Response) StatusMap() map[int]int {
	return r.statusMap
}

func (r *Response) Errors() []error {
	return r.errors
}

func (r *Response) ErrorsCount() int {
	return len(r.errors)
}

func (r *Response) ErrorsList() []string {
	errMsgList := make([]string, 0)
	for _, e := range r.errors {
		errMsgList = append(errMsgList, e.Error())
	}
	return errMsgList
}
