package main

import (
	"github.com/PuerkitoBio/goquery"
	"dianying_spider/bt_crawler"
	"fmt"
	"strings"
	"github.com/kevin-zx/go-util/mysqlUtil"
	"strconv"
	"time"
)

func main()  {
	mu := mysqlutil.MysqlUtil{}
	mu.InitMySqlUtil("111.231.24.24",3306,"remote","Iknowthat@@!221","crawler_data")
	id := getStartId(&mu)
	for i:=id+1;i<1000000;i++  {
		doc,err := goquery.NewDocument(fmt.Sprintf("https://www.77kp.com/vod-detail-id-%d.html",i));
		if err != nil {
			fmt.Println(err)
			i--
			time.Sleep(10*time.Second)
			continue
		}
		btProgram := bt_crawler.BtProgram{}

		detailBoxElement := doc.Find("#detail-box")
		detailListElement := doc.Find("#detail-list")

		btProgram.PlatformUnique = fmt.Sprintf("77kp%d",i)
		btProgram.Title,btProgram.Alias = extractTitle(detailBoxElement);
		fmt.Println(btProgram.Title)
		if btProgram.Title == "" {
			continue
		}
		btProgram.Type = strings.TrimSpace(doc.Find("li.current").Text())
		btProgram.Category = doc.Find("div.position a:nth-child(2)").Text()
		btProgram.Actors = extractActors(detailBoxElement)
		btProgram.Mark = extractMark(detailBoxElement)
		btProgram.Thunders = extractDownloadUrl(detailListElement)
		btProgram.Caption = doc.Find("#detail-intro p").Text()
		insert(&mu,btProgram)
	}

}

func insert(mu *mysqlutil.MysqlUtil,btp bt_crawler.BtProgram)  {
	id,err := mu.InsertId("INSERT IGNORE INTO program_data " +
		"(`title`,`mark`,`actors`,`caption`,`platform_unique`,`type`,`category`) " +
		"VALUE (?,?,?,?,?,?,?)",
			btp.Title,btp.Mark,btp.Actors,btp.Caption,btp.PlatformUnique,btp.Type,btp.Category)
	if err != nil {
		panic(err)
	}
	if id ==0{
		return
	}
	if btp.Alias != nil{
		for _,al := range btp.Alias  {
			err = mu.Insert("INSERT INTO program_alias (`program_id`,`alias`) VALUE (?,?)",id,strings.TrimSpace(al))
			if err != nil {
				panic(err)
			}
		}
	}

	if btp.Thunders != nil {
		for _,td:= range btp.Thunders {
			err = mu.Insert("INSERT INTO program_thunder_data (`program_id`,`url`,`name`,`type`) VALUE (?,?,?,?)",id,td.URL,td.Name,td.Type)
			if err != nil {
				panic(err)
			}
		}
	}
}

func getStartId(mu *mysqlutil.MysqlUtil) int {
	data,err := mu.SelectAll("SELECT platform_unique FROM program_data WHERE  platform_unique LIKE '77kp%' ORDER BY id DESC LIMIT 1")
	if err !=nil {
		panic(err)
	}
	if len((*data))>0 {

		platform_unique := (*data)[0]["platform_unique"]
		idStr := strings.Replace(platform_unique,"77kp","",1)
		id,_ := strconv.Atoi(idStr)
		return id
	}else{
		return 0
	}
}

func extractTitle(detailBox *goquery.Selection) (string,[]string) {
	titleStr := detailBox.Find(".detail-title").Text();
	titleStr = strings.Replace(titleStr,"在线观看","",1)
	titleParts := strings.Split(titleStr,"/")
	title := strings.Trim(strings.TrimSpace(titleParts[0]),"	")
	alias := make([]string,len(titleParts)-1)
	if len(titleParts) > 1 {
		alias = titleParts[1:]
	}
	return title,alias
}

func extractActors(detailBox *goquery.Selection) (string) {
	return detailBox.Find(".detail-info dl:nth-child(1) dd").Text()
}

func extractMark(detailBox *goquery.Selection) (string) {
	return detailBox.Find(".detail-info dl:nth-child(2) dd").Text()
}

func extractDownloadUrl(detailList *goquery.Selection) []bt_crawler.ThunderData {
	detailListElements :=detailList.Find("span>a[title]");
	tds := []bt_crawler.ThunderData{}
	//fmt.Println(detailListElements.Length())
	detailListElements.Each(func(_ int, a *goquery.Selection) {
		url,ok:=a.Attr("href")
		if(ok) {
			name, _ := a.Attr("title")
			tType := ""
			if strings.Contains(url, "thunder://") {
				tType = "thunder"
			} else if strings.Contains(url, "magnet:") {
				tType = "magnet"
			}

			td := bt_crawler.ThunderData{URL: url, Name: name, Type: tType}
			tds = append(tds, td)
		}
	})
	
	return tds
}

func extractorType()  {

}