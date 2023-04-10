package domain

type CarriesReport struct {
	LocalityID   int    `json:"locality_id"`
	LocalityName string `json:"locality_name"`
	CarriesCount int    `json:"carries_count"`
}
