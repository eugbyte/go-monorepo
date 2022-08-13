package queue

import (
	"context"
	"testing"

	"github.com/eugbyte/monorepo/libs/config"
)

func TestQueue(t *testing.T) {
	stage := config.Stage()
	accountName := "my_acc"
	url := config.QueueBaseURL(stage, accountName)
	var qService QueueServicer = NewQueueService(context.Background(), "my_queue", url, accountName, "my_key")
	t.Log(qService)
	t.Log("test passed, qService initialised without panic")
}
