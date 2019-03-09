package mail

import (
	"testing"
)

func TestSendMail(t *testing.T) {
	recipient := "mischa@joerg.ch"
	subject := "test"
	message := "hello world"
	SendMail(recipient, subject, message)
}
