package workflow

import "golang-di-playground/policyholder"

type FakeNotificationService struct {
	Messages []string
}

func (p *FakeNotificationService) Notify(policyholder *policyholder.PolicyHolder, message string) {
	p.Messages = append(p.Messages, message)
}

func (p *FakeNotificationService) NextMessage() (string, bool) {
	if len(p.Messages) == 0 {
		return "", false
	}
	message := p.Messages[0]
	p.Messages = p.Messages[1:]
	return message, true
}
