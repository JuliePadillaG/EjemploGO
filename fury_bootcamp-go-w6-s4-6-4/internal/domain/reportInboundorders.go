package domain

type ReportInBO struct {
	ID                   int    `json:"id"`
	CardNumberID         string `json:"card_number_id"`
	FirstName            string `json:"first_name"`
	LastName             string `json:"last_name"`
	WarehouseID          int    `json:"warehouse_id"`
	Inbound_orders_count int    `json:"inbound_orders_count"`
}
