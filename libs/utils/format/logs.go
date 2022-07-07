package format

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	colors "github.com/TwinProduction/go-color"
)

func Trace(objs ...interface{}) {
	var strs []string
	for _, obj := range objs {
		bytes, err := (json.MarshalIndent(obj, "", "\t"))
		if err != nil {
			log.Panicln(err)
		}
		strs = append(strs, string(bytes))
	}
	fmt.Println(colors.Green, strings.Join(strs, " "), colors.Reset)
}

// Convert objects to string
func Stringify(obj interface{}) string {
	objBytes, err := json.Marshal(obj)
	if err != nil {
		log.Panicln(err)
	}

	return string(objBytes)
}
