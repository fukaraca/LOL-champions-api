package main

import (
	"context"
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
	//routes for version REST API
	lib.R.GET("/rest/champions/list", lib.GetHeroList)
	lib.R.GET("/rest/champions", lib.GetRestRequest)
	lib.R.POST("/rest/champions", lib.PostRestRequestAsForm)
	lib.R.POST("/rest/champions/createWJSON", lib.PostRestRequestAsJSON)
	lib.R.PATCH("/rest/champions", lib.UpdateRestHero)
	lib.R.DELETE("/rest/champions", lib.DeleteRestHero)

}
