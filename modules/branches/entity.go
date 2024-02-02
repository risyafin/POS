package branches

import "gorm.io/gorm"

type Branch struct {
	gorm.Model
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type Respon struct {
	Message string   `json:"massage"`
	Data    any `json:"data"`
}
