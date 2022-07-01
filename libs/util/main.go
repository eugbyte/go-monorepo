package util

import (
	"encoding/json"
	"fmt"
)

func Trace(prefix string, obj interface{}) {
	bytes, err := (json.MarshalIndent(obj, "", "\t"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(prefix+":", string(bytes))
}

// Convert objects to string
func Stringify(obj interface{}) string {
	objBytes, err := json.Marshal(obj)
	if err != nil {
		fmt.Println(err)
	}

	return string(objBytes)
}
