package ut01

import (
	"reflect"
	"testing"
)

// unit test : go test -v

func TestMySplit(t *testing.T) {
	got := Split("abcdbef", "b")
	want := []string{"a", "cd", "ef"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("failed, want:%v but got:%v\n", want, got)
	}
}

func Test2MySplit(t *testing.T) {
	got := Split("abc:db:e:f", ":")
	want := []string{"abc", "db", "e", "f"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("failed, want:%v but got:%v\n", want, got)
	}
}

func Test3MySplit(t *testing.T) {
	got := Split("abbcdbbef", "bb")
	want := []string{"a", "cd", "ef"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("failed, want:%v but got:%v\n", want, got)
	}
}

func Test4MySplit(t *testing.T) {
	got := Split("打麻将的张三有三张麻将打", "麻将")
	want := []string{"打", "的张三有三张", "打"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("failed, want:%v but got:%v\n", want, got)
	}
}
