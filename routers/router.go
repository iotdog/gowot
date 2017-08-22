package routers

import (
	"net/http"

	"github.com/wj15github/gowot/controllers"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

var (
	logMiddleware = negroni.New(
		negroni.HandlerFunc(controllers.LoggerMiddleware),
	)
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/html/").Handler(http.StripPrefix("/html/", http.FileServer(http.Dir("./html")))) // for static pages
	apiRouter := router.PathPrefix("/api/1.0").Subrouter()                                               // for apis of version 1.0
	SetTestRoutes(apiRouter)
	SetWechatRoutes(apiRouter)
	return router
}
