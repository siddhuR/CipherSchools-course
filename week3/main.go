package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/siddhuR/CipherSchools-course/tree/main/Week2/database"
	"github.com/siddhuR/CipherSchools-course/tree/main/Week2/handler"

	//"github.com/siddhuR/CipherSchools-course/tree/main/Week2/routers"
	"github.com/siddhuR/CipherSchools-course/tree/main/Week2/routers"
)

func init() {
	database.Setup()
}
func init() {
	fmt.Println(1)
}
func init() {
	fmt.Println(2)
}
func AuthMiddleware(c) gin.Handlerfunc {
	//fmt.Println("Im a dummy!")
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		fmt.Println("got token: " + token)
		isValid, err := handler.ValidateToken(token)
		if err != nil && !isValid {
			respondwithError(c, 401, "Invalid API token")
			return
		}
		c.Next()
	}
}

// func AuthMiddleware(c *gin.Context) {
// 	token := c.GetHeader("token")
// 	fmt.Println("get token: " + token)
// 	c.Next()
// }

func main() {
	engine := gin.Default() //get the default engine for future customization
	api := handler.Handler{
		DB: database.GetDB(), //set the handler DB
	}
	engine.Use(AuthMiddleware())
	routers.BookRouter(engine, api)
	routers.AuthRouter(engine, api)

	// engine := routers.Router()          //get the customized engine
	//engine.Use(AuthMiddleware)
	err := engine.Run("127.0.0.1:8080") //start the engine
	if err != nil {
		log.Fatal(err)

	}
}

/*
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type something struct {
	Name    string  `json:"name"`
	Married bool    `json:"married"`
	Age     float64 `json:"age"`
	//Address []Address `json:"address"`
	Marks []int `json:"marks`
}

func main() {
	jsonFile, err := os.Open("something.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	jsonByteValues, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	var something something
	json.Unmarshal(jsonByteValues, &something) //convert json data to struct
	log.Println(something)
	//fmt.Print(string(jsonByteValues))

	// newJsonByteValues, err:=json.Marshal(something)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//os.WriteFile("some.json",newJsonByteValues)

	//consuming rest api in go code
	//using the rest api
	//returned vslue 7 = json
	//now you must read the json data
	//what is the problem if your read json data in string format
	//read json in structure format - marshalling
}


*/
