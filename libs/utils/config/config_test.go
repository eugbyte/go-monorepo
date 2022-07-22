package config

import (
	"os"
	"testing"
)

func TestHandler(t *testing.T) {
	myEnv := EnvOrDefault("MY_ENV", "abc")
	if myEnv != "abc" {
		t.Fatalf("test failed, expected %s, but received %s", "abc", myEnv)
	}
	os.Setenv("MY_ENV", "def")
	myEnv = EnvOrDefault("MY_ENV", "abc")
	if myEnv != "def" {
		t.Fatalf("test failed, expected %s, but received %s", "def", myEnv)
	}
	os.Unsetenv("MY_ENV")

	t.Log("test passed")
}

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
