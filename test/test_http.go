/**
  @author: fanyanan
  @date: 2022/6/11
  @note: //http客户端测试
**/
package test

import (
	"fmt"
	ztool "github.com/shenyuanheli/tools"
	"testing"
)

func TestHttp(t *testing.T) {
	get, err := ztool.HttpUtils.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(get)
}
