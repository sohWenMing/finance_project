package errorUtils

import (
	"fmt"
	"testing"
)

func AssertNoError(t *testing.T, testName string, err error) bool {
	if err != nil {
		fmt.Println("test name: ", testName)
		t.Errorf("didn't expect error, got %v\n", err)
		return false
	}
	return true

}

func AssertVals[K comparable](t *testing.T, testName string, got K, want K) bool {
	if got != want {
		fmt.Println("test name:", testName)
		t.Errorf("got:%v\n", got)
		t.Errorf("want:%v\n", want)
		return false
	}
	return true
}
