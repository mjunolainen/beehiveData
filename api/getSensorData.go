package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mjunolainen/GoAPIwithGinGonic/data"
	"net/http"
)

func GetSensorData(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.MockData)
}
