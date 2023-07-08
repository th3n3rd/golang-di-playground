package notification

import (
	"fmt"
	"golang-di-playground/policyholder"
)

type EmailService struct{}

func NewEmailService() Service {
	return &EmailService{}
}

func (s *EmailService) Notify(policyholder *policyholder.PolicyHolder, message string) {
	fmt.Printf("Sending email to %s: %s\n", policyholder.FullName, message)
}
