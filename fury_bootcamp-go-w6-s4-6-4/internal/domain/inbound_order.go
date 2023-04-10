package domain

type Inbound_order struct {
	ID               int    `json:"id"`
	Order_date       string `json:"order_date"`
	Order_number     string `json:"order_number"`
	Employee_id      int    `json:"employee_id"`
	Product_batch_id int    `json:"product_batch_id"`
	Warehouse_id     int    `json:"warehouse_id"`
}
