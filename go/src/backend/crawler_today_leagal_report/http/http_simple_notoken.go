package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/zxwtry/proj_2020/go/src/backend/crawler_today_leagal_report/constant"
)

var globalClient http.Client = http.Client{
	Timeout: 30 * time.Second,
}

// HttpSimpleNoToken 没有鉴权信息，返回页面信息
// dataUrl: 链接
// int: 错误码
// string: 错误原因
// string: 页面内容
func HttpSimpleNoTokenGet(dataUrl string) (int32, string, string) {

	request, err := http.NewRequest("GET", dataUrl, nil)
	if err != nil {
		return constant.HTTP_SIMPLE_NOTOKEN_NEW_REQUEST_ERROR, fmt.Sprintf("create request [dataUrl:%s] [err:%+v]", dataUrl, err), ""
	}

	// 添加请求头
	request.Header.Add("content-type", "application/x-www-form-urlencoded")
	request.Header.Add("cache-control", "no-cache")

	//加入get参数
	// q := request.url.Query()
	// q.Add("sort", "desc")
	// q.Add("time", tunix)
	// q.Add("key", key)
	// fmt.Println("q->", q)
	// request.url.RawQuery = q.Encode()
	// fmt.Println("encode->", q.Encode())

	resp, err := globalClient.Do(request)
	if err != nil {
		return constant.HTTP_SIMPLE_NOTOKEN_DO_REQUEST_ERROR, fmt.Sprintf("http client do request err [request:%+v] [err:%+v]", request, err), ""
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return constant.HTTP_SIMPLE_NOTOKEN_READ_RESPONSE_ERROR, fmt.Sprintf("io read all [resp:%+v] [err:%+v]", resp, err), ""
	}
	return 0, "", string(data)
}
