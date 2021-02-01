package model

import "github.com/cheekybits/genny/generic"

type Response struct {
	Status   string `json:"status"`
	Response Res    `json:"response"`
	Message  string `json:"message"`
}

type Res struct {
	Rank  []generic.Type `json:"rank"`
	Total int            `json:"total"`
}
