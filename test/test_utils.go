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
	dataMap := map[interface{}]interface{}{
		"code": 200,
		"msg":  "success",
		"data": nil,
	}
	data := tools.MyUtils.MapConvert(dataMap)
	fmt.Println(data)
}
