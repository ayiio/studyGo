package protocol

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

/*
大端BigEndian和小端LittleEndian：多字节数先写入高位是大端，先写入低位是小端
*/

// Encode 消息编码
func Encode(message string) ([]byte, error) {
	//获取消息长度，转为int32-4字节
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	//写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	//写入消息体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

// Decode 消息解码
func Decode(reader *bufio.Reader) (string, error) {
	//读取消息长度
	lengthByte, _ := reader.Peek(4) //读取4字节数据
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	//读取真正的消息数据
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}
