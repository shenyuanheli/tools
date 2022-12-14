/**
  @author: 1043193460@qq.com
  @date: 2022/12/13 14:00
  @note: 文件相关操作
**/
package file

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"sync"
)

//Byte转base64类型
func (zFile *ZFile) FileToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

//Base64转Byte类型
func (zFile *ZFile) Base64ToByte(baseStr string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(baseStr))
}

//Byte类型写入到指定文件
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

// 加载json文件返回map
var file_locker sync.Mutex

func (zFile *ZFile) LoadFileJson(filePath string) (map[string]interface{}, bool) {
	var conf map[string]interface{}
	file_locker.Lock()
	data, err := ioutil.ReadFile(filePath)
	file_locker.Unlock()
	if err != nil {
		fmt.Println("read json file error")
		return conf, false
	}
	err = json.Unmarshal(data, &conf)
	if err != nil {
		fmt.Println("unmarshal json file error")
		return conf, false
	}
	return conf, true
}

// 加载yaml文件返回map
func (zFile *ZFile) LoadFileYaml(filePath string) (map[string]interface{}, bool) {
	var conf map[string]interface{}
	file_locker.Lock()
	data, err := ioutil.ReadFile(filePath)
	file_locker.Unlock()
	if err != nil {
		fmt.Println("read yaml file error")
		return conf, false
	}
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		fmt.Println("unmarshal json file error")
		return conf, false
	}
	return conf, true
}
