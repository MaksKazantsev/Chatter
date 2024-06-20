package models

type RegRes struct {
	UUID         string `json:"uuid"`
	RefreshToken string `json:"refreshToken"`
	AccessToken  string `json:"accessToken"`
}

type LogRes struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
