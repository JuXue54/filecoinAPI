package controller

import(
	"github.com/JuXue54/filecoinAPI/service"
	"fmt"
	"net/http"
)

type Resp struct{
	height string
}

// requestJSON 把请求参数转成 JSON 对象
func requestJSON(req *http.Request) gjson.Result {
	jsonString := requestJSONString(req)
	jsonResult := gjson.Result{}
	if jsonString != "" {
		jsonResult = gjson.Parse(jsonString)
	}

	return jsonResult
}

// requestJSONString 把请求参数转成 JSON 字符串
func requestJSONString(req *http.Request) string {
	return util.RequestJSON(req)
}

// 拿到当前的区块头
func chainHead(){
	height,err:=service.ChainHead()
	if err==nil{
		return height
	}else{
		return 0
	}
}

// 返回主网数据
func mainnet(w http.ResponseWriter, req *http.Request){
	fmt.Println(*req)
	w.Header().set("Content-Type","application/json")
	h:=chainHead()
	var resp Resp=Resp{height:h}
	b,err:=json.Marshal(resp)
	if err!=nil{
		json.NewEncoder(w).Encode(b)
	}else{
		json.NewEncoder(w).Encode(err)
	}
}
