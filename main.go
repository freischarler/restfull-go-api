package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/martinpaz/restfulapi/internal/data"
	"github.com/martinpaz/restfulapi/internal/server"
)

func main() {
	//port := os.Getenv("PORT")
	serv, err := server.New("8000")
	if err != nil {
		log.Fatal(err)
	}

	// connection to the database.
	d := data.New()
	// Check the connections
	err = d.DB.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// start the server.
	go serv.Start()

	// Wait for an in interrupt.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Attempt a graceful shutdown.
	serv.Close()
	data.Close()
}
