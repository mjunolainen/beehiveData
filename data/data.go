package data

import (
	"github.com/mjunolainen/GoAPIwithGinGonic/structs"
)

var Hive1 = "1"
var Hive2 = "2"

var Hives = []string{
	Hive1, Hive2,
}

var MockData = []structs.SensorData{
	{ID: "1", Hive: Hives[0], Temperature: 25.434, Humidity: 100.00},
	{ID: "2", Hive: Hives[1], Temperature: 18.3456, Humidity: 65.03},
	{ID: "3", Hive: Hives[0], Temperature: 26.234, Humidity: 99.20},
	{ID: "4", Hive: Hives[1], Temperature: 19.6978, Humidity: 68.83},
}
