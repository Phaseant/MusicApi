package entity

type User struct {
	Id       int    `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Username string `json:"username" bson:"username"`
	Password string `json:"-" bson:"password"`
}
