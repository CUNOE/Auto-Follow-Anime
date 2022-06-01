package global

import (
	"autoFollowAnime/model"
	"gorm.io/gorm"
)

var JsonRpcModel = &model.JsonrpcModel{
	Jsonrpc: "2.0",
	Method:  "aria2.addUri",
	Id:      "Y3Vua3ljaGVuZw",
	Params:  nil,
}
var JsonRpcServer string
var JsonRpcToken string

var RssAddress []RssAddressData

var DB *gorm.DB
