package usecase

import (
	"fmt"
	"time"

	"github.com/israelalvesmelo/desafio-stress-test/internal/domain/dto"
	"github.com/israelalvesmelo/desafio-stress-test/internal/domain/entity"
	"github.com/israelalvesmelo/desafio-stress-test/internal/infra"
)

type StressTestRequest struct {
	requestGateway infra.RequestGateway
	mapper         infra.Mapper
}

func NewStressTestRequest(rq infra.RequestGateway, mapper infra.Mapper) *StressTestRequest {
	return &StressTestRequest{
		requestGateway: rq,
		mapper:         mapper,
	}
}

func (s *StressTestRequest) Execute(rf dto.RequestFlag) ([]byte, error) {
	fmt.Println("Starting stress test...")
	startTime := time.Now()

	var (
		statusChan = make(chan int, rf.MaxRequests)
		errChan    = make(chan error, rf.MaxRequests)
	)

	c := entity.Concurrency{
		Status: statusChan,
		Error:  errChan,
	}

	jobs := make(chan int, rf.MaxRequests)

	for range rf.Concurrency {
		go s.worker(rf.URL, c, jobs)
	}

	for i := range rf.MaxRequests {
		jobs <- i
	}
	close(jobs)

	response := s.processResponse(rf, startTime, &c)
	fmt.Println("Stress test finished...")
	return s.mapper.MarshalJSON(response)
}

func (s *StressTestRequest) worker(url string, c entity.Concurrency, jobs chan int) {
	for range jobs {
		s.requestGateway.SendRequest(url, c)
	}
}

func (s *StressTestRequest) processResponse(rf dto.RequestFlag, startTime time.Time, c *entity.Concurrency) *dto.Response {
	response := entity.NewResponse(rf.Concurrency)

	for range rf.MaxRequests {
		select {
		case status := <-c.Status:
			response.IncrementRequest()
			response.IncrementStatusMap(status)
		case err := <-c.Error:
			response.IncrementRequest()
			response.AddErrors(err)
		}
	}

	response.CalculateTotalDuration(startTime, time.Now())
	return dto.NewResponseByDomain(response)
}
