package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

func Logrus() *logrus.Logger {
	log.Formatter = &logrus.JSONFormatter{}
	f, err := os.OpenFile("./applog.log", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("日志初始化error", err)
	}
	log.Out = f                  // 设置log的默认文件输出
	gin.SetMode(gin.ReleaseMode) // 线上模式，控制台不会打印信息
	gin.DefaultWriter = log.Out  // gin框架自己记录的日志也会输出
	log.Level = logrus.InfoLevel // 设置日志级别
	return log
}

var Logger FieldLogger = NewDefaultLogger()

type defaultLogger struct {
	*logrus.Entry
}

func NewDefaultLogger() *defaultLogger {
	return &defaultLogger{
		logrus.NewEntry(logrus.New()),
	}
}

func (l *defaultLogger) SetNoLock() {
	l.Logger.SetNoLock()
}

func (l *defaultLogger) SetLevel(level string) error {
	parseLevel, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}
	l.Logger.SetLevel(parseLevel)
	return nil
}

func (l *defaultLogger) WithField(key string, value interface{}) FieldLogger {
	l.Logger.SetNoLock()
	return &defaultLogger{l.Entry.WithField(key, value)}
}

func (l *defaultLogger) WithFields(fields map[string]interface{}) FieldLogger {
	return &defaultLogger{l.Entry.WithFields(fields)}
}

func (l *defaultLogger) WithError(err error) FieldLogger {
	return &defaultLogger{l.Entry.WithError(err)}
}

// Logger Define
type FieldLogger interface {
	WithField(key string, value interface{}) FieldLogger
	WithFields(fields map[string]interface{}) FieldLogger
	WithError(err error) FieldLogger

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	Debug(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})

	Debugln(args ...interface{})
	Infoln(args ...interface{})
	Println(args ...interface{})
	Warnln(args ...interface{})
	Warningln(args ...interface{})
	Errorln(args ...interface{})
	Fatalln(args ...interface{})
	Panicln(args ...interface{})

	Tracef(format string, args ...interface{})
	Trace(args ...interface{})
	Traceln(args ...interface{})

	// 设置日志级别
	SetLevel(level string) error
	// 设置读写锁
	SetNoLock()
}
