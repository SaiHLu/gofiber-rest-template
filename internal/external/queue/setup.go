package queue

import (
	"log"
	"net/http"

	"github.com/hibiken/asynq"
	"github.com/hibiken/asynqmon"
)

type queue struct {
	RedisAddr string
}

func NewQueue(RedisAddr string) *queue {
	return &queue{RedisAddr}
}

func (q *queue) SetupQueueClient() *asynq.Client {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: q.RedisAddr})

	return client
}

func (q *queue) ExecuteQueue() {
	queueServer := asynq.NewServer(
		asynq.RedisClientOpt{Addr: q.RedisAddr},
		asynq.Config{Concurrency: 2,
			Queues: map[string]int{
				"high":   6,
				"medium": 3,
				"low":    1,
			},
		},
	)

	if err := queueServer.Start(asynq.HandlerFunc(handler)); err != nil {
		log.Fatalln(err)
	}
}

func (q *queue) MonitorQueues() {
	queueMonitor := asynqmon.New(asynqmon.Options{
		RootPath:     "/monitoring",
		RedisConnOpt: asynq.RedisClientOpt{Addr: q.RedisAddr},
	})

	http.Handle(queueMonitor.RootPath()+"/", queueMonitor)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
