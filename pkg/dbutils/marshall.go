package dbutils

import (
	"encoding/json"
	"fmt"

	"gorm.io/datatypes"
)

// Marshall will merge multiple maps into one
func Marshall(ms ...map[string]interface{}) datatypes.JSON {
	var merged = make(map[string]interface{})
	for _, m := range ms {
		for k, v := range m {
			merged[k] = v
		}
	}
	b, err := json.Marshal(merged)
	if err != nil {
		fmt.Println("Marshall error:", err)
	}
	return b
}
func Unmarshall(j datatypes.JSON) map[string]interface{} {
	var m map[string]interface{}
	_ = json.Unmarshal(j, &m)
	return m
}
