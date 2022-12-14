/**
  @author: 1043193460@qq.com
  @date: 2022/12/13 15:35
  @note:
**/
package test

import (
	"fmt"
	ztool "github.com/shenyuanheli/tools"
)

func main() {
	//生成密钥对，保存到文件
	data, _ := ztool.FileUtils.LoadFileYaml("file/aa.yaml")
	fmt.Println(data)
}
