package main

import (
	"fmt"
)

type Notification interface {
	send(message string)
}

type Email struct{}

func (e *Email) send(message string) {
	fmt.Println("Sending message via email....")
}

type Sms struct{}

func (s *Sms) send(message string) {
	fmt.Println("Sending message via SMS")
}

func notificaionObjectFactory(messagerType string) Notification {
	switch messagerType {
	case "email":
		return &Email{}
	case "sms":
		return &Sms{}
	default:
		return nil
	}
}

func main() {
	email := notificaionObjectFactory("email")
	sms := notificaionObjectFactory("sms")

	email.send("Hello Via email")
	sms.send("Hello Via SMS")
}
