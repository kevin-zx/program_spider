package bt_crawler

import (
	"github.com/kevin-zx/go-util/mysqlUtil"
	"github.com/PuerkitoBio/goquery"
	"fmt"
)

type BaseCrawler struct {
	mu mysqlutil.MysqlUtil
	bum BaseUrlManager
	be BaseExtractor
}

func (bc *BaseCrawler) Run()  {
	bc.be = BaseExtractor{}
	bc.bum.Init()
	for bc.bum.HasNext()  {
		url,platformUnique := bc.bum.Next()
		doc,err := bc.crawlerUrl(url)
		if err != nil {
			fmt.Println(err)
			bc.bum.Failed(url)
			continue
		}
		bc.bum.Complete(url)
		program := bc.be.ExtractorProgram(doc,platformUnique)
		if program.Title != ""{
			continue
		}else {
			
		}

	}
}

func (bc *BaseCrawler) crawlerUrl(url string) (*goquery.Document,error) {
	return goquery.NewDocument(url)
}

func (bc *BaseCrawler) InsertProgram(program *Program)  {
	bc.mu.InsertId("INSERT IGNORE INTO program_data " +
		"(`title`,`mark`,`actors`,`caption`,`platform_unique`,`type`,`category`) " +
		"VALUE (?,?,?,?,?,?,?)",program.Title,program.Mark,program.Actors)
}