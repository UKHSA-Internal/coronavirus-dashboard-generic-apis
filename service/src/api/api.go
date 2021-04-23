package api

import (
	"net/http"

	"generic_apis/apps/areaByType"
	"generic_apis/apps/code"
	"generic_apis/apps/soa"
	"generic_apis/middleware"
	"github.com/gorilla/mux"
	unit "unit.nginx.org/go"
)

type (
	Api struct {
		Router *mux.Router
	}

	routeEntry struct {
		name    string
		path    string
		handler func(http.ResponseWriter, *http.Request)
	}
)

var getRoutes = []routeEntry{
	{
		"code",
		`/generic/code/type/{area_type:[a-zA-Z]{4,10}}/code/{area_code:[a-zA-Z0-9+%\s]{3,12}}`,
		code.QueryByCode,
	},
	{
		"soa",
		`/generic/soa/type/{area_type:[a-zA-Z]{4,10}}/code/{area_code:[a-zA-Z0-9]{3,10}}`,
		soa.SoaQuery,
	},
	{
		"area",
		`/generic/area/type/{area_type:[a-zA-Z]{4,10}}`,
		areaByType.AreaByTypeQuery,
	},
} // routes

func (api *Api) Initialize() {

	api.Router = mux.NewRouter()
	api.initializeRoutes()

} // Initialize

func (api *Api) Run(addr string) {

	api.Initialize()

	// Uncomment for testing
	// if err := http.ListenAndServe(addr, api.Router); err != nil {
	// 	panic(err)
	// }

	// Comment for testing
	// This will only run inside the container.
	if err := unit.ListenAndServe(addr, api.Router); err != nil {
		panic(err)
	}

} // Run

func (api *Api) initializeRoutes() {

	api.Router.Use(middleware.HeadersMiddleware)

	for _, route := range getRoutes {
		api.Router.
			HandleFunc(route.path, route.handler).
			Name(route.name).
			Methods("GET")
	}

} // initializeRoutes
