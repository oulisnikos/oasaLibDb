package logger

import (
	"context"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	gormlogger "gorm.io/gorm/logger"
)

type GormLogger struct {
	LogLevel gormlogger.LogLevel
}

const (
	rootLogsDirpath = "/oasa-telematics"
)

var Logger *logrus.Logger

func InitLogger(applicationName string) {
	Logger = logrus.New()
	directoryPath := path.Join(rootLogsDirpath, applicationName)
	err := os.Mkdir(directoryPath, 0777)
	if err != nil {
		fmt.Printf("error create directory file: %v\n", err)
	}
	var runMode = os.Getenv("application.mode")
	if runMode == "PROD" {
		fileName := path.Join(rootLogsDirpath, applicationName, "opswlog.log")
		//open a file
		f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			fmt.Printf("error opening file: %v\n", err)
		}
		Logger.SetOutput(f)
	}

}

func INFO(str string) {
	Logger.Println(str)
}

func ERROR(str string) {
	Logger.SetLevel(logrus.ErrorLevel)
	Logger.Println(str)
}

// LogMode set log mode
func (l *GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

// Info prints info
func (l GormLogger) Info(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel >= gormlogger.Info {
		// Logger.Infof(str, args...)
		log.WithFields(log.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		}).Infof(str, args)
	}
}

// Warn prints warn messages
func (l GormLogger) Warn(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel >= gormlogger.Warn {
		// Logger.Warnf(str, args...)
		log.WithFields(log.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		}).Warnf(str, args)
	}

}

// Error prints error messages
func (l GormLogger) Error(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel >= gormlogger.Error {
		// Logger.Errorf(str, args...)
		log.WithFields(log.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		}).Infof(str, args)
	}
}

// Trace prints trace messages
func (l GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	// elapsed := time.Since(begin)
	// sql, rows := fc()
	// Logger.Info("[", elapsed.Milliseconds(), " ms, ", rows, " rows] ", "sql -> ", sql)
	// if l.LogLevel <= 0 {
	// 	return
	// }
	// elapsed := time.Since(begin)
	// if l.LogLevel >= gormlogger.Info {
	// 	sql, rows := fc()
	// 	Logger.Debug("[", elapsed.Milliseconds(), " ms, ", rows, " rows] ", "sql -> ", sql)
	// 	return
	// }

	// if l.LogLevel >= gormlogger.Warn {
	// 	sql, rows := fc()
	// 	Logger.Warn("[", elapsed.Milliseconds(), " ms, ", rows, " rows] ", "sql -> ", sql)
	// 	return
	// }

	// if l.LogLevel >= gormlogger.Error {
	// 	sql, rows := fc()
	// 	Logger.Error("[", elapsed.Milliseconds(), " ms, ", rows, " rows] ", "sql -> ", sql)
	// 	return
	// }

	if l.LogLevel <= gormlogger.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= gormlogger.Error: //&& (!errors.Is(err, ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			//l.Printf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			Logger.Error("[", elapsed.Milliseconds(), " ms, ", "sql -> ", sql)
		} else {
			//l.Printf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			Logger.Error("[", elapsed.Milliseconds(), " ms, ", rows, " rows] ", "sql -> ", sql)
		}
	case l.LogLevel >= gormlogger.Warn:
		sql, rows := fc()
		// slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			//l.Printf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			Logger.Warn("[", elapsed.Milliseconds(), " ms, ", "sql -> ", sql)
		} else {
			//l.Printf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			Logger.Info("[", elapsed.Milliseconds(), " ms, ", rows, " rows] ", "sql -> ", sql)
		}
	case l.LogLevel == gormlogger.Info:
		sql, rows := fc()
		if rows == -1 {

			Logger.Info("[", elapsed.Milliseconds(), " ms, ", "sql -> ", sql)
			//l.Info(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			//l.Printf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
			Logger.Info("[", elapsed.Milliseconds(), " ms, ", rows, " rows] ", "sql -> ", sql)
		}
	}
}

func GetGormLogger() *GormLogger {
	return &GormLogger{
		LogLevel: gormlogger.Info,
	}
}
