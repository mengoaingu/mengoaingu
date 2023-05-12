/**
 * Created by liemlhd on 30/Jan/2022
 */

package debezium

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Op int

const (
	OpCreate Op = iota
	OpRead
	OpUpdate
	OpDelete
)

var opNames = map[Op]string{
	OpCreate: "c",
	OpRead:   "r",
	OpUpdate: "u",
	OpDelete: "d",
}

var opValues = map[string]Op{
	"c": OpCreate,
	"r": OpRead,
	"u": OpUpdate,
	"d": OpDelete,
}

func (i Op) String() string {
	return opNames[i]
}

// OpString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func OpString(s string) (Op, error) {
	if val, ok := opValues[s]; ok {
		return val, nil
	}
	if val, ok := opValues[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to CampaignStatus values", s)
}

// MarshalJSON implements the json.Marshaler interface for CampaignStatus
func (i Op) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for CampaignStatus
func (i *Op) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("object should be a string, got %s", data)
	}

	var err error
	*i, err = OpString(s)
	return err
}

type Payload[T any] struct {
	Op          Op    `json:"op"`
	TimestampMs int64 `json:"ts_ms"`
	Before      T     `json:"before"`
	After       T     `json:"after"`
}
