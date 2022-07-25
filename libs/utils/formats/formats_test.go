package formats

import (
	"fmt"
	"testing"
)

func Test_Stringify(t *testing.T) {
	obj := map[string]interface{}{
		"a": 123,
		"b": map[string][]int{
			"arr": {1, 2, 3},
		},
	}

	str := Stringify(obj)
	t.Log(str)
	ans := `{"a":123,"b":{"arr":[1,2,3]}}`
	t.Log("expected no panic, and no panic occured")
	if str != ans {
		t.Fatalf("expected %v, but received %v", ans, str)
	}

	t.Logf("test passed")
}

func Test_FormatURL(t *testing.T) {
	baseUrl := "http://localhost:7071/api"
	pathVariables := []string{"v1", "notifications"}
	queryParams := map[string]string{
		"page_number": "1",
		"page_length": "5",
	}

	res := FormatURL(baseUrl, pathVariables, queryParams)
	t.Log(res)
	ans := "http://localhost:7071/api/v1/notifications?page_number=1&page_length=5"
	if res != ans {
		t.Fatalf("expected %v, but received %v", ans, res)
	}
}

func Test_ValidateAzureParamString(t *testing.T) {
	fmt.Println("Secret name must only contain only 0-9, a-z, A-Z, and -. ")
	var isValid bool = ValidateAzureParamString("dbName")
	if isValid == false {
		t.Fatalf("expected result to be valid, but received %v", isValid)
	}

	isValid = ValidateAzureParamString("db_name")
	if isValid == true {
		t.Fatalf("expected result to be false due to illegal '_', but received %v", isValid)
	}
}
