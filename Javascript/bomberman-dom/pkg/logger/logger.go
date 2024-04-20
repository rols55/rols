package logger

import (
	"log"
	"os"
)

var (
	Info    = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime).Printf
	Warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile).Printf
	Error   = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile).Println
)
