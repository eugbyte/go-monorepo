package config

import (
	"os"
	"testing"
)

func TestHandler(t *testing.T) {
	myEnv := GetEnvOrDefault("MY_ENV", "abc")
	if myEnv != "abc" {
		t.Fatalf("test failed, expected %s, but received %s", "abc", myEnv)
	}
	os.Setenv("MY_ENV", "def")
	myEnv = GetEnvOrDefault("MY_ENV", "abc")
	if myEnv != "def" {
		t.Fatalf("test failed, expected %s, but received %s", "def", myEnv)
	}
	os.Unsetenv("MY_ENV")

	t.Log("test passed")
}
