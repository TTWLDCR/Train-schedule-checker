package model

type Train struct {
	TrainName        string `json:"trainName" gorm:"column:train_name"`                 //车次名称
	StartStation     string `json:"startStation,omitempty" gorm:"column:start_station"` //出发站
	EndStation       string `json:"endStation,omitempty" gorm:"column:end_station"`     //到达站
	StartTime        string `json:"startTime" gorm:"column:start_time"`                 //出发时间
	EndTime          string `json:"endTime" gorm:"column:end_time"`                     //到达时间
	Duration         string `json:"duration" gorm:"column:duration"`                    //历时
	HardSeatPrice    int    `json:"hardSeatPrice" gorm:"column:hard_seat_price"`        //硬座
	HardSleeperPrice int    `json:"hardSleeperPrice" gorm:"column:hard_sleeper_price"`  //硬卧
	SoftSleeperPrice int    `json:"softSleeperPrice" gorm:"column:soft_sleeper_price"`  //软卧
}
