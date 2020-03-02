package service

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/zxwtry/proj_2020/go/src/backend/crawler_today_leagal_report/constant"
	"github.com/zxwtry/proj_2020/go/src/backend/crawler_today_leagal_report/http"
)

// 返回参数
// int: 错误码
// string: 错误信息
// string: mp3文件位置
// string: lrc文件位置
func ServcieUrlMusicNete(dataUrl string) (int32, string, string, string) {
	qMarkIndex := strings.Index(dataUrl, "?")
	dataUrlParse, dataUrlErr := url.ParseQuery(dataUrl[qMarkIndex+1:])
	if dataUrlErr != nil {
		return constant.FUNCTION_PARAM_ERROR, fmt.Sprintf("url parse error [dataUrl:%s] [dataUrlErr:%+v]", dataUrl, dataUrlErr), "", ""
	}
	dataUrlId := dataUrlParse.Get("id")
	if len(dataUrlId) == 0 {
		return constant.FUNCTION_PARAM_ERROR, fmt.Sprintf("data url id empty [dataUrl:%s] [dataUrlParse:%+v]", dataUrl, dataUrlParse), "", ""
	}
	dataUrl = "https://music.163.com/song?id=" + dataUrlId
	urlErrCode, urlErrMsg, urlData := http.HttpSimpleNoTokenGet(dataUrl)
	return urlErrCode, urlErrMsg, urlData, ""
	// return 0, "", "", ""
}
