package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mjunolainen/GoAPIwithGinGonic/api"
)

func main() {
	router := gin.Default()
	// curl localhost:3000/sensor_data
	router.GET("/sensor_data", api.GetSensorData)
	// curl localhost:3000/new_sensor_entry --include --header "Content-Type: application/json" -d @body.json --request "POST"
	router.POST("/new_sensor_entry", api.NewSensorEntry)
	// curl localhost:3000/hives/1 for example
	router.GET("hives/:id", api.GetDataByHiveId)
	router.Run("localhost:3000")
	api.GetHiveId("2")
}
