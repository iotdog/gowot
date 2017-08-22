package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/leesper/holmes"
	"github.com/wj15github/gowot/models"
)

// @Title Test
// @Description test api
// @Accept  json
// @Success 200 {object}	models.CommonResponse
// @Resource /1.0
// @Router /1.0/test [get]
func Test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := &models.CommonResponse{
		Code: 0,
		Msg:  "have fun",
	}
	if err := json.NewEncoder(w).Encode(message); err != nil {
		holmes.Errorln(err)
	}
}
