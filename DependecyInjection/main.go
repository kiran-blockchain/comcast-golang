package main

import "fmt"

// MyService is a service struct.
type MyService struct{}

// Serve serves.
func (s *MyService) Serve() {
    fmt.Println("Serving...")
}

// Consumer is a struct that depends on MyService.
type Consumer struct {
    service MyService
}

// NewConsumer creates a new Consumer.
func NewConsumer() *Consumer {
    return &Consumer{service: MyService{}}
}

// UseService uses the service.
func (c *Consumer) UseService() {
    c.service.Serve()
}

func main() {
    consumer := NewConsumer()
    consumer.UseService()
}


