package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/SaiHLu/rest-template/internal/infrastructure/queue/task"
	"github.com/hibiken/asynq"
)

func HandleEmailDevlieryTask(ctx context.Context, t *asynq.Task) error {
	var p task.SendEmailPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Printf("Sending Email to User: user_id=%d, template_id=%s", p.UserId, p.TemplateId)

	return nil
}
