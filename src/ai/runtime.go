package ai

import (
	"os"
	"runtime"
)

// provides runtime information like host os, shell, etc.

type Runtime struct {
	OS   	string
	Shell 	string
}

var _runtime *Runtime

func initRuntime() {
	_runtime = &Runtime{
		OS: 	getOS(),
		Shell: 	getShell(),
	}
}

func getOS() string {
	return runtime.GOOS
}

func getShell() string {
	return os.Getenv("SHELL")
}

func GetRuntime() *Runtime {
	if _runtime == nil {
		initRuntime()
	}
	return _runtime
}