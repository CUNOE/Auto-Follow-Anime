package model

type JsonrpcModel struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Id      string        `json:"id"`
	Params  []interface{} `json:"params"`
}

type JsonrpcBack struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  string `json:"result,omitempty"`
	Id      string `json:"id"`
	Error   string `json:"error,omitempty"`
}
