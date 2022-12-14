/**
  @author: 1043193460@qq.com
  @date: 2022/12/13 15:35
  @note: Yaml和Json互转
**/
package test

import (
	"fmt"
	"github.com/shenyuanheli/tools"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	//将对象编组为JSON，然后将JSON转换为YAML并返回YAML
	p := Person{"John", 30}
	y, err := tools.YamlUtils.Marshal(p)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(y))

	var p2 Person
	err = tools.YamlUtils.Unmarshal(y, &p2)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(p2)

	//Json2Yaml
	j := []byte(`{"name": "John", "age": 30}`)
	y, err = tools.YamlUtils.JSONToYAML(j)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(y))

	//Yaml2Json
	j2, err := tools.YamlUtils.YAMLToJSON(y)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(j2))
}
