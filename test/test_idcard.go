/**
  @author: fanyanan
  @date: 2022/6/11
  @note: //中华人民共和国身份证校验
**/
package test

import (
	"fmt"
	ztool "github.com/shenyuanheli/tools"
	"testing"
)

//进行身份证校验
func TestCheck(t *testing.T) {
	check, err := ztool.IDCardUtils.Check("130109198904241210")
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(check)
}

//获取身份证所属地
func TestGetProvinceByIdCard(t *testing.T) {
	fmt.Println(ztool.IDCardUtils.GetProvinceByIdCard("130109198904241210"))
}

//func TestCard(t *testing.T) {
//	id:="130109198904241210"
//	id15:="130109198904210"
//	//进行全局校验身份证合法性
//	check, err := ztool.IDCardUtils.Check(id)
//	//获取所在地
//	card := ztool.IDCardUtils.GetProvinceByIdCard(id)
//	//校验身份证号是否是18或者15位
//	validCard, err := ztool.IDCardUtils.IsValidCard(id15)
//	//15位转18位
//	to18 := ztool.IDCardUtils.Convert15To18(id15)
//	//获取生日
//	idCard, err := ztool.IDCardUtils.GetBirthByIdCard(id)
//	//获取年龄
//	age := ztool.IDCardUtils.GetAgeByIdCard(id)
//	//出生年份
//	year := ztool.IDCardUtils.GetYearByIdCard(id)
//	//月份
//	month := ztool.IDCardUtils.GetMonthByIdCard(id)
//	//日
//	day := ztool.IDCardUtils.GetDayByIdCard(id)
//	//获取性别 0女1男3未知
//	sex := ztool.IDCardUtils.GetSex(id)
//}
