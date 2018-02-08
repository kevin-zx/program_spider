package bt_crawler

import (
	"github.com/kevin-zx/go-util/mysqlUtil"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"strings"
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
			bc.InsertProgram(&program)
		}

	}
}

func (bc *BaseCrawler) crawlerUrl(url string) (*goquery.Document,error) {
	return goquery.NewDocument(url)
}

func (bc *BaseCrawler) InsertProgram(program *Program)  {
	id,err := bc.mu.InsertId("INSERT IGNORE INTO program_data " +
		"(`title`,`mark`,`actors`,`caption`,`platform_unique`,`type`" +
			",`category`,`status`,`director`,`languages`) " +
		"VALUE (?,?,?,?,?,?,?,?,?,?)",program.Title,program.Mark,program.ActorsToStr(),
			program.Caption,program.PlatformUnique,program.Type,program.Caption,program.Status,
				program.DirectorsToStr(),program.Languages)

	if err !=nil {
		panic(err)
	}

	if id == 0 {
		return
	}

	if program.Alias != nil{
		for _,al := range program.Alias  {
			err = bc.mu.Insert("INSERT INTO program_alias (`program_id`,`alias`) " +
				"VALUE (?,?)",id,strings.TrimSpace(al))
			if err != nil {
				panic(err)
			}
		}
	}

	if program.Thunders != nil {
		for _,td:=range program.Thunders {
			err = bc.mu.Insert("INSERT INTO program_thunder_data " +
				"(`program_id`,`url`,`name`,`type`) " +
					"VALUE (?,?,?,?)",id,td.URL,td.Name,td.Type)
			if err != nil {
				panic(err)
			}
		}
	}



}