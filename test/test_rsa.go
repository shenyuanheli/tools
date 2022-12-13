/**
  @author: 1043193460@qq.com
  @date: 2022/12/13 15:35
  @note:
**/
package test

import (
	"fmt"
	"github.com/shenyuanheli/tools"
)

func main() {
	//生成密钥对，保存到文件
	tools.EncodeUtils.GenerateRSAKey(2048)
	message := []byte("大萨#￥%……#￥%鬼地方个31AADA3123fsdf！@#！@")
	fmt.Println("原始文字：", string(message))
	//加密
	cipherText := tools.EncodeUtils.RSA_Encrypt(message, "public.pem")
	fmt.Println("加密后为：", cipherText)
	//解密
	plainText := tools.EncodeUtils.RSA_Decrypt(cipherText, "private.pem")
	fmt.Println("解密后为：", plainText)
}
