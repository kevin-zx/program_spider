package main

import (
	"program_spider/bt_crawler/bt_dog"
	"github.com/kevin-zx/go-util/mysqlUtil"
	"program_spider/bt_crawler"
)

func main() {
	mu := mysqlutil.MysqlUtil{}
	mu.InitMySqlUtil("111.231.24.24",3306,"remote","Iknowthat@@!221","crawler_data")
	btc := bt_dog.BtDogCrawler{bt_crawler.BaseCrawler{}}


}
