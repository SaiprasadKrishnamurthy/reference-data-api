package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/saiprasadkrishnamurthy/reference-data-api/routes"
	"github.com/urfave/negroni"

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

	server := &http.Server{
		Addr:         port,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	n := negroni.Classic()
	n.UseHandler(r)

	go func() {
		n.Run(port)
	}()

	interruptChannel := make(chan os.Signal)
	signal.Notify(interruptChannel, os.Interrupt)
	signal.Notify(interruptChannel, os.Kill)

	sig := <-interruptChannel
	fmt.Println("Received Signal: ", sig)

	tc, _ := context.WithTimeout(context.Background(), 5*time.Second)
	server.Shutdown(tc)
}
