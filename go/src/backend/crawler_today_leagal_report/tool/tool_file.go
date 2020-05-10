package tool

import (
	"fmt"
	"os"

	"github.com/zxwtry/proj_2020/go/src/backend/crawler_today_leagal_report/comm_constant"
)

func ToolFileWrite(filePath string, data []byte, openFileFlag int, openFileMod os.FileMode) (int32, string) {
	fileHandler, fileHandlerErr := os.OpenFile(filePath, openFileFlag, openFileMod)
	if fileHandlerErr != nil {
		return comm_constant.FUNCTION_PARAM_ERROR, fmt.Sprintf("open file error [filePath:%s] [openFileFlag:%d] [openFileMod:%d]", filePath, openFileFlag, openFileMod)
	}
	defer fileHandler.Close()

	writeInt, writeErr := fileHandler.Write(data)
	if writeErr != nil || writeInt != len(data) {
		return comm_constant.FUNCTION_EXEC_ERROR, fmt.Sprintf("write file error [filePath:%s] [len(data):%d] [writeInt:%d] [writeErr:%+v]", filePath, len(data), writeInt, writeErr)
	}
	return 0, ""
}
