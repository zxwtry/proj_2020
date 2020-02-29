package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/zxwtry/proj_2020/go/src/backend/crawler_today_leagal_report/constant"
)

// HttpSimpleNoToken 没有鉴权信息，返回页面信息
// url: 链接
// int: 错误码
// string: 错误原因
// string: 页面内容
func HttpSimpleNoTokenGet(url string) (int, string, string) {

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return constant.HTTP_SIMPLLE_NOTOKEN_NEW_REQUEST, fmt.Sprintf("create request [url:%s] [err:%+v]", url, err), ""
	}

	// 添加请求头
	request.Header.Add("content-type", "application/x-www-form-urlencoded")
	request.Header.Add("cache-control", "no-cache")

	//加入get参数
	// q := request.URL.Query()
	// q.Add("sort", "desc")
	// q.Add("time", tunix)
	// q.Add("key", key)
	// fmt.Println("q->", q)
	// request.URL.RawQuery = q.Encode()
	// fmt.Println("encode->", q.Encode())

	httpClient := http.Client{
		Timeout: time.Second,
	}

	resp, err := httpClient.Do(request)
	if err != nil {
		log.Println("err->", err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("err->", err)
	}

	rdata := string(data)
	fmt.Println(rdata)

	var resultdata ResultData
	json.Unmarshal([]byte(rdata), &resultdata)
	fmt.Printf("%s\n", resultdata)
	return 0, "", ""
}
