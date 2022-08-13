package queue

import (
	"context"
	"testing"

	"github.com/eugbyte/monorepo/libs/config"
)

func TestQueue(t *testing.T) {
	queueName := "my-queue"
	var stage config.STAGE = config.Stage()
	queueAccountName := "devstoreaccount1"

	qBaseUrl := config.QueueBaseURL(stage, queueAccountName)
	accountName := "my_acc"
	key := "Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw=="
	var qService QueueServicer = NewQueueService(context.Background(), queueName, qBaseUrl, accountName, key)
	t.Log(qService)
	t.Log("test passed, qService initialised without panic")
}
