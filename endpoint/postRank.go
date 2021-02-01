package endpoint

import (
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"qg-ranking/model"
)

func PostRank(c echo.Context) error {
	var res model.Response
	c.Response().Header().Set("Access-Control-Allow-Origin","*")
	c.Response().Header().Set(echo.HeaderContentType,echo.MIMEApplicationJSONCharsetUTF8)

	if len(c.FormValue("user_id")) > 0 && len(c.FormValue("score")) > 0 {
		uid := c.FormValue("user_id")
		score, _ := strconv.Atoi(c.FormValue("score"))
		rank := model.GetRank()
		rank.RankMember(uid, score)
		res.Status = "OK"
		res.Message = "success"
		c.Response().WriteHeader(http.StatusOK)
		return json.NewEncoder(c.Response()).Encode(res)
	} else {
		res.Status = "NOK"
		res.Message = "Paramêtros inválidos"
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(res)
	}

}
