package grouptest

import (
	"reflect"
	"testing"
)

// 组测试和子测试

// 运行整体测试: go test -v
func TestSplit(t *testing.T) {
	type testCases struct {
		str  string
		seq  string
		want []string
	}

	// 以slice包装的旧方法，无法满足单独运行组内子测试的目的
	testGroup := []testCases{
		{"abcdbef", "b", []string{"a", "cd", "ef"}},          // case1
		{"abc:db:e:f", ":", []string{"abc", "db", "e", "f"}}, // case2
		{"abbcdbbef", "bb", []string{"a", "cd", "ef"}},       // case3
		{"打麻将的张三有三张麻将打", "麻将", []string{"打", "的张三有三张", "打"}}, // case4
	}

	for i, tc := range testGroup {
		got := Split(tc.str, tc.seq)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("the %dth case failed, want:%v but got:%v\n", i+1, tc.want, got)
		}
	}
}

// 实现子测试
// 单独运行某一子案例， go test -run=Test2Split/case4
func Test2Split(t *testing.T) {
	type testCases struct {
		str, sep string
		want     []string
	}
	testGroup := map[string]testCases{
		"case1": {"abcdbef", "b", []string{"a", "cd", "ef"}},
		"case2": {"abc:db:e:f", ":", []string{"abc", "db", "e", "f"}},
		"case3": {"abbcdbbef", "bb", []string{"a", "cd", "ef"}},
		"case4": {"打麻将的张三有三张麻将打", "麻将", []string{"打", "的张三有三张", "打"}},
	}
	for casex, tc := range testGroup {
		t.Run(casex, func(t *testing.T) {
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("failed, want:%v but got:%v\n", tc.want, got)
			}
		})
	}
}

// 测试覆盖率: go test -cover
// 结果输出到文件: go tool cover -coverprofile=c.out
// coverproflie: html/mode/o/var
// 测试函数覆盖率要满足100%, 实际测试代码覆盖率一般要大于60%(iferr判断不能满足完全覆盖)
