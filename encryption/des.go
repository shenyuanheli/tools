/**
  @author: 1043193460@qq.com
  @date: 2022/12/13 14:32
  @note:
**/
package encryption

import (
	"crypto/des"
	"encoding/hex"
	"errors"
)

//DesEncrypt des encrypted function
//salt The parameter is the salt value of encryption and decryption, and the maximum
//length is 8. If the length is exceeded, the corresponding exception will be thrown
func (*Encryption) DesEncrypt(text string, key ...string) (string, error) {
	k := []byte(defaultDesKey)
	if len(key) > 0 {
		k = []byte(key[0])
		if len(k) > 8 {
			return "", errors.New("DES The maximum length of the encrypted salt value is 8")
		}
	}
	src := []byte(text)
	block, err := des.NewCipher(k)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	src = zeroPadding(src, bs)
	if len(src)%bs != 0 {
		return "", errors.New("What is required is an integer multiple of the size")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return hex.EncodeToString(out), nil
}

//DesDecrypt des encrypted function
//salt The parameter is the salt value of encryption and decryption, and the maximum
//length is 8. If the length is exceeded, the corresponding exception will be thrown
func (*Encryption) DesDecrypt(decrypted string, key ...string) (string, error) {
	k := []byte(defaultDesKey)
	if len(key) > 0 {
		if len(key[0]) > 8 {
			return "", errors.New("DES The maximum length of the encrypted salt value is 8")
		}
		k = []byte(key[0])
	}
	src, err := hex.DecodeString(decrypted)
	if err != nil {
		return "", err
	}
	block, err := des.NewCipher(k)
	if err != nil {
		return "", err
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out = zeroUnPadding(out)
	return string(out), nil
}
