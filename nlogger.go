package nlogger

import (
	"github.com/Sirupsen/logrus"
	"github.com/mattn/go-colorable"
)

func init() {
	logrus.SetOutput(colorable.NewColorableStdout())
}

//	新規ロガー
type Loggers struct {
	Info  logger
	Warn  logger
	Error logger
	Fatal logger
}

type logger func(args ...interface{})

//	非表示化
func (l *logger) Idle() {
	*l = func(args ...interface{}) {}
}

//	全て非表示化
func (l *Loggers) AllIdle() {
	l.Info.Idle()
	l.Warn.Idle()
	l.Error.Idle()
	l.Fatal.Idle()
}

//	全て非表示化
func AllIdle() {
	Info.Idle()
	Warn.Idle()
	Error.Idle()
	Fatal.Idle()
}

//	個別ロガー
func NewLogger() *Loggers {
	return &Loggers{
		Info:  logrus.Infoln,
		Warn:  logrus.Warnln,
		Error: logrus.Errorln,
		Fatal: logrus.Fatalln,
	}
}

//	パッケージ固有ロガー
var (
	Info  logger = logrus.Infoln
	Warn  logger = logrus.Warnln
	Error logger = logrus.Errorln
	Fatal logger = logrus.Fatalln
)

func Example() {
	//	すべて非表示
	//	AllIdle()
	//	Info表示を行わない
	Info.Idle()
	Info("succeeded")
	Warn("not correct")
	Error("something error")
	Fatal("panic")
}
