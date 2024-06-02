package models

type RegRes struct {
	UUID  string `json:"uuid"`
	Token string `json:"-"`
}

type LogRes struct {
	Token string `json:"-"`
}
