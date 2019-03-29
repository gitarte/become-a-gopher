package main

import (
	"crypto/tls"
	"strconv"
)

// SMTP is a definition of object that holds basic attributes of communication with SMTP server
type SMTP struct {
	Host      string
	Port      int
	TLSConfig *tls.Config
}

// ComputeAddress is a helper method that joins hostname with port number of ther SMTP server
func (s *SMTP) ComputeAddress() string {
	return s.Host + ":" + strconv.Itoa(s.Port)
}
