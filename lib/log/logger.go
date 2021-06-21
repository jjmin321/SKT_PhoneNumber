package log

import (
	"log"
	"os"
)

// 정상 정보 로그 출력
var Info *log.Logger = log.New(os.Stdout, "INFO: ", log.LstdFlags)

// 오류 정보 로그 출력
var Err *log.Logger = log.New(os.Stdout, "ERROR: ", log.LstdFlags)
