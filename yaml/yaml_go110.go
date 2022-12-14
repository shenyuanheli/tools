/**
  @author: 1043193460@qq.com
  @date: 2022/12/14 12:02
  @note: 这个文件包含的变化只是兼容1.10和以后。
**/
package yaml

import "encoding/json"

//DisableUnknownFields将JSON解码器配置为未知时出错
//字段会出现，而不是默认情况下删除它们。
func DisallowUnknownFields(d *json.Decoder) *json.Decoder {
	d.DisallowUnknownFields()
	return d
}
