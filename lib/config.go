package lib

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"log"
	"os"
)

func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://rootuser:rootpass@localhost:27017"))
	if err != nil {
		log.Fatalln("connection failed:", err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}
	Client = client
	Coll = Client.Database("lol-infos").Collection("champions")

}

func InitServer() {
	logfile, err := os.OpenFile("./logs/log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Could not create/open log file", err)
	}
	errlogfile, err := os.OpenFile("./logs/errlog.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Could not create/open err log file", err)
	}
	gin.DefaultWriter = io.MultiWriter(logfile, os.Stdout)
	gin.DefaultErrorWriter = io.MultiWriter(errlogfile, os.Stdout)
	R = gin.Default()
}
