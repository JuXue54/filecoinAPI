package util

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// RequestJSON 直接获取 请求参数是 JSON 的字符串
func RequestJSON(req *http.Request) string {
	if req != nil && req.Body != nil {
		result, err := ioutil.ReadAll(req.Body)

		if err == nil {
			return string(result)
		}
	}
	return ""
}

// SendHTTPPost 发起 HTTP POST 请求
func SendHTTPPost(url string, param string, mime string) string {
	resp, err := http.Post(url, mime, strings.NewReader(param))
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(body)
}

// SendHTTPDo 发起 HTTP Do 详细请求
func SendHTTPDo(url string, method string, params string, mime string, header map[string]string, cookie string) string {
	req, err := http.NewRequest(method, url, strings.NewReader(params))
	if err != nil {
		return ""
	}
	req.Header.Set("Content-Type", mime)
	req.Header.Set("Cookie", cookie)

	if header != nil && len(header) > 0 {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return string(body)
}
