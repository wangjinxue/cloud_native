package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"sync"
)

func startHttpServer(wg *sync.WaitGroup) *http.Server {
	srv := &http.Server{Addr: ":8080"}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/healthz", healthzHandler)

	go func() {
		defer wg.Done() // let main know we are done cleaning up

		// always returns error. ErrServerClosed on graceful close
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			// unexpected error. port in use?
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	// returning reference so caller can call Shutdown()
	return srv
}
func main() {
	httpServerExitDone := &sync.WaitGroup{}

	httpServerExitDone.Add(1)
	srv := startHttpServer(httpServerExitDone)

	log.Printf("main: serving for 10 seconds")

	//time.Sleep(10 * time.Second)

	log.Printf("main: stopping HTTP server")

	// now close the server gracefully ("shutdown")
	// timeout could be given with a proper context
	// (in real world you shouldn't use TODO()).
	if err := srv.Shutdown(context.TODO()); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}

	// wait for goroutine started in startHttpServer() to stop
	httpServerExitDone.Wait()

	log.Printf("main: done. exiting")
}
func healthzHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
func indexHandler(w http.ResponseWriter, req *http.Request) {
	header := req.Header
	//1、读取请求头并写入响应头
	for k, v := range header {
		w.Header().Set(k, v[0])
	}
	//2、获取当前版本并写入响应头
	v := runtime.Version()
	w.Header().Set("version", v)
	//3、记录日志
	w.WriteHeader(http.StatusOK)

	fmt.Println("statusCode:", http.StatusOK, ",ip:", req.RemoteAddr)
}
