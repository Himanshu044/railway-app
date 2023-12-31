package controllers

import (
	"backend/lib"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetTrains(context *gin.Context) {
	client := lib.GetRailwayClient()
	queryParams := context.Request.URL.Query()
	fromStationCode := queryParams.Get("fromStationCode")
	toStationCode := queryParams.Get("toStationCode")
	dateOfJourney := queryParams.Get("dateOfJourney")
	if len(dateOfJourney) == 0 {
		dateOfJourney = time.Now().UTC().Format("2006-01-02")
	}
	response, err := client.GetTrains(fromStationCode, toStationCode, dateOfJourney)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err})
	}
	context.JSON(http.StatusOK, response)
}
