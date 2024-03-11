package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Manas8803/learn-microservice/product-api/handlers"
)

var port_addr string = ":9020"

func main() {
	product_logger := log.New(os.Stdout, "products-api", log.LstdFlags)

	ph := handlers.NewProducts(product_logger)

	sm := http.NewServeMux()
	sm.Handle("/", ph)

	server := http.Server{
		Addr:         port_addr,
		Handler:      ph,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	//? These lines of code are explained in the sigChan.txt file
	//* Graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	wait_for_sig := <-sigChan
	product_logger.Println("RECEIVED TERMINATE, Graceful shutdown", wait_for_sig)

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	server.Shutdown(ctx)
}
