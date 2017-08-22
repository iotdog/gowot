package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/wj15github/gowot/controllers"
)

func SetWechatRoutes(subRouter *mux.Router) {
	wxTokenHandler := negroni.Wrap(http.HandlerFunc(controllers.WxToken))
	subRouter.Handle("/wx_token", logMiddleware.With(wxTokenHandler)).Methods(http.MethodGet)
}
