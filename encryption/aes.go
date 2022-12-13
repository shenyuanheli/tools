/**
  @author: 1043193460@qq.com
  @date: 2022/12/13 14:32
  @note:
**/
package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

//AesEncrypt function
//original The original password, salt is an optional variable.
//If the salt exists, the passed variable is used.
//If it does not exist, the system default salt value is used.
func (e *Encryption) AesEncrypt(original string, key ...string) (string, error) {
	var saltValue = defaultAesKey
	if len(key) > 0 {
		saltValue = key[0]
	}

	// Convert to byte array
	origData := []byte(original)
	k := []byte(saltValue)
	if len(k) != 16 && len(k) != 24 && len(k) != 32 {
		return "", errors.New("The length of the key should be 16 or 24 or 32")
	}
	// group key
	block, _ := aes.NewCipher(k)
	// Get the length of the key block
	blockSize := block.BlockSize()
	// Completion code
	origData = pKCS7Padding(origData, blockSize)
	// encryption mode
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// create array
	crated := make([]byte, len(origData))
	// encryption
	blockMode.CryptBlocks(crated, origData)
	return base64.StdEncoding.EncodeToString(crated), nil

}

//AesDecrypt function
//original The original password, salt is an optional variable.
//If the salt exists, the passed variable is used.
//If it does not exist, the system default salt value is used.
func (e *Encryption) AesDecrypt(crated string, salt ...string) (string, error) {
	var saltValue = defaultAesKey
	if len(salt) > 0 {
		saltValue = salt[0]
	}
	// Convert to byte array
	cratedByte, _ := base64.StdEncoding.DecodeString(crated)
	k := []byte(saltValue)
	if len(k) != 16 && len(k) != 24 && len(k) != 32 {
		return "", errors.New("The length of the key should be 16 or 24 or 32")
	}
	// group key
	block, _ := aes.NewCipher(k)
	// Get the length of the key block
	blockSize := block.BlockSize()
	// encryption mode
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// create array
	orig := make([]byte, len(cratedByte))
	// decrypt
	blockMode.CryptBlocks(orig, cratedByte)
	// to complete the code
	orig = pKCS7UnPadding(orig)
	return string(orig), nil
}
