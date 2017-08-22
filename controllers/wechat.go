package controllers

import (
	"net/http"

	"github.com/leesper/holmes"
)

// @Title WxToken
// @Description response for wechat token verification
// @Accept  json
// @Success 200 {object}	models.CommonResponse
// @Resource /1.0
// @Router /1.0/wx_token [get]
func WxToken(w http.ResponseWriter, r *http.Request) {
	signature := r.URL.Query().Get("signature")
	timestamp := r.URL.Query().Get("timestamp")
	nonce := r.URL.Query().Get("nonce")
	echostr := r.URL.Query().Get("echostr")

	holmes.Infoln(signature, timestamp, nonce, echostr)
}
