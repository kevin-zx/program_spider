package bt_dog

import (
	"program_spider/bt_crawler"
	"strconv"
	"strings"
	"fmt"
)

type BtDogUrlManager struct {
	bt_crawler.BaseUrlManager
	startId int


}

func (bum *BtDogUrlManager) Init() {
	bum.PlatformPrefix = "bt_dog"
	d,err := bum.Mu.SelectAll("SELECT platform_unique FROM program_data WHERE platform_unique LIKE" +
		" '"+bum.PlatformPrefix +"%' ORDER BY id DESC LIMIT 1")
	if err !=nil {
		panic(err)
	}
	bum.SeedUrl = "http://www.btdog.com/bt/dog%d/"
	bum.startId = 0
	if len(*d) > 0 {
		platformUnique := (*d)[0]["platform_unique"]
		bum.startId,_ = strconv.Atoi(strings.Replace(platformUnique,"bt_dog","",1))
	}
}




func (bum *BtDogUrlManager) Next() (string,string) {
	bum.startId ++
	return fmt.Sprintf(bum.SeedUrl,bum.startId),fmt.Sprintf(bum.PlatformPrefix+"%d",bum.startId)
}
