package grouptest

import (
	"fmt"
	"strings"
)

// Split 字符串切割
func Split(str, seq string) (ret []string) {
	// example: str="abcdbe" seq="b" ret=[]string{"a", "cd", "e"}
	index := strings.Index(str, seq)
	for index > -1 {
		if string(str[:index]) != "" {
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
