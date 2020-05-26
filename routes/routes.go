package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/saiprasadkrishnamurthy/reference-data-api/config"
	"github.com/saiprasadkrishnamurthy/reference-data-api/controllers"
)

const apiBase = "/api"

// InitialiseAllRoutes initialises all routes in the API.
func InitialiseAllRoutes(r *httprouter.Router) {
	// CORS.
	r.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", r.Header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
		}

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})
	apiContext := apiBase + "/" + config.APIVersion()
	appController := controllers.BaseController{}

	// List all your controllers here.
	tagsController(apiContext, r, appController)
}

func tagsController(apiContext string, r *httprouter.Router, baseController controllers.BaseController) {
	c := &controllers.TagsController{BaseController: baseController}
	r.GET(apiContext+"/tags", c.Action(c.Tags))
}
