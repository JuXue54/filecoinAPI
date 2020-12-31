package model

type ReqBody struct {
	Id      int         `json:"id"`
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

type RespBody struct {
	Height int `json:"height"`
}
