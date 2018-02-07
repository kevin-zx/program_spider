package bt_crawler

type Program struct {
	Title string
	Alias []string
	Status string
	Mark string
	Director Director
	Type string
	Language string
	Area string
	AddTime string
	Score float64
	ScoreCount int
	Caption string
	BaiduYuns []BaiduYunData
	Thunders []ThunderData
	Actors string
	PlatformUnique string
	Category string
}


type Actor struct {
	Name string

}

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