package model

type TrainStation struct {
	TrainName     string `json:"trainName"`     //车次名称
	StationName   string `json:"stationName"`   //车站名称
	ArrivalTime   string `json:"arrivalTime"`   //到达时间
	DepartureTime string `json:"departureTime"` //出发时间
	Distance      int    `json:"distance"`      //里程
}
