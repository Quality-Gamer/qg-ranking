package endpoint

import (
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"github.com/Quality-Gamer/qg-ranking/model"
)

func GetRank(c echo.Context) error {
	var res model.Response
	c.Response().Header().Set("Access-Control-Allow-Origin","*")
	c.Response().Header().Set(echo.HeaderContentType,echo.MIMEApplicationJSONCharsetUTF8)

	if len(c.FormValue("page")) > 0 {
		rank := model.GetRank()
		page,_ := strconv.Atoi(c.FormValue("page"))
		leaders := rank.GetLeaders(page)
		res.Status = "OK"
		res.Message = "success"
		res.Response.Total = rank.TotalMembers()
		res.Response.Rank = leaders
		c.Response().WriteHeader(http.StatusOK)
		return json.NewEncoder(c.Response()).Encode(res)
	} else {
		res.Status = "NOK"
		res.Message = "Paramêtros inválidos"
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(res)
	}

}
