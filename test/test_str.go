package test

import (
	"fmt"
	ztool "github.com/shenyuanheli/tools"
	"testing"
)

//驼峰转蛇形
func TestSnakeString(t *testing.T) {
	str := "userName"
	fmt.Printf("驼峰转蛇形------->%s", ztool.StrUtils.SnakeString(str))
}

//蛇形转驼峰
func TestCamelString(t *testing.T) {
	str := "user_name"
	fmt.Printf("蛇形转驼峰------->%s", ztool.StrUtils.CamelString(str))
}

//不可见字符串判空
func TestHasBlank(t *testing.T) {
	str := "    "
	str1 := ""
	str3 := "222222"
	fmt.Printf("不可见字符串------->%t\n", ztool.StrUtils.HasBlank(str))
	fmt.Printf("空字符串------->%t\n", ztool.StrUtils.HasBlank(str1))
	fmt.Printf("正常字符串------->%t\n", ztool.StrUtils.HasBlank(str3))
}

//判断是否是空字符串
func TestHasEmpty(t *testing.T) {
	str := "    "
	str1 := ""
	str2 := "test"
	fmt.Printf("不可见字符串------->%t\n", ztool.StrUtils.HasEmpty(str))
	fmt.Printf("空字符串------->%t\n", ztool.StrUtils.HasEmpty(str1))
	fmt.Printf("正常字符串------->%t\n", ztool.StrUtils.HasEmpty(str2))
}

//删除文件后缀获取文件名
func TestRemoveSuffix(t *testing.T) {
	s := "/opt/image.png"
	fmt.Printf("删除文件后缀获取文件名------->%s\n", ztool.StrUtils.RemoveSuffix(s))
}

//获取文件拓展名
func TestRemovePrefix(t *testing.T) {
	s := "/opt/image.png"
	fmt.Printf("获取文件拓展名------->%s\n", ztool.StrUtils.RemovePrefix(s))
}

//字符串16进制编码
func TestEncodeHexString(t *testing.T) {
	s := "hello world"
	fmt.Printf("编码前------->%s\n", s)
	fmt.Printf("编码后------->%s\n", ztool.StrUtils.EncodeHexString(s))
}

//16进制转字符串
func TestDecodeHexString(t *testing.T) {
	s := "hello world"
	hexString := ztool.StrUtils.EncodeHexString(s)
	fmt.Printf("解码前------->%s\n", hexString)
	decodeHexString, err := ztool.StrUtils.DecodeHexString(hexString)
	if err != nil {
		return
	}
	fmt.Printf("解码后------->%s\n", decodeHexString)
}
