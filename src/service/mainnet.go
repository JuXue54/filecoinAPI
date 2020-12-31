package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func addReqHeader(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Toekn", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJyZWFkIiwid3JpdGUiLCJzaWduIiwiYWRtaW4iXX0.8jNrK7zUE25ekawT-xIx4x_vTbl3GWa05PjvGkN1Wzo")
}

// 查找区块高度的服务
func ChainHead() (int, error) {
	// 安装请求体
	Post.Method = "Filecoin.ChainHead"
	Post.Params = [0]interface{}{}

	jsonStr, err := json.Marshal(Post)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonStr))
	addReqHeader(req)

	// 获取响应结果
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	m := make(map[string]interface{})
	err = json.Unmarshal(body, &m)
	if err != nil {
		panic(err)
	}
	height := int(m["result"].(map[string]interface{})["Height"].(float64))
	return height, err
}
