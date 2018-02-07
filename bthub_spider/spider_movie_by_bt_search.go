package bthub_spider

import (
	"github.com/kevin-zx/go-util/httpUtil"
	"github.com/PuerkitoBio/goquery"
)

func main()  {
	resp,err :=httpUtil.GetWebResponseFromUrl("http://bthub.io/search?key=生活大爆炸")
	if err != nil{
		panic(err)
	}
	//print(htmldata)
	doc,err :=goquery.NewDocumentFromResponse(resp)
	if err!=nil{
		panic(err)
	}
	doc.Find(".bt-item").Each(func(i int, selection *goquery.Selection) {
		selection.Text()
	})
}