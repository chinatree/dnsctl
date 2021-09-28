package convert

import (
	"encoding/json"

	"github.com/axgle/mahonia"
)

// StructToJSON Struct 转 JSON
func StructToJSON(v interface{}) (string, error) {
	buf, err := json.Marshal(v)
	return string(buf), err
}

// StructToJSONWithIndent Struct 转 JSON 格式化输出
func StructToJSONWithIndent(v interface{}) (string, error) {
	buf, err := json.MarshalIndent(v, "", "\t")
	return string(buf), err
}

// StringToJSON String 转 JSON
func StringToJSON(s string) (string, error) {
	var m interface{}
	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		return "", err
	}
	buf, err := json.Marshal(m)
	return string(buf), err
}

// StringToJSONWithIndent String 转 JSON 格式化输出
func StringToJSONWithIndent(s string) (string, error) {
	var m interface{}
	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		return "", err
	}
	buf, err := json.MarshalIndent(m, "", "\t")
	return string(buf), err
}

// GBKToUTF8 GBK convert UTF-8
func GBKToUTF8(s string) (string, error) {
	srcCoder := mahonia.NewDecoder("gbk")
	srcResult := srcCoder.ConvertString(s)
	dstCoder := mahonia.NewDecoder("utf-8")

	_, data, err := dstCoder.Translate([]byte(srcResult), true)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
