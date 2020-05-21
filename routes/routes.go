package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/saiprasadkrishnamurthy/reference-data-api/config"
	"github.com/saiprasadkrishnamurthy/reference-data-api/controllers"
)

const apiBase = "/api"

// InitialiseAllRoutes initialises all routes in the API.
func InitialiseAllRoutes(r *httprouter.Router) {
	apiContext := apiBase + "/" + config.APIVersion()
	appController := controllers.BaseController{}

	// List all your controllers here.
	tagsController(apiContext, r, appController)
}

func tagsController(apiContext string, r *httprouter.Router, baseController controllers.BaseController) {
	c := &controllers.TagsController{BaseController: baseController}
	r.GET(apiContext+"/tags", c.Action(c.Tags))
}
