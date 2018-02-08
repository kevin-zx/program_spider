package bt_dog

import "program_spider/bt_crawler"

type BtDogUrlManager struct {
	bt_crawler.BaseUrlManager
}

func (bum *BtDogUrlManager) Init() {

}


func (bum *BtDogUrlManager) Next() (string,string) {

	return "",""
}
