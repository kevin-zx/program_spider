package douban_api

import (
	"github.com/kevin-zx/go-util/httpUtil"
	"fmt"
	"net/url"
)
var URL_FORMAT = "https://api.douban.com/v2/movie/search?q=%s"
func GetQueryData(q string) (string,error) {
	q = url.QueryEscape(q)
	htmlData,err := httpUtil.GetWebConFromUrl(fmt.Sprintf(URL_FORMAT,q))
	return htmlData,err
}