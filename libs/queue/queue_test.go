package queue

import (
	"testing"

	"github.com/web-notify/api/monorepo/libs/utils/config"
)

func TestQueue(t *testing.T) {
	var qService QueueServicer = &queueService{}
	t.Log(qService)
	t.Log("test passed, qService initialised without panic")
}

func TestGetConnectionString(t *testing.T) {
	connection := GetBaseConnectionString(config.DEV, "my_account")
	t.Log(connection)

	devAns := "http://127.0.0.1:10001/my_account"
	if connection != devAns {
		t.Fatalf("test failed. expectected %s, received %s", devAns, connection)
	}

	connection = GetBaseConnectionString(config.GetStage(), "my_account")
	t.Log(connection)

	stgAns := "https://my_account.queue.core.windows.net"
	if connection != stgAns {
		t.Fatalf("test failed. expectected %s, received %s", stgAns, connection)
	}
}
