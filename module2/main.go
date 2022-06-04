package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func main(){
	http.HandleFunc("/",indexHandler)
	http.HandleFunc("/healthz",healthzHandler)
	http.ListenAndServe(":80",nil)
}
func healthzHandler(w http.ResponseWriter,req *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
func indexHandler(w http.ResponseWriter,req *http.Request){
	header := req.Header
	//1、读取请求头并写入响应头
	for k,v:=range header{
		w.Header().Set(k,v[0])
	}
	//2、获取当前版本并写入响应头
	v := runtime.Version()
	w.Header().Set("version",v)
	//3、记录日志
	w.WriteHeader(http.StatusOK)

	fmt.Println("statusCode:",http.StatusOK,",ip:",req.RemoteAddr)
}

