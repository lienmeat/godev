package controller

import (
	"github.com/gin-gonic/gin"
)

/**
* This file exists to make it simple to convert swagger API codegen Routes output for a go server (http://editor.swagger.io/#/)
* into routes usable by gin.
* 
* Copy/paste the generated var routes = {...} routers.go into here.
* Change any of the handlers to whatever you want to use for each endpoint.
*
* You will need to convert the generated handler funcs be of type gin.HandlerFunc if you want to use the same names as codegen made though.
*
* After doing this though, calling setRoutes() in main.go to pull in your routes should "just work"
*/

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

type Routes []Route

func SetRoutes(e *gin.Engine) {
	//assign all routes to router
	for _, route := range routes {

		//there's probably a way to do exactly this in one line...but...
		//Eric is new to go...dynamic method calls didn't want to happen
		switch route.Method {
			case "GET":
				e.GET(route.Pattern, route.HandlerFunc)
			case "POST":
				e.POST(route.Pattern, route.HandlerFunc)
			case "PUT":
				e.PUT(route.Pattern, route.HandlerFunc)
			case "PATCH":
				e.PATCH(route.Pattern, route.HandlerFunc)
			case "DELETE":
				e.DELETE(route.Pattern, route.HandlerFunc)
			case "HEAD":
				e.HEAD(route.Pattern, route.HandlerFunc)
			case "OPTIONS":
				e.OPTIONS(route.Pattern, route.HandlerFunc)
		}
	}
}

//###############  Router defn below here #################

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/ping",
		ping,
	}
}