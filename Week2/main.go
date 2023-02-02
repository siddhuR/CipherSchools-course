package main

import (
	"log"

	"github.com/siddhuR/CipherSchools-course/tree/main/Week2/database"
	"github.com/siddhuR/CipherSchools-course/tree/main/Week2/routers"
)

func main() {
	database.Setup()                    //establish the database connection
	engine := routers.Router()          //get the customized engine
	err := engine.Run("127.0.0.1:8080") //start the engine
	if err != nil {
		log.Fatal(err)

	}
}
