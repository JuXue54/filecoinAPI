package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	//"fmt"
	"math"
	"strconv"
	//"context"

	/*"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api/apistruct"*/
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
		return 0,err
	}
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonStr))
	addReqHeader(req)

	// 获取响应结果
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0,err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	m := make(map[string]interface{})
	err = json.Unmarshal(body, &m)
	if err != nil {
		return 0,err
	}
	height := int(m["result"].(map[string]interface{})["Height"].(float64))
	return height, err
}

// 获取全网算力
func TotalPower() (float64,error){
	// 安装请求体
	Post.Method = "Filecoin.StateMinerPower"
	Post.Params =[2]interface{}{"f02770",nil}

	jsonStr, err := json.Marshal(Post)
	if err != nil {
		return 0.0,err
	}
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonStr))
	addReqHeader(req)

	// 获取响应结果
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0.0,err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	m := make(map[string]interface{})
	err = json.Unmarshal(body, &m)
	if err != nil {
		return 0.0,err
	}
	temp,err:=strconv.ParseFloat(m["result"].(map[string]interface{})["TotalPower"].(map[string]interface{})["QualityAdjPower"].(string),64)
	if err!=nil{
		return 0.0,err
	}
	qualityAdjPower := temp/math.Pow(1024,6)
	return qualityAdjPower, err
}


/*func ChainHead() (int,error){
	authToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJyZWFkIiwid3JpdGUiLCJzaWduIiwiYWRtaW4iXX0.8jNrK7zUE25ekawT-xIx4x_vTbl3GWa05PjvGkN1Wzo"
	headers := http.Header{"Authorization": []string{"Bearer " + authToken}}
	addr := "3.87.90.124:1234"

	var api apistruct.FullNodeStruct
	closer, err := jsonrpc.NewMergeClient(context.Background(), "ws://"+addr+"/rpc/v0", "Filecoin", []interface{}{&api.Internal, &api.CommonStruct.Internal}, headers)
	if err != nil {
		log.Fatalf("connecting with lotus failed: %s", err)
	}
	defer closer()

       // Now you can call any API you're interested in.
	tipset, err := api.ChainHead(context.Background())
	if err != nil {
		log.Fatalf("calling chain head: %s", err)
	}
	fmt.Printf("Current chain head is: %s", tipset.String())
	return tipset.Height,err
}*/