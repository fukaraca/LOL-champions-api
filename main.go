package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"lol-champs-api/lib"
)

func init() {
	lib.ConnectDB()
	lib.InitServer()

}

func main() {
	routes()
	defer func() {
		if err := lib.Client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
	log.Fatalln("Router encountered and error while main.Run:", lib.R.Run(":8080"))

}

func routes() {
	lib.R.GET("/rest/champions", lib.GetRestRequest)

}

func karsilama(c *gin.Context) {
	url := c.Request.URL
	met := c.Param("method")
	fmt.Println(url.String())
	a := url.Query()
	counter := 0
	for s, v := range a {
		counter++
		fmt.Println("key:", s, "value:", v)

	}
	fmt.Println(counter)
	c.JSON(200, gin.H{
		"key":               a,
		"url:":              url.String(),
		"method":            met,
		"query parameters:": a,
	})
}
