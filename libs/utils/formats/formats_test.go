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
