package log

import (
	"runtime"

	"github.com/sirupsen/logrus"
)

var l *logrus.Entry

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	l = logrus.WithFields(logrus.Fields{
		"module": "gungnir",
	})
}

func Debugf(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.WithFields(logrus.Fields{"file": file, "line": line}).Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.WithFields(logrus.Fields{"file": file, "line": line}).Infof(format, args...)
}

func Printf(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.WithFields(logrus.Fields{"file": file, "line": line}).Printf(format, args...)
}

func Warnf(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.WithFields(logrus.Fields{"file": file, "line": line}).Warnf(format, args...)
}

func Warningf(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.WithFields(logrus.Fields{"file": file, "line": line}).Warningf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.WithFields(logrus.Fields{"file": file, "line": line}).Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.WithFields(logrus.Fields{"file": file, "line": line}).Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.WithFields(logrus.Fields{"file": file, "line": line}).Panicf(format, args...)
}

func Debug(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.WithFields(logrus.Fields{"file": file, "line": line}).Debug(args...)
}

func Info(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.WithFields(logrus.Fields{"file": file, "line": line}).Info(args...)
}

func Print(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.WithFields(logrus.Fields{"file": file, "line": line}).Print(args...)
}

func Warn(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.WithFields(logrus.Fields{"file": file, "line": line}).Warn(args...)
}

func Warning(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.WithFields(logrus.Fields{"file": file, "line": line}).Warning(args...)
}

func Error(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.WithFields(logrus.Fields{"file": file, "line": line}).Error(args...)
}

func Fatal(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.WithFields(logrus.Fields{"file": file, "line": line}).Fatal(args...)
}

func Panic(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.WithFields(logrus.Fields{"file": file, "line": line}).Panic(args...)
}

func Debugln(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.WithFields(logrus.Fields{"file": file, "line": line}).Debugln(args...)
}

func Infoln(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.WithFields(logrus.Fields{"file": file, "line": line}).Infoln(args...)
}

func Println(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.WithFields(logrus.Fields{"file": file, "line": line}).Println(args...)
}

func Warnln(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.WithFields(logrus.Fields{"file": file, "line": line}).Warnln(args...)
}

func Warningln(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.WithFields(logrus.Fields{"file": file, "line": line}).Warningln(args...)
}

func Errorln(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.WithFields(logrus.Fields{"file": file, "line": line}).Errorln(args...)
}

func Fatalln(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.WithFields(logrus.Fields{"file": file, "line": line}).Fatalln(args...)
}

func Panicln(args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	l.WithFields(logrus.Fields{"file": file, "line": line}).Panicln(args...)
}
