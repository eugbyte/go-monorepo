package queue

import (
	"testing"

	"github.com/web-notify/api/monorepo/libs/utils/config"
)

func TestQueue(t *testing.T) {
	var qService QueueServiceImpl = &QueueService{}
	t.Log(qService)
	t.Log("test passed, qService initialised without panic")
}

func TestGetConnectionString(t *testing.T) {
	t.Log("stage", config.STAGE)
	connection := GetConnectionString(config.STAGE, "my_account")

	t.Log(connection)
	devAns := "http://127.0.0.1:10001/my_account/my_queue"
	if connection != devAns {
		t.Fatalf("test failed. expectected %s, received %s", devAns, connection)
	}
}
