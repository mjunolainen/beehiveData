package structs

type SensorData struct {
	ID          string  `json:"id"`
	Hive        string  `json:"hive"`
	Temperature float64 `json:"temperature,omitempty"`
	Humidity    float64 `json:"humidity,omitempty"`
}
