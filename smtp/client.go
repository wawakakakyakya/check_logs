package smtp

import (
	"errors"
	"fmt"
	"io"
	"net/smtp"
	"strings"
)

var (
	CRLF string = "\r\n"
)

type SMTPClient struct {
	client   *smtp.Client
	addr     string
	userName string
	password string
}

func NewSMTPClient(hostname string, port int, username string, password string) *SMTPClient {
	return &SMTPClient{addr: fmt.Sprintf("%s:%s", hostname, port), userName: username, password: password}
}

func (s *SMTPClient) getHost() string {
}

func (s *SMTPClient) Dial() error {

	if s.client != nil {
		return errors.New("smtp client is not nil")
	}

	client, err := smtp.Dial(fmt.Sprintf("%s:%d", s.hostName, s.port))
	if err != nil {
		return err
	}
	s.client = client
	return nil
}

func (s *SMTPClient) isConnect() error {

	if s.client == nil {
		return errors.New("smtp client client is nil")
	} else if err := s.client.Noop(); err != nil {
		return err
	}
	return nil
}

func (s *SMTPClient) Auth() error {
	err := s.isConnect()
	if err != nil {
		return err
	}
	return s.client.Auth(smtp.CRAMMD5Auth(s.userName, s.password))
}

func (s *SMTPClient) setRecipients(data *SMTPData) error {
	for _, addr := range data.Recipients {
		if err := s.client.Rcpt(addr); err != nil {
			return err
		}
	}
	return nil
}

func (s *SMTPClient) setMailFrom(mailFrom string) error {
	if err := s.client.Mail(mailFrom); err != nil {
		return err
	}
	return nil
}

func (s *SMTPClient) Send(data *SMTPData) error {

	// auth := smtp.CRAMMD5Auth(s.userName, s.password)
	// msg := []byte(strings.ReplaceAll(fmt.Sprintf("To: %s\nSubject: %s\n\n%s", strings.Join(data.Recipients, ","), data.Subject, "bodybody"), "\n", "\r\n"))
	// if err := smtp.SendMail(s.addr, auth, data.From, data.Recipients, msg); err != nil {
	// 	return err
	// }

	//reuse need server setting
	w, err := s.client.Data()
	if err != nil {
		return err
	}
	defer w.Close()
	msg := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", strings.Join(data.Recipients, ","), data.Subject, "bodybody")
	if _, err := io.WriteString(w, msg); err != nil {
		return err
	}

	err = s.client.Quit()
	if err != nil {
		return err
	}
	return nil
}

func (s *SMTPClient) Clsoe() {
	s.client.Close()
}

func main() {
	// from := "gopher@example.net"
	// recipients := []string{"foo@example.com", "bar@example.com"}
	// subject := "hello"
	body := "Hello World!\nHello Gopher!"

	defer client.Close()

	// if err := client.Mail(from); err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// 	return
	// }

	// for _, addr := range recipients {
	// 	if err := client.Rcpt(addr); err != nil {
	// 		fmt.Fprintln(os.Stderr, err)
	// 		return
	// 	}
	// }

	// if err := func() error {
	// 	w, err := client.Data()
	// 	if err != nil {
	// 		return err
	// 	}
	// 	defer w.Close()
	// 	msg := strings.ReplaceAll(fmt.Sprintf("To: %s\nSubject: %s\n\n%s", strings.Join(recipients, ","), subject, body), "\n", "\r\n")
	// 	if _, err := io.WriteString(w, msg); err != nil {
	// 		return err
	// 	}
	// 	return nil
	// }(); err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// 	return
	// }

	// if err := client.Quit(); err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// }
}
