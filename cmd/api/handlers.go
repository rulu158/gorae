package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rulu158/gorae"
)

func (srv *Server) GetWordDefinition(ctx *gin.Context) {
	word := ctx.Param("word")

	minify := true
	if ctx.Param("minify") != "" &&
		(ctx.Param("minify") == "0" || ctx.Param("minify") == "false") {
		minify = false
	}

	json, err := gorae.QueryWordJSON(word, minify)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "query failed; word not in DRAE? : "+err.Error()+" : "+word)
		return
	}

	ctx.String(http.StatusOK, json)
}
