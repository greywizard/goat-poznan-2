package main

import (
	"log"

	"github.com/fvbock/endless"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	Logger := logrus.New()

	server := endless.NewServer("127.0.0.1:8000", mux.NewRouter())
	server.ErrorLog = log.New(Logger.WriterLevel(logrus.ErrorLevel), "", 0)
	server.BeforeBegin = func(add string) {
		err := writePID()
		if err != nil {
			Logger.Error(err)
		}
	}

	err := server.ListenAndServe()
	if err != nil {
		Logger.Error(err)
	}
}

func writePID() error {
	//pid := syscall.Getpid()

	return nil
}
