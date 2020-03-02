package tool

import (
	"fmt"
	"regexp"

	"github.com/zxwtry/proj_2020/go/src/backend/crawler_today_leagal_report/constant"
)

// ToolRegexSimpleFindAll 正则表达简单查找全部
func ToolRegexSimpleFindAll(regPatternPrefix, regPatternSuffix, data string) (int32, string, []string) {
	findString := make([]string, 0, 10)
	regPattern := regPatternPrefix + "(.+)" + regPatternSuffix
	reg, regErr := regexp.Compile(regPattern)
	if regErr != nil {
		Log("ToolRegexSimpleFindAll", fmt.Sprintf("reg pattern compile error [regPattern:%s] [regErr:%+v]", regPattern, regErr))
		return constant.FUNCTION_PARAM_ERROR, fmt.Sprintf("tool regex [regPatternPrefix:%s] [regPatternSuffix:%s]", regPatternPrefix, regPatternSuffix), findString
	}
	arrPlayUrlBs := reg.FindAll([]byte(data), -1)
	for _, playUrlBs := range arrPlayUrlBs {
		if len(playUrlBs) > len(regPattern) {
			playUrl := string(playUrlBs[len(regPatternPrefix) : len(playUrlBs)-len(regPatternSuffix)])
			findString = append(findString, playUrl)
		}
	}
	return 0, "", findString
}
