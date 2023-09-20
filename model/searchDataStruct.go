package model

type StationToStation struct {
	StartStation string `json:"startStation"` //出发站
	EndStation   string `json:"endStation"`   //到达站
}

type TrainSearch struct {
	TrainName string `json:"trainName"`
}

type StationSearch struct {
	StationName string `json:"stationName"`
}
