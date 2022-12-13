/**
  @author: 1043193460@qq.com
  @date: 2022/12/13 15:24
  @note:
**/
package encryption

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"os"
)

//生成RSA私钥和公钥，保存到文件中
func (r *Encryption) GenerateRSAKey(bits int, params ...string) {
	saveDir := "./"
	if len(params) > 0 {
		saveDir = params[0]
		if err := os.MkdirAll(saveDir, 0755); err != nil {
			panic(err)
		}
	}
	//GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥
	//Reader是一个全局、共享的密码用强随机数生成器
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
	}
	//保存私钥
	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	//使用pem格式对x509输出的内容进行编码
	//创建文件保存私钥
	privateFile, err := os.Create(saveDir + "private.pem")
	if err != nil {
		panic(err)
	}
	defer privateFile.Close()
	//构建一个pem.Block结构体对象
	privateBlock := pem.Block{Type: "RSA Private Key", Bytes: X509PrivateKey}
	//将数据保存到文件
	pem.Encode(privateFile, &privateBlock)
	//保存公钥
	//获取公钥的数据
	publicKey := privateKey.PublicKey
	//X509对公钥编码
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	//pem格式编码
	//创建用于保存公钥的文件
	publicFile, err := os.Create(saveDir + "public.pem")
	if err != nil {
		panic(err)
	}
	defer publicFile.Close()
	//创建一个pem.Block结构体对象
	publicBlock := pem.Block{Type: "RSA Public Key", Bytes: X509PublicKey}
	//保存到文件
	pem.Encode(publicFile, &publicBlock)
}

//RSA加密
func (r *Encryption) RSA_Encrypt(plainText []byte, path string) string {
	//打开文件
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//读取文件的内容
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	//pem解码
	block, _ := pem.Decode(buf)
	//x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	//对明文进行加密
	cipherByte, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		panic(err)
	}
	//hex加密输出
	cipherText := r.opputStr(cipherByte)
	//返回密文
	return cipherText
}

//RSA解密
func (r *Encryption) RSA_Decrypt(cipherText string, path string) string {
	//打开文件
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//获取文件内容
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	//pem解码
	block, _ := pem.Decode(buf)
	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//对密文进行解密
	cipherByte := r.inputStr(cipherText)
	plainText, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherByte)
	//返回明文
	return string(plainText)
}

// byte类型输出str
func (r *Encryption) opputStr(cipherByte []byte) string {
	//hex加密输出
	cipherText := hex.EncodeToString(cipherByte)
	//base64加密输出
	cipherText = base64.URLEncoding.EncodeToString(cipherByte)
	return cipherText
}

// str类型输出byte类型
func (r *Encryption) inputStr(cipherText string) []byte {
	//hex解密输出
	cipherByte, _ := hex.DecodeString(cipherText)
	//base64加密输出
	cipherByte, _ = base64.URLEncoding.DecodeString(cipherText)
	return cipherByte
}

// 签名：采用sha1算法进行签名并输出为hex格式（私钥PKCS8格式）
func (r *Encryption) RsaSignWithSha1Hex(data string, prvKey string) (string, error) {
	keyByts, err := hex.DecodeString(prvKey)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(keyByts)
	if err != nil {
		fmt.Println("ParsePKCS8PrivateKey err", err)
		return "", err
	}
	h := sha1.New()
	h.Write([]byte([]byte(data)))
	hash := h.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA1, hash[:])
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
		return "", err
	}
	out := hex.EncodeToString(signature)
	return out, nil
}

// 验签：对采用sha1算法进行签名后转base64格式的数据进行验签
func (r *Encryption) RsaVerySignWithSha1Base64(originalData, signData, pubKey string) error {
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return err
	}
	public, _ := base64.StdEncoding.DecodeString(pubKey)
	pub, err := x509.ParsePKIXPublicKey(public)
	if err != nil {
		return err
	}
	hash := sha1.New()
	hash.Write([]byte(originalData))
	return rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA1, hash.Sum(nil), sign)
}
