package email

type Dispatcher interface {
	SendEmail(email *Message) (ok bool, err error)
}
