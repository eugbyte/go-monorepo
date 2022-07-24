package vault

import (
	"fmt"
	"testing"
)

func TestVault(t *testing.T) {
	fmt.Println("Secret name must only contain only 0-9, a-z, A-Z, and -. ")
	var isValid bool = ValidateSecretName("dbName")
	if isValid == false {
		t.Fatalf("expected result to be valid, but received %v", isValid)
	}

	isValid = ValidateSecretName("db_name")
	if isValid == true {
		t.Fatalf("expected result to be false due to illegal '_', but received %v", isValid)
	}
}
