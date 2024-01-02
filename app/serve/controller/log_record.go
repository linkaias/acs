package controller

import (
	"APICallerStats/model"
	"fmt"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strings"
	"time"
)

const (
	LogLevelInfo = iota
	LogLevelError
)

var logClient *logrus.Logger

type logData struct {
	LogLevel int
	title    string
	content  string
	v        interface{}
	Track    string
}

var waitWrite = make(chan *logData, 10000)

func initLoad(env *model.Env) {
	//禁止logrus的输出
	logClient = logrus.New()
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	logClient.Out = src
	logClient.SetLevel(logrus.DebugLevel)
	apiLogInfoPath := strings.Trim(env.LogFilePath, "/") + "/info.log"
	apiLogErrorPath := strings.Trim(env.LogFilePath, "/") + "/error.log"

	if _, err := os.Stat(strings.Trim(env.LogFilePath, "/")); os.IsNotExist(err) {
		err = os.MkdirAll(strings.Trim(env.LogFilePath, "/"), os.ModePerm)
		if err != nil {
			panic("创建日志文件失败，请检查权限：" + err.Error())
		}
	}

	maxAge := time.Duration(env.LogFileMaxAge) * 24 * time.Hour
	logInfoWriter, err := rotatelogs.New(
		apiLogInfoPath+".%Y-%m-%d-%H-%M.log",
		rotatelogs.WithLinkName(apiLogInfoPath),                              // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),                                        // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Duration(env.LogFileGap)*time.Hour), // 日志切割时间间隔
	)
	logErrorWriter, err := rotatelogs.New(
		apiLogErrorPath+".%Y-%m-%d-%H-%M.log",
		rotatelogs.WithLinkName(apiLogErrorPath),                             // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),                                        // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Duration(env.LogFileGap)*time.Hour), // 日志切割时间间隔
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logInfoWriter,
		logrus.ErrorLevel: logErrorWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{})
	logClient.AddHook(lfHook)
}

func RunLog(env *model.Env) {
	initLoad(env)
	for i := 0; i < 3; i++ {
		go func() {
			for log := range waitWrite {
				switch log.LogLevel {
				case LogLevelInfo:
					logClient.Infof(
						"| %s | %s | %+v | %s |",
						log.title,
						log.content,
						log.v,
						log.Track,
					)
				case LogLevelError:
					logClient.Errorf(
						"| %s | %s | %+v | %s |",
						log.title,
						log.content,
						log.v,
						log.Track,
					)
				}
			}
		}()
	}
}

type LogClient struct{}

func (l LogClient) Error(title, content string, v interface{}) {
	l.writeLog(LogLevelError, title, content, v)
}

func (l LogClient) Info(title, content string, v interface{}) {
	l.writeLog(LogLevelInfo, title, content, v)
}

func (l LogClient) writeLog(level int, title, content string, v interface{}) {
	_, file, line, _ := runtime.Caller(2)
	track := fmt.Sprintf("[%s:%d]", file, line)
	waitWrite <- &logData{
		LogLevel: level,
		title:    title,
		content:  content,
		v:        v,
		Track:    track,
	}
}
