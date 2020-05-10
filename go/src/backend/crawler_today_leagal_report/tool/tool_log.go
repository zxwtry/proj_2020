package tool

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/zxwtry/proj_2020/go/src/backend/backend_comm_constant"
	"github.com/zxwtry/proj_2020/go/src/backend/backend_comm_func"
	"github.com/zxwtry/proj_2020/go/src/backend/crawler_today_leagal_report/comm_constant"
)

// 自己实现，不怎么高效的写日至api
const LOG_FORMAT_FILE_INFO = "%s [%s:%d:%+v] %s\r\n"
const LOG_FORMAT_ORIGIN = "%s %s\r\n"

var globalLogLock sync.Mutex

type LogNameLock struct {
	LogName     string
	FileHandler *os.File
	Lock        *sync.Mutex
}

var logMap map[string]LogNameLock = make(map[string]LogNameLock)

func Log(logName, logStr string) {
	// 获取logNameLock对象全局加锁
	logNameLock, logNameLockErr := getLogNameLock(logName)
	if logNameLockErr != nil {
		fmt.Fprintf(os.Stderr, "log err [logName:%s] [logNameLockErr:%+v]", logName, logNameLockErr)
		return
	}

	fileFunc, fileName, fileLine, ok := runtime.Caller(1)
	strYmdHis := time.Now().Format(comm_constant.TIME_FORMAT_YMDHIS)
	if ok {
		logStr = fmt.Sprintf(LOG_FORMAT_FILE_INFO, strYmdHis, fileName, fileLine, fileFunc, logStr)
	} else {
		logStr = fmt.Sprintf(LOG_FORMAT_ORIGIN, strYmdHis, logStr)
	}

	logNameLock.Lock.Lock()
	defer logNameLock.Lock.Unlock()

	logNameLock.FileHandler.WriteString(logStr)
}

func getLogNameLock(logName string) (*LogNameLock, error) {
	globalLogLock.Lock()
	defer globalLogLock.Unlock()
	if logNameLock, ok := logMap[logName]; ok {
		return &logNameLock, nil
	}
	strYmdHis := time.Now().Format(comm_constant.TIME_FORMAT_YMDH_ONLY)

	var fileNameForm string
	if backend_comm_func.CheckZxwPcEnv() {
		fileNameForm = backend_comm_constant.LOG_FILE_ZXW_PC
	} else {
		fileNameForm = backend_comm_constant.LOG_FILE_NOT_ZXW_PC
	}

	fileName := fmt.Sprintf(fileNameForm, strYmdHis, logName)
	fileHandler, fileHandlerErr := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if fileHandlerErr != nil {
		return nil, fileHandlerErr
	}
	var logLock sync.Mutex
	logNameLock := LogNameLock{
		LogName:     logName,
		FileHandler: fileHandler,
		Lock:        &logLock,
	}
	logMap[logName] = logNameLock
	return &logNameLock, nil
}
