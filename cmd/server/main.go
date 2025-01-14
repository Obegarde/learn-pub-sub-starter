package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"log"
	"github.com/Obegarde/learn-pub-sub-starter/internal/pubsub"
)

func main() {
	connectionString := "amqp://guest:guest@localhost:5672/"
	connection , err := amqp.Dial(connectionString)
	if err != nil{
		log.Fatal("Failed to create connection")

	}
	defer connection.Close()
	fmt.Println("Connection successful")
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	newChannel,err := connection.Channel()
	if err != nil{
		log.Fatal("Failed to create channel")	
	}

		
}
