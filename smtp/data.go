package smtp

import "strings"

type SMTPData struct {
	recipients []string
	subject    string
	msg        []byte
}

const crlf = "\r\n"

func (s *SMTPData) Body() []byte {
	to := "To: " + strings.Join(s.recipients, ",")
	subject := "Subject: " + s.subject + crlf
	return []byte(strings.Join([]string{to, subject, string(s.msg)}, crlf))
}

func NewSMTPData(recipients []string, subject string, msg []byte) *SMTPData {
	return &SMTPData{recipients: recipients, subject: subject, msg: msg}
}
