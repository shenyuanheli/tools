/**
  @author: fanyanan
  @date: 2022/6/11
  @note: //id生成工具测试
**/
package test

import (
	"fmt"
	ztool "github.com/shenyuanheli/tools"
	"testing"
)

//带有下划线的UUID
func TestRandomUUID(t *testing.T) {
	uuid, err := ztool.IdUtils.RandomUUID()
	if err != nil {
		panic(err)
		return
	}
	fmt.Printf("带下划线的UUID---------------%s\n", uuid)
}

//没有下划线的UUID
func TestSimpleUUID(t *testing.T) {
	uuid, err := ztool.IdUtils.SimpleUUID()
	if err != nil {
		panic(err)
		return
	}
	fmt.Printf("没有下划线的UUID---------------%s\n", uuid)
}

//SnowFlake算法的ID
func TestGenerateSnowflakeId(t *testing.T) {
	fmt.Printf("SnowFlake算法的ID-----------------%d\n", ztool.IdUtils.GenerateSnowflakeId())
}
