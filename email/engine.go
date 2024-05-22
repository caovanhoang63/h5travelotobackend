package email

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/asyncjob"
	"log"
	"sync"
	"time"
)

type Engine interface {
	Send(mail Mail)
	Start() error
	Stop()
}

type engine struct {
	mailSender MailSender
	queue      []asyncjob.Job
	worker     chan asyncjob.Job
	newJob     chan struct{}
	done       chan struct{}
	mu         sync.Mutex
}

func (e *engine) Send(mail Mail) {
	job := asyncjob.NewJob(func(ctx context.Context) error {
		return e.mailSender.Send(mail)
	})
	e.mu.Lock()
	e.queue = append(e.queue, job)
	e.mu.Unlock()
	e.newJob <- struct{}{}
}

func NewEngine(mailSender MailSender) *engine {
	return &engine{
		mailSender: mailSender,
		worker:     make(chan asyncjob.Job, 5),
		newJob:     make(chan struct{}, 1),
		done:       make(chan struct{}),
	}
}

func (e *engine) Start() error {
	go func() {
		defer common.AppRecover()
		for {
			select {
			case <-e.newJob:
				e.mu.Lock()
				if len(e.queue) != 0 {
					job := e.queue[0]
					e.worker <- job
					e.queue = append(e.queue[:0], e.queue[1:]...)
				}
				e.mu.Unlock()
			case <-e.done:
				return
			default:
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	go func() {
		defer common.AppRecover()
		for {
			select {
			case job := <-e.worker:
				go func() {
					defer common.AppRecover()
					if err := job.Execute(context.Background()); err != nil {
						for {
							log.Println(err)
							if job.State() == asyncjob.StateRetryFailed {
								log.Println("Err")
								break
							}

							if job.Retry(context.Background()) == nil {
								log.Println("Success")
								break
							}
						}
					}
				}()
			case <-e.done:
				return
			}
		}
	}()
	return nil
}

func (e *engine) Stop() {
	close(e.done)
}
