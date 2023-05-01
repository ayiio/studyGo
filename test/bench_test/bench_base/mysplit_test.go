package bt01

import "testing"

// 基准测试
// 运行命令： go test -bench=Splits
// 追加参数： -benchmem 表示内存使用情况
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c:d:e", ":")
	}
}
