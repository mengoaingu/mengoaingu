package dbutils

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type SliceInt64 []int64

func (a *SliceInt64) Value() (driver.Value, error) {
	val := strings.Join(lo.Map(*a, func(t int64, i int) string {
		return strconv.FormatInt(t, 10)
	}), ",")
	return val, nil
}
func (a *SliceInt64) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of SliceInt64: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return nil
	}

	val := strings.Split(str, ",")
	var result []int64
	for _, v := range val {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return err
		}
		result = append(result, i)
	}
	*a = result
	return nil
}

type SliceString []string

func (a *SliceString) Value() (driver.Value, error) {
	val := strings.Join(*a, ",")
	return val, nil
}
func (a *SliceString) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of SliceString: %[1]T(%[1]v)", value)
	}
	if len(str) == 0 {
		return nil
	}

	val := strings.Split(str, ",")
	*a = val
	return nil
}
