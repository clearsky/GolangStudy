package main

import (
	"GolangStudy/7/errorhandling/filelistgingserver/filelisting"
	"github.com/gpmgo/gopm/modules/log"
	"net/http"
	"os"
)

type appHandle func(write http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandle) func(http.ResponseWriter, *http.Request){  //类似装饰器
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func(){
			r := recover()
			log.Warn("Panic: %v", r)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}()
		err := handler(writer, request)
		if err != nil{
			log.Warn("Error handling request: %s", err.Error())
			if userError, ok := err.(userError);ok{
				http.Error(writer, userError.Message(), http.StatusBadRequest)
				return
			}
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer,http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/list/", errWrapper(filelisting.HandleFileList))
	err :=http.ListenAndServe(":8888", nil)
	if err != nil{
		panic(err)
	}
}
