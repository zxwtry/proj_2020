package task

import (
	"fmt"

	"github.com/zxwtry/proj_2020/go/src/backend/crawler_today_leagal_report/service"
	"github.com/zxwtry/proj_2020/go/src/backend/crawler_today_leagal_report/tool"
)

func TaskShenDuGuoJi() {
	errCode, errMsg := service.ServiceUrlShenDuGuoJi()
	if errCode != 0 {
		tool.Log("shenduguoji", fmt.Sprintf("service url [errCode:%d] [errmsg:%s]", errCode, errMsg))
	}
}
