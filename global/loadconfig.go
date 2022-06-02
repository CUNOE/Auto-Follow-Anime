package global

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Jsonrpc struct {
		RpcAddress  string `json:"rpc_address"`
		RpcProtocol string `json:"rpc_protocol"`
		RpcToken    string `json:"rpc_token"`
	} `json:"jsonrpc"`
	WsPort     string           `json:"ws_port"`
	RssAddress []RssAddressData `json:"rss_address"`
	QQGroupId  int64            `json:"qq_group_id"`
}

type RssAddressData struct {
	AppendParams json.RawMessage `json:"append_params,omitempty"`
	Rss          string          `json:"rss"`
}

func LoadConfig() {
	var conf Config

	jsonFile, err := ioutil.ReadFile("./afa/config.json")
	if err != nil {
		log.Printf(err.Error())
	}

	err = json.Unmarshal(jsonFile, &conf)
	if err != nil {
		log.Printf(err.Error())
	}

	setJsonRpc(conf)

	RssAddress = conf.RssAddress
	JsonRpcToken = conf.Jsonrpc.RpcToken
	WsPort = conf.WsPort
	QQGroupId = conf.QQGroupId

}

func setJsonRpc(conf Config) {
	JsonRpcServer = conf.Jsonrpc.RpcProtocol + "://" + conf.Jsonrpc.RpcAddress
	log.Printf("setJsonrpcServer: %v", JsonRpcServer)
	return
}
