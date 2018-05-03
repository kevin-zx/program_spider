//boss movie 的添加接口数据格式
package movie_api

type Program struct {
	Title    string `json:"title"`
	Type     string `json:"type"`
	Category []string `json:"category"`
	Aliases  []string `json:"aliases"`
	Poster string `json:"poster"`
	Images []string `json:"images"`
	PubDate string `json:"pub_date"`
	Year string `json:"year"`
	Summary string `json:"summary"`
	CurrentSeasons int `json:"current_seasons"`
	EpisodesCount int `json:"episodes_count"`
	Mark string `json:"mark"`
	Resources []Resource `json:"resources"`
}

type Resource struct {
	Description string `json:"description"`
	Title string `json:"title"`
	Link string `json:"link"`
	Type string `json:"type"`
}



