/**
  @author: 1043193460@qq.com
  @date: 2022/12/13 14:32
  @note:
**/
package tools

import (
	"github.com/shenyuanheli/tools/date"
	"github.com/shenyuanheli/tools/encryption"
	"github.com/shenyuanheli/tools/file"
	"github.com/shenyuanheli/tools/id"
	"github.com/shenyuanheli/tools/idcard"
	"github.com/shenyuanheli/tools/net"
	"github.com/shenyuanheli/tools/str"
	"github.com/shenyuanheli/tools/utils"
)

var (
	DateUtils       date.DateTime
	IdUtils         id.Id
	EncodeUtils     encryption.Encryption
	IDCardUtils     idcard.Idcard
	StrUtils        str.Str
	FileUtils       file.ZFile
	EncryptionUtils encryption.Encryption
	HttpUtils       net.Client
	MyUtils         utils.UtilsLogic
)
