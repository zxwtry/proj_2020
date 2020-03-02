package service

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/zxwtry/proj_2020/go/src/backend/crawler_today_leagal_report/constant"
	"github.com/zxwtry/proj_2020/go/src/backend/crawler_today_leagal_report/http"
	"github.com/zxwtry/proj_2020/go/src/backend/crawler_today_leagal_report/tool"
)

func ServiceUrlXinWenLianBoMp3(d time.Time) (int32, string) {
	strYmd := d.Format(constant.TIME_FORMAT_YMD)
	dataUrl := fmt.Sprintf("http://tv.cctv.com/lm/xwlb/day/%s.shtml", strYmd)
	getErrCode, getErrMsg, getData := http.HttpSimpleNoTokenGet(dataUrl)
	if getErrCode != 0 {
		tool.Log("xinwenlianbo", fmt.Sprintf("no token get [dataUrl:%s] [getErrCode:%d] [getErrMsg:%s]", dataUrl, getErrCode, getErrMsg))
		return getErrCode, getErrMsg
	}

	// 正则匹配链接
	urlPatternPrefix := "<a href=\""
	urlPatternSuffix := "\" target=\"_blank\">"
	urlErrCode, urlErrMsg, urlList := tool.ToolRegexSimpleFindAll(urlPatternPrefix, urlPatternSuffix, getData)
	if urlErrCode != 0 || len(urlList) == 0 {
		tool.Log("xinwenlianbo", fmt.Sprintf("regex url data err [urlPatternPrefix:%s] [urlPatternSuffix:%s] [getData:%s] [urlErrCode:%d] [urlErrMsg:%s]", urlPatternPrefix, urlPatternSuffix, getData, urlErrCode, urlErrMsg))
		return urlErrCode, urlErrMsg
	}
	playUrl := urlList[0] // 一般取第一个

	// 下载播放页面
	playErrCode, playErrMsg, playData := http.HttpSimpleNoTokenGet(playUrl)
	if playErrCode != 0 {
		tool.Log("xinwenlianbo", fmt.Sprintf("no token get [playUrl:%s] [playErrCode:%d] [playErrMsg:%s]", playUrl, playErrCode, playErrMsg))
		return playErrCode, playErrMsg
	}

	// 正则匹配 pid
	pidPatternPrefix := "var guid_Ad_VideoCode = \""
	pidPatternSuffix := "\";"
	pidErrCode, pidErrMsg, pidList := tool.ToolRegexSimpleFindAll(pidPatternPrefix, pidPatternSuffix, playData)
	if pidErrCode != 0 || len(pidList) == 0 {
		tool.Log("xinwenlianbo", fmt.Sprintf("regex url data err [pidPatternPrefix:%s] [pidPatternSuffix:%s] [playData:%s] [pidErrCode:%d] [pidErrMsg:%s]", pidPatternPrefix, pidPatternSuffix, playData, pidErrCode, pidErrMsg))
		return pidErrCode, pidErrMsg
	}

	videoInfoUrl := fmt.Sprintf("http://vdn.apps.cntv.cn/api/getHttpVideoInfo.do?url=%s&pid=%s", playUrl, pidList[0])

	// 下载播放页面
	videoInfoErrCode, videoInfoErrMsg, videoInfoData := http.HttpSimpleNoTokenGet(videoInfoUrl)
	if videoInfoErrCode != 0 || len(videoInfoData) == 0 {
		tool.Log("xinwenlianbo", fmt.Sprintf("no token get [videoInfoUrl:%s] [videoInfoErrCode:%d] [videoInfoErrMsg:%s]", videoInfoUrl, videoInfoErrCode, videoInfoErrMsg))
		return videoInfoErrCode, videoInfoErrMsg
	}
	xinWenLianBoVideoInfo := constant.XinWenLianBoVideoInfo{}
	videoInfoUnmarshalErr := json.Unmarshal([]byte(videoInfoData), &xinWenLianBoVideoInfo)
	if videoInfoUnmarshalErr != nil {
		tool.Log("xinwenlianbo", fmt.Sprintf("videoInfo Unmarshal error [videoInfoData:%s] [videoInfoUnmarshalErr:%d]", videoInfoData, videoInfoUnmarshalErr))
		return constant.FUNCTION_JSON_UNMARSHAL_ERROR, videoInfoUnmarshalErr.Error()
	}

	mp3Url := strings.Replace(xinWenLianBoVideoInfo.Manifest.AudioMp3, "main.m3u8", "16.mp3", 1)
	// 下载播放页面
	mp3ErrCode, mp3ErrMsg, mp3Data := http.HttpSimpleNoTokenGet(mp3Url)
	if mp3ErrCode != 0 || len(mp3Data) == 0 {
		tool.Log("xinwenlianbo", fmt.Sprintf("no token get [mp3Url:%s] [mp3ErrCode:%d] [mp3ErrMsg:%s]", mp3Url, mp3ErrCode, mp3ErrMsg))
		return mp3ErrCode, mp3ErrMsg
	}

	mp3FilePath := fmt.Sprintf(constant.FILE_PATH_XIN_WEN_LIAN_BO, strYmd)
	mp3WriteErrCode, mp3WriteErrMsg := tool.ToolFileWrite(mp3FilePath, []byte(mp3Data), os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if mp3WriteErrCode != 0 {
		tool.Log("xinwenlianbo", fmt.Sprintf("mp3 write file [mp3FilePath:%s] [mp3WriteErrCode:%d] [mp3WriteErrMsg:%s]", mp3FilePath, mp3WriteErrCode, mp3WriteErrMsg))
		return mp3WriteErrCode, mp3WriteErrMsg
	}
	return 0, ""
}
