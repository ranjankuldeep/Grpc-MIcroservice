package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"working/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	sm := http.NewServeMux()
	sm.Handle("/", hh)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	// graceful shutdown
	signChan := make(chan os.Signal)
	signal.Notify(signChan, os.Interrupt)
	signal.Notify(signChan, os.Kill)

	// block until signal is received
	sig := <-signChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	defer cancel()
	s.Shutdown(tc)
}
