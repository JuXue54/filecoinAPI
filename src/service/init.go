package service

import (
	"zerone/filecoinAPI/src/model"
)

var URL string
var Post model.ReqBody

func init() {
	URL = "http://3.87.90.124:1234/rpc/v0"
	Post = model.ReqBody{Id: 0, Jsonrpc: "2.0"}
}
