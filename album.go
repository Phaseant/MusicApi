package MusicAPI

type Album struct {
	Id          int    `json:"albumId"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Year        int    `json:"year"`
	CoverURL    string `json:"coverURL"`
	Description string `json:"description"`
	Duration    string `json:"duration"`
	Songs       []struct {
		Id       int    `json:"songId"`
		Title    string `json:"title"`
		Duration string `json:"duration"`
	} `json:"songs"`
	GeniousLink string `json:"geniousLink"`
}
