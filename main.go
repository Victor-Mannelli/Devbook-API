package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

//* Creates JWT secret key
// func init() {
// 	key := make([]byte, 64)

// 	_, err := rand.Read(key)
// 	utils.CheckError(err)

// 	base64String := base64.StdEncoding.EncodeToString(key)
// 	fmt.Println(base64String)
// }
//* Creates JWT secret key

func main() {
	config.Load()
	r := router.NewRouter()

	// fmt.Println(config.JwtSecret)
	// fmt.Println(config.DBConnectionString)
	fmt.Printf("Listening to port %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
