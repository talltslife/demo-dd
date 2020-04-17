package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/talltslife/demo-dd/services"
	"net/http"
)

func SearchBars(c *gin.Context) {
	params := c.Request.URL.Query()
	limit := params.Get("limit")
	location := params.Get("location")
	longitude := params.Get("longitude")
	latitude := params.Get("latitude")

	var longLatBol = true
	if longitude == "" {
		longLatBol = false
	}
	if latitude == "" {
		longLatBol = false
	}
	if (location == "") && (longLatBol == false) {
		err := errors.New("you did not provide location or coordinates")
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := services.SearchBars(longitude, latitude, location, limit)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, res)
}
