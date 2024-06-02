package queue

import (
	"context"
	"fmt"

	queuetype "github.com/SaiHLu/rest-template/internal/infrastructure/queue/type"
	"github.com/SaiHLu/rest-template/internal/infrastructure/queue/worker"
	"github.com/hibiken/asynq"
)

func handler(ctx context.Context, t *asynq.Task) error {
	switch t.Type() {
	case queuetype.SendEmail.String():
		if err := worker.HandleEmailDevlieryTask(ctx, t); err != nil {
			return err
		}

	default:
		return fmt.Errorf("unexpected task type: %s", t.Type())
	}

	return nil
}
