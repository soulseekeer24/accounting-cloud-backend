package core

import (
	"piwi-backend-clean/communication/core/dispatcher/email"
)

type Module struct {
	emailDispatcher email.Dispatcher
}

func NewModule(emailSender email.Dispatcher) *Module {
	return &Module{emailDispatcher: emailSender}
}

func (m *Module) SendEmail(message *email.Message) (ok bool, err error) {

	return
}
