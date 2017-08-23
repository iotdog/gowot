package controllers

import (
	"net/http"
	"sort"

	"github.com/leesper/holmes"
	"github.com/wj15github/gowot/models"
)

// @Title WxToken
// @Description response for wechat token verification
// @Success 200 {object}	models.CommonResponse
// @Resource /1.0
// @Router /1.0/wx_token [get]
func WxToken(w http.ResponseWriter, r *http.Request) {
	signature := r.URL.Query().Get("signature")
	timestamp := r.URL.Query().Get("timestamp")
	nonce := r.URL.Query().Get("nonce")
	echostr := r.URL.Query().Get("echostr")

	if signature == "" || timestamp == "" || nonce == "" || echostr == "" {
		holmes.Warnln("not a wechat request")
		w.WriteHeader(http.StatusBadRequest)
		// w.Write([]byte("not a wechat request"))
		return
	}

	holmes.Infoln(signature, timestamp, nonce, echostr)

	token := "havefun"
	list := []string{token, timestamp, nonce}

	sort.Sort(models.Alphabetic(list))

	holmes.Infoln(list)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(echostr))
}
