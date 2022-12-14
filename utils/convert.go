/**
  @author: 1043193460@qq.com
  @date: 2022/12/14 12:01
  @note:
**/
package utils

import "fmt"

//递归转换map的key
func (c *UtilsLogic) MapConvert(m map[interface{}]interface{}) map[string]interface{} {
	res := map[string]interface{}{}
	for k, v := range m {
		switch v2 := v.(type) {
		case map[interface{}]interface{}:
			res[fmt.Sprint(k)] = c.MapConvert(v2)
		default:
			res[fmt.Sprint(k)] = v
		}
	}
	return res
}
