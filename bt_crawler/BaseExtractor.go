package bt_crawler

import "github.com/PuerkitoBio/goquery"

type BaseExtractor struct {
	doc            goquery.Document
	platformUnique string
}

func (be *BaseExtractor) ExtractorProgram(doc *goquery.Document,platformUnique string) (Program) {
	btp := Program{
		Title:be.ExtractorTitle(),
		Alias:be.ExtractorAlias(),
		Status:be.ExtractorStatus(),
		Mark:be.ExtractorMark(),
		Directors:be.ExtractorDirectors(),
		Type:be.ExtractorType(),
		Languages:be.ExtractorLanguages(),
		Area:be.ExtractorArea(),
		ReleaseDate:be.ExtractorReleaseDate(),
		Thunders:be.ExtractorThunders(),
		Actors: be.ExtractorActors(),
		PlatformUnique: be.platformUnique,
		Category: be.ExtractorCategory(),
		Caption: be.ExtractorCaption(),
		}
	return btp
}



func (be *BaseExtractor) ExtractorTitle() (string){
	return ""
}
func (be *BaseExtractor) ExtractorCategory() (string){
	return ""
}
func (be *BaseExtractor) ExtractorCaption() (string){
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
func (be *BaseExtractor) ExtractorReleaseDate() ([]string){
	return nil
}

func (be *BaseExtractor) ExtractorDirectors() ([]Director){
	return []Director{}
}

func (be *BaseExtractor) ExtractorType() (string){
	return ""
}
func (be *BaseExtractor) ExtractorArea() (string){
	return ""
}

func (be *BaseExtractor) ExtractorLanguages() ([]string){
	return []string{}
}

func (be *BaseExtractor) ExtractorThunders() ([]ThunderData){
	return []ThunderData{}
}

func (be *BaseExtractor) ExtractorActors() ([]Actor){
	return []Actor{}
}
