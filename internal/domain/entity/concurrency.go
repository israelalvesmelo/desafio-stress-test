package entity

type Concurrency struct {
	Status chan int
	Error  chan error
}
