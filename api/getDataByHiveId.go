package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mjunolainen/GoAPIwithGinGonic/data"
	"github.com/mjunolainen/GoAPIwithGinGonic/structs"
	"net/http"
)

// GetHiveId filters the data based on Hive id and returns a slice
func GetHiveId(id string) ([]*structs.SensorData, error) {
	var hiveData []*structs.SensorData

	for i, v := range data.MockData {
		if v.Hive == id {
			hiveData = append(hiveData, &data.MockData[i])
		}
	}
	if len(hiveData) == 0 {
		errorMessage := fmt.Sprintf("Error: Hive with id: %v not found", id)
		return nil, errors.New(errorMessage)
	}
	return hiveData, nil
}

func GetDataByHiveId(c *gin.Context) {
	id := c.Param("id")
	hive, err := GetHiveId(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Hive not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, hive)
}
