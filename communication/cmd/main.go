package main

import (
	"flag"
	"fmt"
	"piwi-backend-clean/communication/core/dispatcher/email"
)

func main() {

	emailPassFlag := flag.String("EMAIL_PASS", "", "password")

	flag.Parse()
	pass := *emailPassFlag

	emailDispatcher := email.BuildSMTPDispatcher(
		"devcommunicationgo@gmail.com",
		pass,
		"smtp.gmail.com", "587")

	msg := email.Message{
		Recipient: []string{"miguelesalador@gmail.com", "manuelclimb24@gmail.com", "jejoalca14@gmail.com"},
		Body:      "SOMOS EL PUTO AMAZON TIO!!!",
	}

	ok, err := emailDispatcher.SendEmail(&msg)
	if err != nil {
		panic(err)
	}

	fmt.Println(ok)
}
