package main

import "fmt"

// Service is an interface that defines the behavior of a service.
type Service interface {
    Serve()
}

// MyService is a concrete implementation of the Service interface.
type MyService struct{}

// MyService is a concrete implementation of the Service interface.
type MyService2 struct{}

// Serve implements the Serve method of the Service interface.
func (s *MyService) Serve() {
    fmt.Println("Serving...")
}

// Serve implements the Serve method of the Service interface.
func (s *MyService2) Serve() {
    fmt.Println("Serving 2...")
}

// Consumer is a struct that depends on a service.
type Consumer struct {
    service Service
}

// NewConsumer creates a new Consumer with the provided Service dependency.
func NewConsumer(service Service) *Consumer {
    return &Consumer{service: service}
}

// UseService uses the injected Service.
func (c *Consumer) UseService() {
    c.service.Serve()
}

func main() {
    service := &MyService2{} // Create a concrete instance of the service.
    consumer := NewConsumer(service) // Inject the service dependency.
    consumer.UseService() // Use the injected service.
}
