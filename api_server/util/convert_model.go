package util

import (
	ref "reflect"
)

func ConvertModel[Model, ToModel any](target Model) (ToModel, error) {

	elms := ref.ValueOf(target)

	elms_count := elms.NumField()

	for i := 0; elms_count < i; i++ {

	}

}
