/**
  @author: 1043193460@qq.com
  @date: 2022/12/13 14:00
  @note:
**/
package file

import (
	"encoding/base64"
	"os"
)

//FileToBase64 file to base64 string function
func (zFile *ZFile) FileToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

//Base64ToByte base64 string to byte
//baseStr parameter is base64 string
func (zFile *ZFile) Base64ToByte(baseStr string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(baseStr))
}

//ByteToFile The byte array is converted into a file and landed
//The parameter b is the bate array filePath file path
func (zFile *ZFile) ByteToFile(b []byte, filePath string) error {
	//Open file Create a file if there is no file
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	_, err = file.Write(b)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	return nil
}
