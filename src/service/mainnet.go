package main

import(
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)


func chainHead() string{
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
	fmt.Println(string(body))
	return string(body)
}