package main

import (
	"log"
	"net/http"
	"runtime"
)

func main(){
	http.HandleFunc("/",indexHandler)
	log.Fatal(http.ListenAndServe(":9999",nil))
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
	statusCode :=req.Response.StatusCode
	println(statusCode)
}

