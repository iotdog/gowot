// gowot - GOlang server for Wechat Official accounTs

package main

import (
	"net/http"

	"github.com/wj15github/gowot/routers"

	"github.com/leesper/holmes"
	"github.com/urfave/negroni"
)

func main() {
	defer holmes.Start().Stop()

	router := routers.InitRoutes()
	n := negroni.New()
	n.UseHandler(router)

	holmes.Errorln(http.ListenAndServe(":3000", n))
}
