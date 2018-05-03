// 向boss_movie添加电影
// json格式
//  {
//	"title" : "tile", #must
//	"type": "电影" , #must
//	"category": ["动作","爱情"], #must
//	"aliases": ["",""], #
//	"poster": "http://xxx.jpg", # 海报图
//	"images": ["",""], # 电影的一些图片
//	"pub_date": "", # 发行日期 1999-01-02 形式
//	"year": "", # 发行年份
//	"summary": "", #节目简介
//	#"seasons_count": 10, # 总季数 只针对电视剧
//	"current_seasons": 8, # 当前的季数 只针对电视剧
//	"episodes_count" : 24, # 当前集数
//	"mark":"",
//	"resources": [
//	{"description":"","title":"","link":"","type":""},
//	{"description":"","title":"","link":"","type":""}
//	] # 节目资源 type 类型 如 thunder magnet 百度云
//	}
//

package main

import (
	"github.com/kevin-zx/go-util/mysqlUtil"
	"program_spider/movie_api"
	"strconv"
	"fmt"
	"encoding/json"
	"github.com/kevin-zx/go-util/httpUtil"
	"time"
	"net/http"
	"strings"
	"io"
)

var mu mysqlutil.MysqlUtil
func main() {
	err := mu.InitMySqlUtil("111.231.24.24",3306, "remote","Iknowthat@@!221","crawler_data")
	if err != nil {
		panic(err)
	}
	currentId := 289
	for {
		data,err := mu.SelectAll("SELECT * FROM program_data where id > ? limit 1000", currentId)
		if err != nil {
			panic(err)
		}
		for _,d :=range *data{

			pid,_ := strconv.Atoi(d["id"])
			currentId = pid
			title := d["title"]
			pType := d["type"]
			category := []string{d["category"]}
			mark := d["mark"]
			summary := d["caption"]
			poster := d["poster"]
			actorStr := d["actors"]
			var members []movie_api.Member
			for _,actName := range strings.Split(actorStr," "){
				if strings.TrimSpace(actName) !="" {
					member := movie_api.Member{Name:actName,Role:"主演"}
					members = append(members, member)
				}
			}

			resources := getResources(pid)
			aliases := getAliases(pid)

			p := movie_api.Program{Title:title,
						Type: pType,
						Category:category,
						Mark:mark,
						Summary:summary,
						Poster:poster,
						Resources:resources,
						Aliases:aliases,
						Members:members,}


			pJson,err := json.Marshal(&p)
			if err != nil {
				panic(err)
			}

			rb := sendRequest("http://127.0.0.1:8000/sp_movie/api/programs/",string(pJson))
			fmt.Println(p.Title,rb,currentId)

		}
	}

}

func sendRequest(url string,data string) string  {
	var r http.Request
	r.ParseForm()
	r.Form.Add("data",data)
	bodyStr := strings.TrimSpace(r.Form.Encode())
	response,err := httpUtil.DoRequest(url,nil,"POST",[]byte(bodyStr),10*time.Second)
	if err != nil {
		panic(err)
	}
	//fmt.Println(getContentFromResponse(response))
	rb,err := getContentFromResponse(response)
	if err != nil {
		panic(err)
	}
	return rb
}

func getContentFromResponse(response *http.Response) (string, error) {
	defer response.Body.Close()
	var c []byte
	for {
		buf := make([]byte, 1024)
		n, err := response.Body.Read(buf)
		if n == 0 {
			break
		}
		if err != nil && err != io.EOF {
			return "", err
		}

		c = append(c, buf[0:n]...)
	}
	return string(c), nil
}
func getAliases(pid int) []string {
	aliases := []string{}
	data,err := mu.SelectAll("SELECT * FROM program_aliases WHERE program_id = ?",pid)
	if err != nil {
		panic(err)
	}
	for _,d := range *data  {
		aliases = append(aliases,d["alias"])
	}
	return aliases

}

func getResources(pid int) []movie_api.Resource {
	 resources := []movie_api.Resource{}
	data,err := mu.SelectAll("SELECT * FROM program_thunder_data WHERE program_id = ?",pid)
	if err != nil {
		panic(err)
	}
	for _,d := range *data {
		name := d["name"]
		rType := d["type"]
		url := d["url"]

		resource := movie_api.Resource{Title:name,Type: rType,Link:url}
		resources = append(resources,resource)
	}
	return resources
}

