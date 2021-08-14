package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	host                       = "localhost"
	port                       = "8080"
	readtimeout  time.Duration = time.Duration(5) * time.Second
	writetimeout time.Duration = time.Duration(10) * time.Second
	idletimeout  time.Duration = time.Duration(20) * time.Second
)

func newHTTPServer(mux http.Handler) *http.Server {
	return &http.Server{
		Addr:         host + ":" + port,
		ReadTimeout:  readtimeout,
		WriteTimeout: writetimeout,
		IdleTimeout:  idletimeout,
		Handler:      mux,
	}
}

func main() {
	api := &API{}
	srv := newHTTPServer(api)
	ts := make(chan os.Signal, 1)
	signal.Notify(ts, os.Interrupt, syscall.SIGABRT, syscall.SIGTERM)
	log.Printf("Server starting at %s", srv.Addr)
	go func() {
		log.Fatalf("Server fail to start with error: %v \n", srv.ListenAndServe())
	}()
	log.Println(<-ts)
	signal.Stop(ts)
}
