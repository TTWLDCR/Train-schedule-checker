package main

import (
	"Train-schedule-checker/db"
	"Train-schedule-checker/model"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	//连接users数据库
	UserDB, err := db.Mysql("127.0.0.1", 3306, "root", "123456", "users")
	if err != nil {
		panic(err)
	}
	defer UserDB.Close()

	//连接train数据库
	TrainDB, err := db.Mysql("127.0.0.1", 3306, "root", "123456", "train")
	if err != nil {
		panic(err)
	}
	defer TrainDB.Close()

	UserDB.AutoMigrate(&model.User{})

	TrainDB.AutoMigrate(&model.Train{})

	TrainDB.AutoMigrate(&model.TrainStation{})

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", gin.H{})
	})
	r.POST("/login", func(c *gin.Context) {
		var user model.User //接收前端传过来的登录信息
		var user1 model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 验证用户名和密码
		if err := UserDB.Where(&user).First(&user1).Error; err == nil && user1.Password == user.Password {
			c.JSON(http.StatusOK, gin.H{"message": "登录成功"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "用户名或密码错误"})
		}
	})

	r.GET("/register", func(c *gin.Context) {
		c.HTML(200, "register.html", gin.H{})
	})
	r.POST("/register", func(c *gin.Context) {
		var user model.RegisterUser //接收前端传过来的注册信息
		var user1 model.User
		var user2 model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user1.Username = user.Username
		user1.Password = user.Password
		//验证用户名和密码
		if err := UserDB.Where("username = ?", user1.Username).First(&user2).Error; err == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "用户名已存在"})
		} else if user.Password != user.ConfirmPassword {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "密码不一致"})
		} else {
			err := UserDB.Create(&user1).Error
			if err != nil {
				log.Fatal(err)
			}
			c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
		}
		//if err := UserDB.Where("username = ?", user1.Username).First(&user2).Error; err != nil {
		//	if errors.Is(err, gorm.ErrRecordNotFound) {
		//		// 用户不存在，可以执行注册操作
		//		// 将user1插入到数据库中
		//		if user.Password != user.ConfirmPassword {
		//			c.JSON(http.StatusUnauthorized, gin.H{"message": "密码不一致"})
		//		} else {
		//			err := UserDB.Create(&user1).Error
		//			if err != nil {
		//				// 处理插入操作中的错误
		//				log.Fatal(err)
		//			} else {
		//				// 注册成功
		//				c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
		//			}
		//		}
		//	} else {
		//		log.Fatal(err)
		//	}
		//} else {
		//	// 用户已存在，处理重复注册情况
		//	c.JSON(http.StatusUnauthorized, gin.H{"message": "用户名已存在"})
		//}
	})

	r.GET("/homepage", func(c *gin.Context) {
		c.HTML(200, "homepage.html", gin.H{})
	})

	r.GET("/station-to-station", func(c *gin.Context) {
		c.HTML(200, "station-to-station.html", gin.H{})
	})

	r.POST("/station-to-station", func(c *gin.Context) {
		var stationToStation model.StationToStation
		if err := c.ShouldBindJSON(&stationToStation); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var trains []model.Train
		TrainDB.Table("trains").
			Select("train_name, start_station, end_station, start_time, end_time, duration, hard_seat_price, hard_sleeper_price, soft_sleeper_price").
			Where("start_station = ? AND end_station = ?", stationToStation.StartStation, stationToStation.EndStation).
			Find(&trains)
		//trainsJson, _ := json.Marshal(trains)
		fmt.Println(trains)
		c.JSON(http.StatusOK, trains)
	})

	r.GET("/train-search", func(c *gin.Context) {
		c.HTML(200, "train-search.html", gin.H{})
	})

	r.POST("/train-search", func(c *gin.Context) {
		var trainSearch model.TrainSearch
		if err := c.ShouldBindJSON(&trainSearch); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var trainStations []model.TrainStation
		TrainDB.Table("train_stations").
			Select("train_name, station_name, arrival_time, departure_time, distance").
			Where("train_name = ?", trainSearch.TrainName).
			Order("distance ASC").
			Find(&trainStations)
		c.JSON(http.StatusOK, trainStations)
	})

	r.GET("/station-search", func(c *gin.Context) {
		c.HTML(200, "station-search.html", gin.H{})
	})

	r.POST("/station-search", func(c *gin.Context) {
		var stationSearch model.StationSearch
		if err := c.ShouldBindJSON(&stationSearch); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var trainStations []model.TrainStation
		TrainDB.Table("train_stations").
			Select("train_name, station_name, arrival_time, departure_time, distance").
			Where("station_name = ?", stationSearch.StationName).
			Find(&trainStations)
		c.JSON(http.StatusOK, trainStations)
	})

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
