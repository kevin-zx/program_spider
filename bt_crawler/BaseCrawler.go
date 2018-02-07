package bt_crawler

import "github.com/kevin-zx/go-util/mysqlUtil"

type BaseCrawler struct {
	mu mysqlutil.MysqlUtil
}

func (bc *BaseCrawler) Run()  {
	
}