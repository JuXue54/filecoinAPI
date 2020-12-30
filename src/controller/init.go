package controller

import(
	"gjson"
)

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


// 暴露filecoin主网参数给外界
func mainnet(req *http.Request){

}