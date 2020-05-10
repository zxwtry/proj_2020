package task

import (
	"fmt"
	"time"

	"github.com/zxwtry/proj_2020/go/src/backend/crawler_today_leagal_report/constant"
	"github.com/zxwtry/proj_2020/go/src/backend/crawler_today_leagal_report/service"
	"github.com/zxwtry/proj_2020/go/src/backend/crawler_today_leagal_report/tool"
)

func TaskXinWenLianBoMp3() {
	startTimeStr := "20200301"
	endTimeStr := "20200405"
	startTime, startTimeErr := time.Parse(constant.TIME_FORMAT_YMD, startTimeStr)
	endTime, endTimeErr := time.Parse(constant.TIME_FORMAT_YMD, endTimeStr)
	if startTimeErr != nil || endTimeErr != nil {
		tool.Log("task", fmt.Sprintf("[startTimeErr:%+v] [endTimeErr:%+v]", startTimeErr, endTimeErr))
		return
	}

	for startTime.Unix() <= endTime.Unix() {
		errCode, errMsg := service.ServiceUrlXinWenLianBoMp3(startTime)
		tool.Log("task", fmt.Sprintf("xinwenlianbo [day:%s] [errCode:%d] [errMsg:%s]", startTime.Format(constant.TIME_FORMAT_YMD), errCode, errMsg))
		time.Sleep(10 * time.Second)
		startTime = startTime.AddDate(0, 0, 1)
	}

}
