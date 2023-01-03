package smtp

type SMTPData struct {
	Recipients []string
	Subject    string
	Body       string
	From       string
}

func NewSMTPData(recipients []string, subject string) *SMTPData {
	return &SMTPData{Recipients: recipients, Subject: subject}
}
