package bt01

import (
	"fmt"
	"strings"
)

// Split 字符串切割
func Split(str, seq string) (ret []string) {
	// example: str="abcdbe" seq="b" ret=[]string{"a", "cd", "e"}
	// 可对初始化的部分进行优化
	ret = make([]string, strings.Count(str, seq)+1) // 以分割符记数初始化固定容量的ret
	index := strings.Index(str, seq)
	for index > -1 {
		if string(str[:index]) != "" {
			//基准测试得到申请了4次内存，是因为ret未合理初始化，导致存在内容多次申请
			ret = append(ret, str[:index])
		}
		str = str[index+len(seq):]
		index = strings.Index(str, seq)
	}
	ret = append(ret, str)

	if 1 < 0 {
		// 验证测试覆盖率
		fmt.Println("can not run this part")
	}
	return
}

