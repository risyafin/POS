package branch

type Branch struct {
	Id   int    `json:"id"`
	Shop string `json:"shop"`
}
type Respon struct {
	Message string  `json:"massage"`
	Data    []Branch `json:"data"`
}
