package bt_crawler

import "strings"

type Program struct {
	Title string
	Alias []string
	Status string
	Mark string
	Directors []Director
	Type string
	Languages []string
	Area string
	ReleaseDate []string //上映日期
	DoubanScore DoubanScore
	IMBScore IMBScore
	Caption string
	BaiduYuns []BaiduYunData
	Thunders []ThunderData
	Actors []Actor
	PlatformUnique string
	Category string

}

func (p *Program) DirectorsToStr() string {
	var directors []string
	for _,drt :=range p.Directors{
		directors = append(directors,drt.Name)
	}
	return strings.Join(directors,"/")
}

func (p *Program) LanguagesToStr() string  {

	return strings.Join(p.Languages,"/")
}

func (p *Program) AliasToStr() string {
	return strings.Join(p.Alias,"/")
}

func (p *Program) ActorsToStr() string  {
	var actors []string
	for _,act := range p.Actors{
		actors = append(actors, act.Name)
	}
	return strings.Join(actors,"/")
}

type Actor struct {
	Name string
}

type DoubanScore struct {
	Score float64
	Count int
	Status int
	URL string
}

type IMBScore struct {
	Score  float64
	Count  int
	Status int
	URL string
}


//type Languages struct {
//	Language string
//}

type Director struct {
	Name string
}

type ThunderData struct{
	URL string
	Name string
	Source string
	Type string //magnet or thunder

}
type BaiduYunData struct{
	URL string
	Passwd string
	Status string
}