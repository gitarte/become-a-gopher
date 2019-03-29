package main

import (
	"crypto/tls"
	"log"
	"net/smtp"
)

func main() {
	//  create message
	m := Mail{}
	m.Sender = "some@address.com"
	m.To = []string{"to1@address.com", "to2@address.com"}
	m.Cc = []string{"cc1@address.com", "cc2@address.com"}
	m.Bcc = []string{"Bcc1@address.com", "Bcc2@address.com"}
	m.Subject = "Party invitation"
	m.Body = `Hi dude!
    Let's get schwifty tonight!
    Regards`

	// create connection with SMTP
	s := SMTP{
		Host: "smtp.provider.com",
		Port: 587, //465,
	}
	s.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         s.Host,
	}
	conn, err := tls.Dial("tcp", s.ComputeAddress(), s.TLSConfig)
	if err != nil {
		log.Panic(err)
	}
	c, err := smtp.NewClient(conn, s.Host)
	if err != nil {
		log.Panic(err)
	}
	defer c.Quit()

	// set authentication
	auth := smtp.PlainAuth("", m.Sender, "password", s.Host)
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	// set data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		err = w.Close()
		if err != nil {
			log.Panic(err)
		}
	}()
	_, err = w.Write(m.BuildMessage())
	if err != nil {
		log.Panic(err)
	}

	// set sender
	err = c.Mail(m.Sender)
	if err != nil {
		log.Panic(err)
	}

	// set recipients
	rcvs := []string{}
	rcvs = append(rcvs, m.To...)
	rcvs = append(rcvs, m.Cc...)
	rcvs = append(rcvs, m.Bcc...)
	for _, v := range rcvs {
		if err = c.Rcpt(v); err != nil {
			log.Panic(err)
		}
	}

	log.Println("Mail sent successfully")
}
