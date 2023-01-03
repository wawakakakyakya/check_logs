package smtp

import "sync"

var (
	mu sync.Mutex
)

type MailQueue struct {
	queue []*SMTPClinet
}

func (q *MailQueue) Add(c *SMTPClinet) {
	mu.Lock()
	defer mu.Unlock()
	q.queue = append(q.queue, c)
}

type Queue struct {
}
