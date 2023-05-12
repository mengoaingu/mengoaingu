package debezium

import (
	"encoding/json"
	"strings"
)

type Event[T any] struct {
	Schema  Schema     `json:"schema"`
	Payload Payload[T] `json:"payload"`
}

func (e Event[T]) UnmarshallInto(v interface{}) error {
	b, err := json.Marshal(e)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}

func (e Event[T]) Name() string {
	// <kafka-connect-name>.<db-name>.<table-name>.Envelope
	ss := strings.Split(e.Schema.Name, ".")
	if len(ss) < 4 {
		return ""
	}
	return ss[1] + ".cdc." + ss[2]
}

type Schema struct {
	// Name of the schema
	// docs https://debezium.io/documentation/reference/stable/connectors/mysql.html#mysql-change-event-keys
	Name string `json:"name"`
}
