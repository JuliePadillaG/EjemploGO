package domain

type ProductRecordsReport struct {
	ProductID   int    `json:"product_id"`
	ProductDescription string `json:"description"`
	ProductRecordsCount int    `json:"records_count"`
}