package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"encoding/json"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root:Smith@2022@tcp(127.0.0.1:3306)/db_call_api?charset=utf8mb4&parseTime=True&loc=Local"
var db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 测试数据库设计一张表，字段有
type ApiCalling struct {
	ID   int64  `json:"id"`
	Type string `json:"type"`
	Url  string `json:"url"`
}

// db.AutoMigrate(&ApiCalling{})
//Create
// db.Create(&ApiCalling{Type: "dd", Url: "13sex.xyz"})

func doEveryOneHour(d time.Duration, f func(t time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func helloworld(t time.Time) {
	fmt.Printf("%v: Hello, World!\n", t)

	var url string
	fmt.Print("Enter the URL of a website: ")
	fmt.Scan(&url)

	// url := "https://api.new.urlzt.com/api/aa"

	// url = r.URL.Path[1:]

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	status_code := resp.StatusCode
	if status_code != 200 {
		// send to telegram
		fmt.Println(status_code, ":", "send to telegram group")
	} else {
		// nothing to do
		fmt.Println(status_code, ":", url)
	}
}

func getData(w http.ResponseWriter, r *http.Request) {
	apiCallings := []*ApiCalling{}
	db.Model(&ApiCalling{}).Where("type", "dd").Find(&apiCallings)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(apiCallings)
}

func main() {
	//1 hour = 360*time.Second
	doEveryOneHour(2*time.Second, helloworld)
	// http.HandleFunc("/getdata", getData)
	http.ListenAndServe(":8080", nil)
}
