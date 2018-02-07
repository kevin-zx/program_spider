package main

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
	"fmt"
	"strconv"

	"dianying_spider/bt_crawler"
)



func main() {
	doc,err := goquery.NewDocument("http://www.btdog.com/bt/dog40823/");
	if err != nil {
		panic(err)
	}

	btp := bt_crawler.BtProgram{}
	btp.Title = doc.Find("#detail-box > div.box.box-blue > div.detail-title.fn-left > h2").Text()
	aliasStr := doc.Find("#detail-box > div.box.box-blue > div.detail-title.fn-left > div").Text()
	btp.Alias = strings.Split(aliasStr,"/")
	btp.Status = doc.Find("#detail-box > div.box.box-blue > div.detail-cols.fn-clear > div.detail-info.fn-left > div.info.fn-clear > dl:nth-child(1) > dd > span").Text()
	//btp.Director = doc.Find("#detail-box > div.box.box-blue > div.detail-cols.fn-clear > div.detail-info.fn-left > div.info.fn-clear > dl:nth-child(3) > dd > span > a").Text()
	btp.Type = doc.Find("#detail-box > div.box.box-blue > div.detail-cols.fn-clear > div.detail-info.fn-left > div.info.fn-clear > dl:nth-child(4) > dd > a").Text()
	btp.Language = doc.Find("#detail-box > div.box.box-blue > div.detail-cols.fn-clear > div.detail-info.fn-left > div.info.fn-clear > dl:nth-child(5) > dd > span").Text()
	btp.Area = doc.Find("#detail-box > div.box.box-blue > div.detail-cols.fn-clear > div.detail-info.fn-left > div.info.fn-clear > dl:nth-child(6) > dd > span").Text()
	btp.AddTime = doc.Find("#addtime").Text()
	btp.Score,_ =	strconv.ParseFloat(doc.Find("#detail-box > div.box.box-blue > div.detail-cols.fn-clear > div.detail-info.fn-left > div.info.fn-clear > dl:nth-child(9) > dd > span.Goldnum").Text(),64)
	btp.ScoreCount,_ = strconv.Atoi(doc.Find("#detail-box > div.box.box-blue > div.detail-cols.fn-clear > div.detail-info.fn-left > div.info.fn-clear > dl:nth-child(9) > dd > span.Golder").Text());

	dirtyCaption :=doc.Find("#detail-intro > div > div:nth-child(2)").Text()
	btp.Caption = strings.Replace(dirtyCaption,doc.Find("#detail-intro > div > div:nth-child(2) > p").Text(),"",1);

	ThunderElements := doc.Find("#down-pl-list input")
	btp.Thunders = make([]bt_crawler.ThunderData,ThunderElements.Length())
	for i:=0;i< ThunderElements.Length() ;i++  {
		Te := doc.Find(fmt.Sprintf("#down-pl-list input:nth-child(%d)",i+1))
		fmt.Println(Te.Attr("value"))
	}

	//playListBox := doc.Find("#detail-list .play-list-box")
	//playListBox.Find("#btdown-pl-list")



	fmt.Print(btp)
}
