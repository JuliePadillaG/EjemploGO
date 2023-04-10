package domain

type ProductRecords struct {
	ID             	  int     	`json:"id"`
	LastUpdateDate    string  	`json:"last_update_date"`
	PurchasePrice 	  float64   `json:"purchase_price"`
	SalePrice   	  float64   `json:"sale_price"`
	ProductID         int 		`json:"products_id"`
}