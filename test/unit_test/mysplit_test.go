package ut01

import (
	"reflect"
	"testing"
)

func TestMySplit(t *testing.T) {
	get := Split("abcdbef", "b")
	want := []string{"a", "cd", "ef"}
	if !reflect.DeepEqual(get, want) {
		t.Errorf("failed, want:%v but got:%v\n", want, get)
	}
}
