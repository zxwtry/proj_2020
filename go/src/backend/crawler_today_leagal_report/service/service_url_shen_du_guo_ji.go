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

func ServiceUrlShenDuGuoJi() (int32, string) {
	videoSetUrl := "http://api.cntv.cn/NewVideo/getVideoListByColumn?id=TOPC1451540709098112&sort=desc&serviceId=tvcctv&mode=0&n=25&p=1"

	videoSetErrCode, videoSetErrMsg, videoSetData := http.HttpSimpleNoTokenGet(videoSetUrl)
	if videoSetErrCode != 0 {
		tool.Log("shenduguoji", fmt.Sprintf("http simple notoken get [videoSetUrl:%s] [videoSetErrCode:%d] [videoSetErrMsg:%s]", videoSetUrl, videoSetErrCode, videoSetErrMsg))
		return videoSetErrCode, videoSetErrMsg
	}

	shenDuGuoJiVideoSet := constant.ShenDuGuoJiVideoSet{}
	shenDuGuoJiVideoSetErr := json.Unmarshal([]byte(videoSetData), &shenDuGuoJiVideoSet)
	if shenDuGuoJiVideoSetErr != nil {
		tool.Log("shenduguoji", fmt.Sprintf("http simple notoken get [videoSetUrl:%s] [videoSetErrCode:%d] [videoSetErrMsg:%s] [videoSetData:%s]", videoSetUrl, videoSetErrCode, videoSetErrMsg, videoSetData))
		return constant.FUNCTION_JSON_UNMARSHAL_ERROR, shenDuGuoJiVideoSetErr.Error()
	}

	for _, shenDuGuoJiVideo := range shenDuGuoJiVideoSet.Data.List {
		time.Sleep(5 * time.Second)
		mp3FilePath := fmt.Sprintf(constant.FILE_PATH_SHEN_DU_GUO_JI, strings.Replace(shenDuGuoJiVideo.Title, " ", "-", -1))
		mp3FilePath = strings.ReplaceAll(mp3FilePath, "《", "")
		mp3FilePath = strings.ReplaceAll(mp3FilePath, "》", "")
		_, fileExistErr := os.Stat(mp3FilePath)
		if fileExistErr == nil {
			// 代表文件已经存在
			continue
		}

		videoInfoUrl := fmt.Sprintf("http://vdn.apps.cntv.cn/api/getHttpVideoInfo.do?url=%s&pid=%s", shenDuGuoJiVideo.Url, shenDuGuoJiVideo.Guid)

		// 下载播放页面
		videoInfoErrCode, videoInfoErrMsg, videoInfoData := http.HttpSimpleNoTokenGet(videoInfoUrl)
		if videoInfoErrCode != 0 || len(videoInfoData) == 0 {
			tool.Log("shenduguoji", fmt.Sprintf("no token get [videoInfoUrl:%s] [videoInfoErrCode:%d] [videoInfoErrMsg:%s]", videoInfoUrl, videoInfoErrCode, videoInfoErrMsg))
			continue
		}
		xinWenLianBoVideoInfo := constant.XinWenLianBoVideoInfo{}
		videoInfoUnmarshalErr := json.Unmarshal([]byte(videoInfoData), &xinWenLianBoVideoInfo)
		if videoInfoUnmarshalErr != nil {
			tool.Log("shenduguoji", fmt.Sprintf("videoInfo Unmarshal error [videoInfoData:%s] [videoInfoUnmarshalErr:%d]", videoInfoData, videoInfoUnmarshalErr))
			continue
		}
		tool.Log("shenduguoji", fmt.Sprintf("xinWenLianBoVideoInfo:%+v", xinWenLianBoVideoInfo))
		fmt.Printf("%+v\n", xinWenLianBoVideoInfo)

		mp3Url := strings.Replace(xinWenLianBoVideoInfo.Manifest.AudioMp3, "main.m3u8", "16.mp3", 1)
		// 下载播放页面
		mp3ErrCode, mp3ErrMsg, mp3Data := http.HttpSimpleNoTokenGet(mp3Url)
		if mp3ErrCode != 0 || len(mp3Data) == 0 {
			tool.Log("xinwenlianbo", fmt.Sprintf("no token get [mp3Url:%s] [mp3ErrCode:%d] [mp3ErrMsg:%s]", mp3Url, mp3ErrCode, mp3ErrMsg))
			continue
		}

		mp3WriteErrCode, mp3WriteErrMsg := tool.ToolFileWrite(mp3FilePath, []byte(mp3Data), os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
		if mp3WriteErrCode != 0 {
			tool.Log("xinwenlianbo", fmt.Sprintf("mp3 write file [mp3FilePath:%s] [mp3WriteErrCode:%d] [mp3WriteErrMsg:%s]", mp3FilePath, mp3WriteErrCode, mp3WriteErrMsg))
			continue
		}
	}
	return 0, ""
}
