package main

import (
	"log"
	"net/http"

	"apascualco.com/api-rest/src/server"
	sqlitle3db "apascualco.com/api-rest/src/sqlite3"
)

func main() {
	sqlitle3db.LaunchDDL()

	server := server.BuildServerRoutes()
	log.Fatal(http.ListenAndServe(":8080", server.Handler()))
}
