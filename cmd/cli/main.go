package main

import (
	"flag"
	"fmt"

	"github.com/israelalvesmelo/desafio-stress-test/internal/domain/dto"
	"github.com/israelalvesmelo/desafio-stress-test/internal/domain/usecase"
	"github.com/israelalvesmelo/desafio-stress-test/internal/infra"
)

func main() {
	var (
		url         = flag.String("url", "", "URL to test")
		maxRequests = flag.Int("requests", 0, "Maximum amount of requests to send")
		concurrency = flag.Int("concurrency", 0, "Amount of concurrent requests to send")
	)
	flag.Parse()

	requestFlag := dto.RequestFlag{
		URL:         *url,
		MaxRequests: *maxRequests,
		Concurrency: *concurrency,
	}
	requestFlag.Validate()

	gateway := infra.NewRequestGateway()
	mapper := infra.NewMapper()
	usecase := usecase.NewStressTestRequest(*gateway, *mapper)

	execResult, err := usecase.Execute(requestFlag)
	if err != nil {
		msgErr := fmt.Sprintf(`{"error": "%s"}`, err.Error())
		fmt.Println(msgErr)
	}

	fmt.Println(string(execResult))
}
