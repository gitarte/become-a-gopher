package main

import (
	"fmt"
	"strings"
)

// Mail is a definition of object that holds basic attributes of E-Mail message
type Mail struct {
	Sender  string
	To      []string
	Cc      []string
	Bcc     []string
	Subject string
	Body    string
}

// BuildMessage is a helper method that computes E-Mail's header
func (mail *Mail) BuildMessage() []byte {
	header := ""
	header += fmt.Sprintf("From: %s\r\n", mail.Sender)
	if len(mail.To) > 0 {
		header += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	}
	if len(mail.Cc) > 0 {
		header += fmt.Sprintf("Cc: %s\r\n", strings.Join(mail.Cc, ";"))
	}
	if len(mail.Bcc) > 0 {
		header += fmt.Sprintf("Bcc: %s\r\n", strings.Join(mail.Cc, ";"))
	}

	header += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	header += "\r\n" + mail.Body

	return []byte(header)
}
