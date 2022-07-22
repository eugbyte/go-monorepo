package formats

import "testing"

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
