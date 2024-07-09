package models

type RegRes struct {
	UUID         string `json:"uuid"`
	RefreshToken string `json:"refreshToken"`
	AccessToken  string `json:"accessToken"`
}
