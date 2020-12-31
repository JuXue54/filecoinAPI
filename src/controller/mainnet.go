package controller

import (
	"encoding/json"
	"net/http"

	"zerone/filecoinAPI/src/model"
	"zerone/filecoinAPI/src/service"
)

// 拿到当前的区块头
func chainHead() int {
	height, err := service.ChainHead()
	if err == nil {
		return height
	} else {
		return -1
	}
}

// 返回主网数据
func Mainnet(w http.ResponseWriter, req *http.Request) {
	//fmt.Println(*req)
	w.Header().Set("Content-Type", "application/json")
	h := chainHead()
	resp := model.RespBody{Height: h}
	b, err := json.Marshal(resp)
	if err == nil && h > -1 {
		w.Write(b)
	} else {
		w.Write([]byte("some error happen 505"))
	}
}
