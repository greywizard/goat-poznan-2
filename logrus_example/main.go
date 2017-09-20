package main

import (
	"log/syslog"
	"os"

	"github.com/sirupsen/logrus"
	logrus_syslog "github.com/sirupsen/logrus/hooks/syslog"
)

func main() {
	Logger := logrus.New()
	Logger.Formatter = &logrus.JSONFormatter{}
	Logger.Out = os.Stderr
	Logger.Level = logrus.ErrorLevel

	sysLog, err := logrus_syslog.NewSyslogHook("udp", "127.0.0.1:500",
		syslog.LOG_ERR, "go-example")

	if err != nil {
		panic("Can't attach syslog hook: " + err.Error())
	}
	Logger.Hooks.Add(sysLog)
}


