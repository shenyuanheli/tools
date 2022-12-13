/**
  @author: 1043193460@qq.com
  @date: 2022/12/13 14:32
  @note:
**/
package encryption

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

//Md5Check md5 Check method
func (e *Encryption) Md5Check(content, encrypted string) bool {
	return strings.EqualFold(e.Md5Encode(content), encrypted)
}

//Md5Encode md5 Signature function
func (e *Encryption) Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
