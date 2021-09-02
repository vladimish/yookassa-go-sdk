package yookassa

type Error struct {
	Type        string `json:"type"`
	Id          string `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Parameter   string `json:"parameter"`
}