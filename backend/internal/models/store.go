package models

type Store struct {
	ID          int    `json:"id"`
	DisplayName string `json:"displayName"`
	SystemName  string `json:"systemName"`
	ChainID     int    `json:"chainName"`
}
