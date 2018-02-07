package main

import (
	"github.com/kevin-zx/go-util/mysqlUtil"
	"dianying_spider/douban_api"
	"fmt"
	"strconv"
	"time"
	//"github.com/henrylee2cn/pholcus/common/util"
)

type DoubanSearchSpider struct {
	mu mysqlutil.MysqlUtil
	tasks []task
}
type task struct {
	MovieId int
	MovieName string
	Result string
	SuccessFlag int
}


func (dbss *DoubanSearchSpider) init()  {
	dbss.initdb()
	go dbss.keepDbAlive()
	dbss.tasks = []task{}
}

func (dbss *DoubanSearchSpider) initdb()  {
	dbss.mu = mysqlutil.MysqlUtil{}
	err := dbss.mu.InitMySqlUtil("111.231.24.24",3306, "remote","Iknowthat@@!221","crawler_data")
	for  {
		if err != nil {
			time.Sleep(time.Second*10)
			err := dbss.mu.InitMySqlUtil("111.231.24.24",3306, "remote","Iknowthat@@!221","crawler_data")
			if err != nil {
				println("init",err.Error())
			}
		}else{
			break
		}
	}
}


func (dbss *DoubanSearchSpider) getTask(limit int) error {

	rows,err := dbss.mu.SelectAll(fmt.Sprintf("select * from movies where crawler_flag = 0 order by rand() limit %d", limit))
	if err != nil{
		return err
	}
	for _,row := range *rows{
		mId,_ := strconv.Atoi(row["id"])
		dbss.tasks = append(dbss.tasks,task{MovieId:mId, MovieName:row["movie_name"]} )
	}
	return nil
}

func (dbss *DoubanSearchSpider) start()  {
	if dbss.tasks == nil {
		dbss.init()
	}
	for {
		dbss.getTask(5)
		if len(dbss.tasks) == 0{
			time.Sleep(time.Hour * 1)
			continue
		}
		for _,task_it := range dbss.tasks {
			html,err := douban_api.GetQueryData(task_it.MovieName)
			if err != nil {
				println("start",err.Error())
				time.Sleep(time.Second * 20)
				task_it.SuccessFlag = -1

			}else{
				task_it.SuccessFlag = 1
				task_it.Result = html
			}
			dbss.store(task_it)
			time.Sleep(time.Millisecond * 500)
		}
		dbss.tasks = []task{}
		time.Sleep(time.Second * 5)
	}

}

func (dbss *DoubanSearchSpider) store(t task) {
	if t.SuccessFlag == 1{
		err := dbss.mu.Insert("INSERT INTO movie_search_data (`movie_id`,`data`) value (?,?)",t.MovieId,t.Result)
		for err != nil  {
			time.Sleep(time.Second * 10)
			//print(err.Error())
			err := dbss.mu.Insert("INSERT INTO movie_search_data (`movie_id`,`data`) value (?,?)",t.MovieId,t.Result)
			if err !=nil{
				println("store",err.Error())
			}
		}
	}
	dbss.completeTask(t)

}

func (dbss *DoubanSearchSpider) completeTask(t task)  {
	err := dbss.mu.Exec("UPDATE movies SET crawler_flag = ? WHERE id= ? ",t.SuccessFlag,t.MovieId)
	for err != nil  {
		time.Sleep(time.Second * 10)
		err := dbss.mu.Exec("UPDATE movies SET crawler_flag = ? WHERE id= ? ",t.SuccessFlag,t.MovieId)
		if err != nil{
			println("cpt",err.Error())
		}
	}
	println("完成电影", t.MovieName )
}
func (dbss *DoubanSearchSpider) keepDbAlive()  {
	for  {
		if !dbss.mu.IsAlive() {
			dbss.initdb()
		}
		time.Sleep(time.Second*5)
	}
}
func main(){
	//mysqlutil.GlobalMysqlUtil.InitMySqlUtil("115.159.3.51",3306, "remote","Iknowthat","eb_bigdata")

	//mu.InitMySqlUtil("111.231.24.24",3306, "remote","Iknowthat@@!221","crawler_data")
	//mu.SelectAll("")
	//html,err := douban_api.GetQueryData("")
	//5 35 55
	dss := DoubanSearchSpider{}
	dss.start()
}