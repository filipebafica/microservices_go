package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/filipebafica/microservices_go/src/handlers"
)

func main () {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// creates an end point "object"
	hello_handler := handlers.NewHello(l)
	goodbye_handler := handlers.NewGoodbye(l)

	// matches the request to the correspont handler function
	mux := http.NewServeMux()
	mux.Handle("/", hello_handler)
	mux.Handle("/goodbye", goodbye_handler)

	// sets the server parameters
	server := &http.Server{
		Addr: ":9090",
		Handler: mux,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// light weighted thread that will execute in concurrency mode
	go func (){
		err := server.ListenAndServe()
		if err != nil {
			l.Println(err)
		}
	}()

	// Signal handling
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <- sigChan
	l.Println("Recieved terminate, graceful shutdown\n", sig)

	// waits until all current requests finishes before a server shoutdown happen
	ctx, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	server.Shutdown(ctx)
}
