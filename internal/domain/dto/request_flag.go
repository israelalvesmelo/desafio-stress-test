package dto

import (
	"log"
	"net/url"
)

type RequestFlag struct {
	URL         string
	MaxRequests int
	Concurrency int
}

const (
	ErrURLIsRequired                    = "URL is required"
	ErrMaxRequestsIsRequired            = "MaxRequests is required"
	ErrConcurrencyIsRequired            = "Concurrency is required"
	ErrMaxRequestsIsLessThanConcurrency = "Concurrency must be less than or equal to MaxRequests"
	ErrInvalidURL                       = "Invalid URL"
)

func (r *RequestFlag) Validate() {
	if r.URL == "" {
		log.Fatal(ErrURLIsRequired)
	}
	if r.MaxRequests <= 0 {
		log.Fatal(ErrMaxRequestsIsRequired)
	}
	if r.Concurrency <= 0 {
		log.Fatal(ErrConcurrencyIsRequired)
	}
	if r.Concurrency > r.MaxRequests {
		log.Fatal(ErrMaxRequestsIsLessThanConcurrency)
	}
	if !r.isValidURL(r.URL) {
		log.Fatal(ErrInvalidURL)
	}
}

func (r *RequestFlag) isValidURL(u string) bool {
	parsedURL, err := url.Parse(u)
	if err != nil {
		return false
	}

	return (parsedURL.Scheme == "http" ||
		parsedURL.Scheme == "https") &&
		parsedURL.Host != ""
}
