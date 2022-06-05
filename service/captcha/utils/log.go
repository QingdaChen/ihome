package utils

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

var Logger = NewLog()

type Log struct {
	log *logrus.Logger
}

func NewLog() *logrus.Logger {
	mLog := logrus.New()               //新建一个实例
	mLog.SetOutput(os.Stderr)          //设置输出类型
	mLog.SetReportCaller(true)         //开启返回函数名和行号
	mLog.SetFormatter(&LogFormatter{}) //设置自己定义的Formatter
	mLog.SetLevel(logrus.DebugLevel)   //设置最低的Level
	return mLog
}

type LogFormatter struct{}

//实现Formatter(entry *logrus.Entry) ([]byte, error)接口
func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//根据不同的level去展示颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	//自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		//自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		//自定义输出格式
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s\n", timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

////封装一些会用到的方法
//func (l *Log) Debug(args ...interface{}) {
//	l.log.Debugln(args...)
//}
//func (l *Log) Debugf(format string, args ...interface{}) {
//	l.log.Debugf(format, args...)
//}
//func (l *Log) Info(args ...interface{}) {
//	l.log.Infoln(args...)
//}
//func (l *Log) Infof(format string, args ...interface{}) {
//	l.log.Infof(format, args...)
//}
//func (l *Log) Error(args ...interface{}) {
//	l.log.Errorln(args...)
//}
//func (l *Log) Errorf(format string, args ...interface{}) {
//	l.log.Errorf(format, args...)
//}
//func (l *Log) Trace(args ...interface{}) {
//	l.log.Traceln()
//}
//func (l *Log) Tracef(format string, args ...interface{}) {
//	l.log.Tracef(format, args...)
//}
//func (l *Log) Panic(args ...interface{}) {
//	l.log.Panicln()
//}
//func (l *Log) Panicf(format string, args ...interface{}) {
//	l.log.Panicf(format, args...)
//}
//
//func (l *Log) Print(args ...interface{}) {
//	l.log.Println()
//}
//func (l *Log) Printf(format string, args ...interface{}) {
//	l.log.Printf(format, args...)
//}
