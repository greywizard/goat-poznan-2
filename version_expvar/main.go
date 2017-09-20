package main

import (
	"expvar"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

const serviceName = "Example"

var (
	Version   = "unknown"
	BuildTime = ""
)

func main() {
	go func() {
		http.ListenAndServe("localhost:6565", nil)
	}()

	expVersion := expvar.String{}
	expBuildAt := expvar.String{}
	expGoVersion := expvar.String{}
	expRunAt := expvar.String{}

	expInfo := expvar.NewMap("build-info")
	expInfo.Set("Version", &expVersion)
	expInfo.Set("BuildAt", &expBuildAt)
	expInfo.Set("GoVersion", &expGoVersion)
	expInfo.Set("RunAt", &expRunAt)

	expVersion.Set(Version)
	expBuildAt.Set(BuildTime)
	expGoVersion.Set(runtime.Version())
	expRunAt.Set(time.Now().Format(time.RFC3339))

	fmt.Println("Pause")

	time.Sleep(10 * time.Second)
}
