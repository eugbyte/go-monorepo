package config

import "testing"

func TestQueueBaseURL(t *testing.T) {
	connection := QueueBaseURL(DEV, "my_account")
	t.Log(connection)

	devAns := "http://127.0.0.1:10001/my_account"
	if connection != devAns {
		t.Fatalf("test failed. expectected %s, received %s", devAns, connection)
	}

	connection = QueueBaseURL(STAGING, "my_account")
	t.Log(connection)

	stgAns := "https://my_account.queue.core.windows.net"
	if connection != stgAns {
		t.Fatalf("test failed. expectected %s, received %s", stgAns, connection)
	}
}
