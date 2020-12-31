package service

import(
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 查找区块高度的服务
func ChainHead() (string,error){
	url:="http://3.87.90.124:1234/rpc/v0"
	post:="{\"id\": 0,\"jsonrpc\": \"2.0\",\"method\": \"Filecoin.ChainHead\",\"params\": []}"

	var jsonStr=[]byte(post)
	req,err:=http.NewRequest("POST",url,bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type","application/json")
	req.Header.Set("Toekn","Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJyZWFkIiwid3JpdGUiLCJzaWduIiwiYWRtaW4iXX0.8jNrK7zUE25ekawT-xIx4x_vTbl3GWa05PjvGkN1Wzo")

	client:=&http.Client{}
	resp,err:=client.Do(req)
	if err!=nil{
		panic(err)
	}
	defer resp.Body.Close()

	body,_:=ioutil.ReadAll(resp.Body)
	m:=make(map[string]interface{})
	err0:=json.Unmarshal(body, &m)
	if err0!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(m["result"].(map[string]interface{})["Height"])
	}
	height:=m["result"].(map[string]interface{})["Height"]
	return string(height),err0
}