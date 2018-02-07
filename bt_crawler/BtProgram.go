package bt_crawler

type BtProgram struct {
	Title string
	Alias []string
	Status string
	Mark string
	Director string
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