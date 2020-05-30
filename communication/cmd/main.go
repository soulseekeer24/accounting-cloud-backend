package main

import (
	"fmt"
	"piwi-backend-clean/communication/core/dispatcher/email"
)

func main() {
	emailDispatcher := email.BuildSMTPDispatcher(
		"devcommunicationgo@gmail.com",
		"Mierda624.",
		"smtp.gmail.com", "587")

	msg := email.Message{
		Recipient: []string{"miguelesalador@gmail.com"},
		Body:      "Pappu sos la ostia",
	}

	ok, err := emailDispatcher.SendEmail(&msg)
	if err != nil {
		panic(err)
	}

	fmt.Println(ok)
}
