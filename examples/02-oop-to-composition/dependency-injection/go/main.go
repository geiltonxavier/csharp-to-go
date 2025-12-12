package main

import (
	"fmt"
	"time"
)

type MessageSender interface {
	Send(to, message string)
}

type EmailSender struct{}

func (EmailSender) Send(to, message string) {
	fmt.Printf("[email] para %s: %s\n", to, message)
}

type Clock interface {
	Now() time.Time
}

type SystemClock struct{}

func (SystemClock) Now() time.Time {
	return time.Now().UTC()
}

type NotificationService struct {
	sender MessageSender
	clock  Clock
}

func NewNotificationService(sender MessageSender, clock Clock) *NotificationService {
	return &NotificationService{sender: sender, clock: clock}
}

func (n *NotificationService) Notify(user, message string) {
	timestamp := n.clock.Now().Format(time.RFC3339)
	n.sender.Send(user, fmt.Sprintf("[%s] %s", timestamp, message))
}

func main() {
	// Composition root: wiring explícito sem container.
	sender := EmailSender{}
	clock := SystemClock{}
	notifier := NewNotificationService(sender, clock)

	notifier.Notify("ana", "olá usando DI explícita em Go")
}
