package main
import (
	"fmt"
	"encoding/json"
)
type AirQualityReading struct {
	SensorID  string    'json:"sensor_id"'
	Timestamp time.Time 'json:"timestamp"'
	PM25      float64   'json:"pm25"'
	CO2       float64 	'json:"co2"'
}
func parseReading(data []byte)([]AirQualityReading, error){
	jsonData, err := json.Marshal(data)
	if err!= nil {
		fmt.Println("could not marshal json %s\n", err)
	}
	json.Unmarshal([]byte(data), &result)
}
func calculateAverage(readings []AirQualityReading) map[string]float64{
	n:= len(data)
	sum1:=0
	sum2:=0
	for i := 0; i < n; i++ {
		sum += a6[i]
		sum2 += a8[i]
	}
	avg1 := (float64(sum1))/(float64(n))
	avg2 := (float64(sum2))/(float64(n))
	return avg1, avg2
}

func main(){
	data = [{"sensor_id": "S001", "timestamp": "2023-12-28T10:00:00Z", "pm25": 25.5, "co2": 410.2},
	{"sensor_id": "S002", "timestamp": "2023-12-28T10:05:00Z", "pm25": 30.8, "co2": 405.7},
	{"sensor_id": "S001", "timestamp": "2023-12-28T11:00:00Z", "pm25": 18.2, "co2": 395.1}]
	fmt.Println("avarage of pm25 =", avg1)
	fmt.Println("avarage of CO2 =", avg2))
}