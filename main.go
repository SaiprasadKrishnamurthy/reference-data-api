package main

import (
	"log"
	"net/http"

	"github.com/saiprasadkrishnamurthy/reference-data-api/routes"

	"github.com/julienschmidt/httprouter"
	"github.com/saiprasadkrishnamurthy/reference-data-api/config"
)

var r *httprouter.Router

func init() {
	r = httprouter.New()
	config.InitConfigs()
	routes.InitialiseAllRoutes(r)

}

// @title Reference Data API written in Golang
// @version 1.0
// @description Reference Data API  Golang to demonstrate it's simplicity.
// @termsOfService http://swagger.io/terms/

// @contact.name Sai Kris
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8082
// @BasePath /api/v1
// @query.collection.format multi

func main() {
	port := config.GetPort()
	log.Println(" Starting the Server on port: ", port)
	fs := http.FileServer(http.Dir("./swaggerui"))
	http.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", fs))
	http.Handle("/", r)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalln(err)
	}
}
