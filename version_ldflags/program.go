package main

import (
	"fmt"
	"runtime"
)

const serviceName = "Example"

var (
	Version   = "unknown"
	BuildTime = ""
)

func main() {
	InfoMessage := "      ..:: %s ::..\nBuilt at: %s \nVersion: %s (with: %s)\n\n"
	fmt.Printf(InfoMessage, serviceName, BuildTime, Version, runtime.Version())
}


