package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		log.Println("eba")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("chu"))
	})

	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)

	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		sig := <-gracefulStop
		log.Printf("caught sig: %+v\n", sig)
		log.Println("Wait for 2 second to finish processing")
		time.Sleep(2 * time.Second)
		os.Exit(0)
	}()

	log.Println("started")
	log.Fatal(srv.ListenAndServe())
}
