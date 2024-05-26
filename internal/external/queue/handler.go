package queue

import (
	"context"
	"fmt"
	"log"

	queuetype "github.com/SaiHLu/rest-template/internal/external/queue/type"
	"github.com/SaiHLu/rest-template/internal/external/queue/worker"
	"github.com/hibiken/asynq"
)

func handler(ctx context.Context, t *asynq.Task) error {
	log.Println("t.Type(): ", t.Type())
	log.Println("queuetype.SendEmail.String(): ", queuetype.SendEmail.String())
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
