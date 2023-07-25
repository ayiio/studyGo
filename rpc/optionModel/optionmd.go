package optionmodel

import "fmt"

//结构体
type Options struct {
	stringOption1 string
	stringOption2 string
	stringOption3 string
	intOption1    int
	intOption2    int
	intOption3    int
}

//声明一个函数类型的变量，用于传参
type Option func(ops *Options)

//初始化结构体
func NewOptins(opfuncs ...Option) {
	ops := &Options{}
	//遍历opfuncs，得到每一个函数
	for _, opfunc := range opfuncs {
		//调用函数，在函数里，传入对象并赋值
		opfunc(ops)
	}
	fmt.Println(ops)
}

func SetStringOption1(s string) Option {
	return func(ops *Options) {
		ops.stringOption1 = s
	}
}

func SetStringOption2(s string) Option {
	return func(ops *Options) {
		ops.stringOption2 = s
	}
}

func SetStringOption3(s string) Option {
	return func(ops *Options) {
		ops.stringOption3 = s
	}
}

func SetIntOption1(i int) Option {
	return func(ops *Options) {
		ops.intOption1 = i
	}
}

func SetIntOption2(i int) Option {
	return func(ops *Options) {
		ops.intOption2 = i
	}
}

func SetIntOption3(i int) Option {
	return func(ops *Options) {
		ops.intOption3 = i
	}
}
