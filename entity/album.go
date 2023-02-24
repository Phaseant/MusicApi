package entity

type Album struct {
	Id          int    `json:"albumId" bson:"_id"`
	Title       string `json:"title" bson:"title"`
	Author      string `json:"author" bson:"author"`
	Year        int    `json:"year" bson:"year"`
	CoverURL    string `json:"coverURL" bson:"coverURL"`
	Description string `json:"description" bson:"description"`
	Duration    string `json:"duration" bson:"duration"`
	Songs       []struct {
		Id       int    `json:"songId" bson:"_id"`
		Title    string `json:"title" bson:"title"`
		Duration string `json:"duration" bson:"duration"`
	} `json:"songs" bson:"songs"`
	GeniousLink string `json:"geniousLink" bson:"geniousLink"`
}
