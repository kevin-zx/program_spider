package bt_crawler

import "github.com/PuerkitoBio/goquery"

type BaseExtractor struct {
	doc goquery.Document
}

func (be *BaseExtractor) ExtractorProgram() (Program) {
	btp := Program{
		Title:be.ExtractorTitle(),
		Alias:be.ExtractorAlias(),
		Status:be.ExtractorStatus(),
		Mark:be.ExtractorMark(),
		//Director:be.Ex
		}
	return btp
}

func (be *BaseExtractor) ExtractorTitle() (string){
	return ""
}

func (be *BaseExtractor) ExtractorAlias() ([]string){
	return nil
}

func (be *BaseExtractor) ExtractorStatus() (string){
	return ""
}

func (be *BaseExtractor) ExtractorMark() (string){
	return ""
}

func (be *BaseExtractor) ExtractorTitle() (string){
	return ""
}

func (be *BaseExtractor) ExtractorTitle() (string){
	return ""
}

func (be *BaseExtractor) ExtractorTitle() (string){
	return ""
}