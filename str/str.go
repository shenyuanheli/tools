/**
  @author: 1043193460@qq.com
  @date: 2022/12/13 14:32
  @note:
**/
package str

import (
	"encoding/hex"
	"path"
	"strings"
)

// SnakeString HUMP TURN SNAKE
func (str *Str) SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or convert upper and lower case through ASCII code
		// 65-90（A-Z），97-122（a-z）
		//Judging if the letters are uppercase A-Z, splicing a _ in front
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	//To Lower converts uppercase letters to lowercase uniformly
	return strings.ToLower(string(data[:]))
}

// CamelString Snake turn hump
func (str *Str) CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

//HasBlank will also count invisible characters as empty
func (str *Str) HasBlank(s string) bool {
	s = strings.Replace(s, " ", "", -1)
	return len(s) == 0
}

//HasEmpty Only judge whether it is ("")
func (str *Str) HasEmpty(s string) bool {
	return len(s) == 0
}

// RemoveSuffix get filename That is, remove the file suffix
func (str *Str) RemoveSuffix(s string) string {
	filenameWithSuffix := path.Base(s)
	fileSuffix := path.Ext(filenameWithSuffix)
	filenameOnly := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	return filenameOnly
}

//RemovePrefix Get file extension
func (str *Str) RemovePrefix(s string) string {
	filenameWithSuffix := path.Base(s)
	return path.Ext(filenameWithSuffix)
}

// EncodeHexString String to hexadecimal
func (str *Str) EncodeHexString(s string) string {
	byteData := []byte(s)
	return hex.EncodeToString(byteData)
}

// DecodeHexString hexadecimal to string
func (str *Str) DecodeHexString(s string) (string, error) {
	hexData, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
		return "", err
	}
	return string(hexData), nil
}
