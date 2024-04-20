package logger

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime"
)

var (
	Info    = log.New(os.Stdout, "INFO - App: ", log.Ldate|log.Ltime).Printf
	Warning = log.New(os.Stdout, "WARNING - App: ", log.Ldate|log.Ltime|log.Lshortfile).Printf
	Error   = log.New(os.Stdout, "ERROR - App: ", log.Ldate|log.Ltime|log.Lshortfile).Println
)

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func GetCurrentFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return fmt.Sprintf("%s", runtime.FuncForPC(pc).Name())
}
