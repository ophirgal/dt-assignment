package models

type Store struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	ChainID int    `json:"chain_id"`
}
