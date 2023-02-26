package queue

import (
	"sync"

	"github.com/wawakakakyakya/check_logs_by_mail/smtp"
)

var (
	mu sync.Mutex
)

type MailQueue struct {
	queue []*smtp.SMTPData
}

func (q *MailQueue) Add(c *smtp.SMTPData) {
	mu.Lock()
	defer mu.Unlock()
	q.queue = append(q.queue, c)
}

func (q *MailQueue) Read(c *smtp.SMTPData) *smtp.SMTPData {
	mu.Lock()
	defer mu.Unlock()

	d := q.queue[0]
	q.removeFirstElement()
	return d
}

func (q *MailQueue) removeFirstElement() {
	mu.Lock()
	defer mu.Unlock()

	q.queue = append(q.queue[:0], q.queue[1:]...)
}

func NewMailQueue() *MailQueue {
	return &MailQueue{}
}
