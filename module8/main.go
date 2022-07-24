package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func main() {
	//ctx, cancel := context.WithCancel(context.Background())
	//
	//mux := http.NewServeMux()
	httpServer := &http.Server{Addr: ":8080"}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/healthz", healthzHandler)

	log.Printf("main: serving for 10 seconds")

	//time.Sleep(10 * time.Second)

	log.Printf("main: stopping HTTP server")

	// Run server
	go func() {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			// it is fine to use Fatal here because it is not main gorutine
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()

	signalChan := make(chan os.Signal, 1)

	signal.Notify(
		signalChan,
		syscall.SIGHUP,  // kill -SIGHUP XXXX
		syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
		syscall.SIGQUIT, // kill -SIGQUIT XXXX
	)

	<-signalChan
	log.Print("os.Interrupt - shutting down...\n")

	go func() {
		<-signalChan
		log.Fatal("os.Kill - terminating...\n")
	}()

	gracefullCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := httpServer.Shutdown(gracefullCtx); err != nil {
		log.Printf("shutdown error: %v\n", err)
		defer os.Exit(1)
		return
	} else {
		log.Printf("gracefully stopped\n")
	}

	// manually cancel context if not using httpServer.RegisterOnShutdown(cancel)
	cancel()

	defer os.Exit(0)
	return
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
