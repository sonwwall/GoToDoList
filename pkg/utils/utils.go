package utils

import "fmt"

// AnyToString 将any类型转换为string
func AnyToString(value any) string {
	return fmt.Sprintf("%v", value)
}
