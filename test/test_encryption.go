/**
  @author: fanyanan
  @date: 2022/6/11
  @note: //加密工具测试
**/
package test

import (
	"fmt"
	ztool "github.com/shenyuanheli/tools"
	"testing"
)

//AES加密方法
func TestAesEncrypt(t *testing.T) {
	str := "1234567789"
	//自定义加密key
	key := "ksjdu7372hsy43l;"
	encrypt, _ := ztool.EncryptionUtils.AesEncrypt(str)
	fmt.Printf("使用默认key进行加密---------->%s\n", encrypt)
	aesEncrypt, _ := ztool.EncryptionUtils.AesEncrypt(str, key)
	fmt.Printf("使用自定义key进行加密---------->%s\n", aesEncrypt)
}

//AES解密方法
func TestAesDecrypt(t *testing.T) {
	str := "1234567789"
	encrypt, _ := ztool.EncryptionUtils.AesEncrypt(str)
	fmt.Printf("解密前---------->%s\n", encrypt)
	decrypt, err := ztool.EncryptionUtils.AesDecrypt(encrypt)
	if err != nil {
		return
	}
	fmt.Printf("解密后---------->%s\n", decrypt)
}

//DES加密方法
func TestDesEncrypt(t *testing.T) {
	str := "1234567789"
	//自定义加密key
	key := "ujugygyi"
	encrypt, _ := ztool.EncryptionUtils.DesEncrypt(str)
	fmt.Printf("使用默认key进行加密---------->%s\n", encrypt)
	desEncrypt, _ := ztool.EncryptionUtils.DesEncrypt(str, key)
	fmt.Printf("使用自定义key进行加密---------->%s\n", desEncrypt)
}

//DES解密方法
func TestDesDecrypt(t *testing.T) {
	str := "1234567789"
	encrypt, _ := ztool.EncryptionUtils.DesEncrypt(str)
	fmt.Printf("解密前---------->%s\n", encrypt)
	decrypt, err := ztool.EncryptionUtils.DesDecrypt(encrypt)
	if err != nil {
		return
	}
	fmt.Printf("解密后---------->%s\n", decrypt)
}

//Md5进行签名
func TestMd5Encode(t *testing.T) {
	str := "123456789q"
	fmt.Printf("签名前-----------------%s\n", str)
	fmt.Printf("签名后-----------------%s\n", ztool.EncryptionUtils.Md5Encode(str))
}

//Md5签名验证，校验签名前和签名后是否是一个字符串
func TestMd5Check(t *testing.T) {
	str := "123456789q"
	fmt.Printf("签名前-----------------%s\n", str)
	encode := ztool.EncryptionUtils.Md5Encode(str)
	fmt.Printf("签名后-----------------%s\n", encode)
	fmt.Printf("签名校验-----------------%t\n", ztool.EncryptionUtils.Md5Check(str, encode))
}
