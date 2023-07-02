package store

type Brand struct {
	Id   int    `json:"id"`
	Shop string `json:"shop"`
}
type Respon struct {
	Massage string  `json:"massage"`
	Data    []Brand `json:"data"`
}
