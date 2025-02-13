package utils

import (
	"errors"
	"fmt"
)

// AnyToString 将any类型转换为string
func AnyToString(value any) string {
	return fmt.Sprintf("%v", value)
}

func AnytoUint(a any) (uint, error) {
	if u, ok := a.(uint); ok {
		return u, nil
	}
	return 0, errors.New("value is not of type uint")
}
