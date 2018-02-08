package bt_crawler

import (
	"container/list"
	"github.com/kevin-zx/go-util/mysqlUtil"
)

type BaseUrlManager struct {
	QueryMap map[string]QueryInfo
	URLQueue list.List
	Mu mysqlutil.MysqlUtil
	PlatformPrefix string
	SeedUrl string
}
type QueryInfo struct {
	Status int
	TryTimes int

}

func (bum *BaseUrlManager) Init() {

}

func (bum *BaseUrlManager) InsertUrl(url string)  {

}

func (bum *BaseUrlManager) HasNext() bool {
	return false
}

func (bum *BaseUrlManager) Next() (string,string) {
	return "",""
}

func (bum *BaseUrlManager) Complete(url string)  {

}


func (bum *BaseUrlManager) Failed(url string)  {

}
