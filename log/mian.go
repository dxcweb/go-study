package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func initLog() {
	// //设置输出样式，自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	// logrus.SetFormatter(&logrus.JSONFormatter{})
	//设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
	logrus.SetOutput(os.Stdout)
	//设置最低loglevel
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	initLog()
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Warn("A walrus appears")
	level()
}
func level() {
	logrus.Debug("Useful debugging information.")
	logrus.Info("Something noteworthy happened!")
	logrus.Warn("You should probably take a look at this.")
	logrus.Error("Something failed but I'm not quitting.")
	logPanic()
	logrus.Fatal("Bye.") //log之后会调用os.Exit(1)
}
func logPanic() {
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	logrus.Panic("I'm bailing.") //log之后会panic()
}
