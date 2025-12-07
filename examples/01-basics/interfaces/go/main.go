package main

import "fmt"

type MessageSender interface {
	Send(to, message string)
}

type EmailSender struct{}

func (EmailSender) Send(to, message string) {
	fmt.Printf("[email] para %s: %s\n", to, message)
}

type SmsSender struct{}

func (SmsSender) Send(to, message string) {
	fmt.Printf("[sms] para %s: %s\n", to, message)
}

type Notifier struct {
	sender MessageSender
}

func NewNotifier(sender MessageSender) Notifier {
	return Notifier{sender: sender}
}

func (n Notifier) Notify(user, message string) {
	n.sender.Send(user, message)
}

func main() {
	var sender MessageSender = EmailSender{}
	notifier := NewNotifier(sender)
	notifier.Notify("ana", "olá de Go")

	sender = SmsSender{}
	notifier = NewNotifier(sender)
	notifier.Notify("joao", "olá por SMS")
}
