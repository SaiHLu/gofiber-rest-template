package task

import (
	"encoding/json"
	"log"

	queuetype "github.com/SaiHLu/rest-template/internal/infrastructure/queue/type"
	"github.com/hibiken/asynq"
)

type SendEmailPayload struct {
	UserId     int
	TemplateId string
}

func NewEmailDeliveryTask(userId int, templateId string) (*asynq.Task, error) {
	payload, err := json.Marshal(SendEmailPayload{UserId: userId, TemplateId: templateId})
	if err != nil {
		return nil, err
	}

	log.Println("queue.SendEmail.String(): ", queuetype.SendEmail.String())

	return asynq.NewTask(queuetype.SendEmail.String(), payload), nil
}
