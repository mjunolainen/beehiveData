package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mjunolainen/GoAPIwithGinGonic/data"
	"github.com/mjunolainen/GoAPIwithGinGonic/structs"
	"net/http"
)

func NewSensorEntry(c *gin.Context) {

	var newSensorEntry structs.SensorData

	if err := c.BindJSON(&newSensorEntry); err != nil {
		return
	}
	data.MockData = append(data.MockData, newSensorEntry)
	c.IndentedJSON(http.StatusCreated, newSensorEntry)
}
