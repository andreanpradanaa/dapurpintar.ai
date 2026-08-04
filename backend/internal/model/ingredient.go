package model

type PantryItem struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	NameID   string  `json:"nameId"`
	Category Category `json:"category"`
	Common   bool    `json:"common"`
}
