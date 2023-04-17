package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type MySqlConf struct {
	Address  string `ini:"address"`
	Port     int64  `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConf struct {
	Host     string `ini:"host"`
	Port     int64  `ini:"port"`
	Password string `ini:"password"`
	User     string `ini:"user"`
}

type ConfIni struct {
	MySqlConf `ini:"mysql"`
	RedisConf `ini:"redis"`
}

func loadIni(filepath string, data interface{}) (err error) {
	//参数校验, 指针类型
	td := reflect.TypeOf(data)
	if td.Kind() != reflect.Ptr {
		err = errors.New("data param should be a pointer")
		return
	}
	//结构体指针
	if td.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be a struct pointer")
		return
	}

	//1.读文件
	bcont, err := ioutil.ReadFile(filepath)
	if err != nil {
		err = fmt.Errorf("无法打开%s处的文件,err=%v", filepath, err)
		return
	}
	//2.以换行符切割
	scont := strings.Split(string(bcont), "\r\n")
	// fmt.Println(scont)
	//3.处理文件内容
	var structName string
	for idx, cont := range scont {
		line := strings.Trim(cont, " ")
		//处理空行和注释行
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		//以[]节的样式
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("err1, line: %d syntax err", idx+1)
				return
			}
			//取到节的内容
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("err2, line: %d syntax err", idx)
				return
			}
			//根据section去data中根据反射找到对应结构体
			for i := 0; i < td.Elem().NumField(); i++ {
				fieldObj := td.Elem().Field(i)
				if sectionName == fieldObj.Tag.Get("ini") {
					structName = fieldObj.Name
				}
			}
		} else {
			//不以[]节的样式
			//以=号开头或不含=则无效
			if !strings.Contains(line, "=") || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("err3, line: %d syntax err", idx+1)
				return
			}

			valdata := reflect.ValueOf(data)
			structVal := valdata.Elem().FieldByName(structName) //获取嵌套结构体值信息
			structTyp := structVal.Type()                       //获取嵌套结构体类型信息

			if structTyp.Kind() != reflect.Struct {
				err = fmt.Errorf("err4, line: %d syntax err", idx)
				return
			}

			key, val := strings.Trim((strings.Split(line, "=")[0]), " "), strings.Trim((strings.Split(line, "=")[1]), " ")
			var fieldName string
			var fieldType reflect.StructField
			for j := 0; j < structVal.NumField(); j++ {
				fieldType = structTyp.Field(j)
				if key == structTyp.Field(j).Tag.Get("ini") { //Tag存储在类型信息中
					//找到字段
					fieldName = structTyp.Field(j).Name

					//赋值， 根据fieldName取出字段并赋值
					field := structVal.FieldByName(fieldName)
					// fmt.Println(fieldName, fieldType)
					switch fieldType.Type.Kind() {
					case reflect.String:
						field.SetString(val)
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						intv, _ := strconv.Atoi(val)
						field.SetInt(int64(intv))
					}
					break
				}
			}
		}

	}

	return nil

}

func main() {
	var conf ConfIni
	err := loadIni("conf.ini", &conf)
	if err != nil {
		fmt.Println("出现错误", err)
	}
	fmt.Printf("%#v\n", conf)
}
