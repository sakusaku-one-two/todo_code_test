package util

import (
	"reflect"
)

func IsNil(target any) bool {
	return target == nil || reflect.ValueOf(target).IsNil()
}
