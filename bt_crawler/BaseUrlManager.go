package bt_crawler

import "container/list"

type BaseUrlManager struct {
	QueryMap map[string]QueryInfo
	URLQueue list.List
}
type QueryInfo struct {

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
