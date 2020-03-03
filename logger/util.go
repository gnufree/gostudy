package logger

import (
	"fmt"
	"path"
	"runtime"
	"time"
)

// 异步写日志对象

type LogData struct {
	Message      string
	TimeStr      string
	LevelStr     string
	FileName     string
	FuncName     string
	LineNo       int
	WarnAndFatal bool
}

// 获取打日志的文件名、函数名及行号
func GetLineInfo() (fileName string, funcName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(4) // 栈深度 0 1 2 3
	if ok {
		fileName = file
		funcName = runtime.FuncForPC(pc).Name()
		lineNo = line
	}
	return
}

// 写日志到对象
func writeLog(level int, format string, args ...interface{}) *LogData {
	now := time.Now()
	nowStr := now.Format("2006-01-02 15:04:05.999")
	levelStr := getLevelText(level)
	fileName, funcName, lineNo := GetLineInfo()
	fileName = path.Base(fileName)
	funcName = path.Base(funcName)
	msg := fmt.Sprintf(format, args...)

	logData := &LogData{
		Message:      msg,
		TimeStr:      nowStr,
		LevelStr:     levelStr,
		FileName:     fileName,
		FuncName:     funcName,
		LineNo:       lineNo,
		WarnAndFatal: false,
	}

	if level == LogLevelWarn || level == LogLevelError || level == LogLevelFatal {
		logData.WarnAndFatal = true
	}

	return logData

	// fmt.Fprintf(file, "%s %s (%s:%s:%d) %s\n", nowStr, levelStr, fileName, funcName, lineNo, msg)
}
