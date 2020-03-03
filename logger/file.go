package logger

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type FileLogger struct {
	level         int
	logPath       string
	logName       string
	file          *os.File // 正常日志文件句柄
	warnfile      *os.File // 错误日志文件句柄
	LogDataChan   chan *LogData
	logSplitType  int
	logSplitSize  int64
	lastSplitHour int
}

// 3 构造文件对象 返回接口
func NewFileLogger(config map[string]string) (log LogInterface, err error) {
	logPath, ok := config["log_path"]
	if !ok {
		err = fmt.Errorf("not found log_path")
		return
	}

	logName, ok := config["log_name"]
	if !ok {
		err = fmt.Errorf("not found log_name")
		return
	}
	logLevel, ok := config["log_level"]
	if !ok {
		err = fmt.Errorf("not found log_level")
		return
	}

	logChanSize, ok := config["log_chan_size"]
	if !ok {
		logChanSize = "50000"
	}
	var logSplitType int = LogSplitTypeHour
	var logSplitSize int64
	logSplitStr, ok := config["log_split_type"]
	// 判断map 中是否有log_split_type的key，如没有则把logSplitStr 等于hour
	// 如有，则判断值的类型,如果是size 则logSplitType 为size
	// 否则为hour
	if !ok {
		logSplitStr = "hour"
	} else {
		if logSplitStr == "size" {
			logSplitSizeStr, ok := config["log_split_size"]
			if !ok {
				logSplitSizeStr = "104857600"
			}
			// 字符串转int64整型
			logSplitSize, err = strconv.ParseInt(logSplitSizeStr, 10, 64)
			if err != nil {
				logSplitSize = 104857600
			}
			// 分割类型由hour改为size分割
			logSplitType = LogSplitTypeSize
		} else {
			logSplitType = LogSplitTypeHour
		}
	}
	// 字符串转整型 设置chanSize
	chanSize, err := strconv.Atoi(logChanSize)
	if err != nil {
		chanSize = 50000
	}
	level := getLogLevel(logLevel)
	log = &FileLogger{
		level:         level,
		logPath:       logPath,
		logName:       logName,
		LogDataChan:   make(chan *LogData, chanSize),
		logSplitSize:  logSplitSize,
		logSplitType:  logSplitType,
		lastSplitHour: time.Now().Hour(),
	}
	// fmt.Println(logSplitSize, logSplitType)
	log.Init()
	return

}

// 对象实例化
func (f *FileLogger) Init() {
	filename := fmt.Sprintf("%s/%s.log", f.logPath, f.logName)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open file %s failed, err: %v", filename, err))
	}
	f.file = file

	// 写错误日志和Fatal日志的文件
	filename = fmt.Sprintf("%s/%s.log.wf", f.logPath, f.logName)
	file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open file %s failed, err: %v", filename, err))
	}
	f.warnfile = file
	// 启用新线程写日志
	go f.warnfileBackgroud()
}

// hour方式分割日志
func (f *FileLogger) splitFileHour(warnFile bool) {
	now := time.Now()
	hour := now.Hour()
	if hour == f.lastSplitHour {
		return
	}

	f.lastSplitHour = hour

	var backupFileName string
	var filename string
	if warnFile {
		backupFileName = fmt.Sprintf("%s/%s.log.wf_%04d%02%02d%02d",
			f.logPath, f.logName, now.Year(), now.Month(), now.Day(), f.lastSplitHour)
		filename = fmt.Sprintf("%s/%s.log.wf", f.logPath, f.logName)
	} else {
		backupFileName = fmt.Sprintf("%s/%s.log_%04d%02%02d%02d",
			f.logPath, f.logName, now.Year(), now.Month(), now.Day(), f.lastSplitHour)
		filename = fmt.Sprintf("%s/%s.log", f.logPath, f.logName)
	}
	file := f.file
	if warnFile {
		file = f.warnfile
	}
	file.Close()
	os.Rename(filename, backupFileName)

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return
	}

	if warnFile {
		f.warnfile = file
	} else {
		f.file = file
	}
}

// size方式分割日志
func (f *FileLogger) splitFileSize(warnFile bool) {
	file := f.file
	if warnFile {
		file = f.warnfile
	}

	startInfo, err := file.Stat()
	if err != nil {
		return
	}
	fileSize := startInfo.Size()
	if fileSize <= f.logSplitSize {
		return
	}

	var backupFileName string
	var filename string
	now := time.Now()
	if warnFile {
		backupFileName = fmt.Sprintf("%s/%s.log.wf_%04d%02d%02d%02d%02d%02d",
			f.logPath, f.logName, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
		filename = fmt.Sprintf("%s/%s.log.wf", f.logPath, f.logName)
	} else {
		backupFileName = fmt.Sprintf("%s/%s.log_%04d%02d%02d%02d%02d%02d",
			f.logPath, f.logName, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
		filename = fmt.Sprintf("%s/%s.log", f.logPath, f.logName)
	}

	file.Close()
	os.Rename(filename, backupFileName)

	file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return
	}

	if warnFile {
		f.warnfile = file
	} else {
		f.file = file
	}
}

// 检查分割日志是否为ware-fatal日志
func (f *FileLogger) checkSplitFile(warnFile bool) {
	if f.logSplitType == LogSplitTypeHour {
		f.splitFileHour(warnFile)
		return
	}
	f.splitFileSize(warnFile)
}

// 后台写日志 把chan中的日志写入到文件中
func (f *FileLogger) warnfileBackgroud() {
	for logData := range f.LogDataChan {
		var file *os.File = f.file
		if logData.WarnAndFatal {
			file = f.warnfile
		}
		f.checkSplitFile(logData.WarnAndFatal)

		fmt.Fprintf(file, "%s %s (%s:%s:%d) %s\n", logData.TimeStr,
			logData.LevelStr, logData.FileName, logData.FuncName, logData.LineNo, logData.Message)
	}
}
func (f *FileLogger) SetLevel(level int) {
	if level < LogLevelDebug || level > LogLevelFatal {
		level = LogLevelDebug
	}
	f.level = level
}

func (f *FileLogger) Debug(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	logData := writeLog(LogLevelDebug, format, args...)
	select {
	case f.LogDataChan <- logData:
	default:
	}
	// f.LogDataChan <- logData
	/*
		if f.level > LogLevelDebug {
			return
		}
		now := time.Now()
		nowStr := now.Format("2006-01-02 15:04:05.999")
		levelStr := getLevelText(LogLevelDebug)
		fileName, funcName, lineNo := GetLineInfo()
		fileName = path.Base(fileName)
		funcName = path.Base(funcName)
		msg := fmt.Sprintf(format, args...)

		fmt.Fprintf(f.file, "%s %s (%s:%s:%d) %s\n", nowStr, levelStr, fileName, funcName, lineNo, msg)
	*/
	/*
		fmt.Fprintf(f.file, nowStr)
		fmt.Fprintf(f.file, format, args...)
		fmt.Fprintln(f.file)
	*/
}

func (f *FileLogger) Trace(format string, args ...interface{}) {
	if f.level > LogLevelTrace {
		return
	}
	logData := writeLog(LogLevelTrace, format, args...)
	select {
	case f.LogDataChan <- logData:
	default:
	}
}

func (f *FileLogger) Info(format string, args ...interface{}) {
	if f.level > LogLevelInfo {
		return
	}
	logData := writeLog(LogLevelInfo, format, args...)
	select {
	case f.LogDataChan <- logData:
	default:
	}
}

func (f *FileLogger) Warn(format string, args ...interface{}) {
	if f.level > LogLevelWarn {
		return
	}
	logData := writeLog(LogLevelWarn, format, args...)
	select {
	case f.LogDataChan <- logData:
	default:
	}
}

func (f *FileLogger) Error(format string, args ...interface{}) {
	if f.level > LogLevelError {
		return
	}
	logData := writeLog(LogLevelError, format, args...)
	select {
	case f.LogDataChan <- logData:
	default:
	}
}

func (f *FileLogger) Fatal(format string, args ...interface{}) {
	if f.level > LogLevelFatal {
		return
	}
	logData := writeLog(LogLevelFatal, format, args...)
	select {
	case f.LogDataChan <- logData:
	default:
	}
}

func (f *FileLogger) Close() {
	f.file.Close()
	f.warnfile.Close()
}
