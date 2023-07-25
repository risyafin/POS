package branches

type Branch struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type Respon struct {
	Message string   `json:"massage"`
	Data    []Branch `json:"data"`
}
