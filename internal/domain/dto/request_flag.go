package dto

import (
	"fmt"
	"net/url"
)

type RequestFlag struct {
	URL         string
	MaxRequests int
	Concurrency int
}

const (
	ErrURLIsRequired                    = "flag -url is required"
	ErrMaxRequestsIsRequired            = "flag -requests is required"
	ErrConcurrencyIsRequired            = "flag -concurrency is required"
	ErrMaxRequestsIsLessThanConcurrency = "flag -concurrency must be less than or equal to -requests"
	ErrInvalidURL                       = "invalid -url"
)

func (r *RequestFlag) Validate() error {
	if r.URL == "" {
		return fmt.Errorf(ErrURLIsRequired)
	}
	if r.MaxRequests <= 0 {
		return fmt.Errorf(ErrMaxRequestsIsRequired)
	}
	if r.Concurrency <= 0 {
		return fmt.Errorf(ErrConcurrencyIsRequired)
	}
	if r.Concurrency > r.MaxRequests {
		return fmt.Errorf(ErrMaxRequestsIsLessThanConcurrency)
	}
	if !r.isValidURL(r.URL) {
		return fmt.Errorf(ErrInvalidURL)
	}
	return nil
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
