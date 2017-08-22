package routers

import (
	"net/http"

	"github.com/wj15github/gowot/controllers"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func SetTestRoutes(subRouter *mux.Router) {
	testHandler := negroni.Wrap(http.HandlerFunc(controllers.Test))
	subRouter.Handle("/test", logMiddleware.With(testHandler)).Methods(http.MethodGet)
}
