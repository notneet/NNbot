package service

import "fmt"

// GreetService is a simple service for greeting messages
type GreetService struct{}

// NewGreetService creates a new instance of GreetService
func NewGreetService() *GreetService {
	return &GreetService{}
}

// GetWelcomeMessage generates a welcome message
func (s *GreetService) GetWelcomeMessage() string {
	return "Welcome to the API!"
}

// GetHelloMessage generates a hello message
func (s *GreetService) GetHelloMessage() string {
	return "Hello, Fiber!"
}

// GetCustomMessage generates a custom message using the provided name
func (s *GreetService) GetCustomMessage(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
