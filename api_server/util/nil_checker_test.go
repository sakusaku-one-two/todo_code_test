package util

import (
	"fmt"
	"testing"
)

type sampleStruct struct {
}

func Test_nil_checker(t *testing.T) {
	sample := &sampleStruct{}

	result := IsNil(sample)
	fmt.Println(sample, result)

	sample = nil
	result = IsNil(sample)
	fmt.Println(sample, result)

}
