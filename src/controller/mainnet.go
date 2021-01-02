package controller

import (
	"encoding/json"
	"net/http"

	"zerone/filecoinAPI/src/model"
	"zerone/filecoinAPI/src/service"
)

// 拿到当前的区块头
func chainHead() (int,error) {
	return service.ChainHead()
}

// 拿到当前的全网算力
func totalPower() (float64,error){
	return service.TotalPower()
}

// 返回主网数据
func Mainnet(w http.ResponseWriter, req *http.Request) {
	// 封装返回数据并做错误处理
	height,err := chainHead()
	power,err:=totalPower()
	resp := model.RespBody{
		Height: height,
		TotalPower: power,
	}
	b, err := json.Marshal(resp)
	// 返回json
	w.Header().Set("Content-Type", "application/json")
	if err == nil {
		w.Write(b)
	} else {
		w.Write([]byte("server error 505: "+err.Error()))
	}
}
