package domain

// "yyyy-mm-dd T00:00:00"
type Product_batches struct {
	ID                 int    `json:"id"`
	BatchNumber        int    `json:"section_number"`
	CurrentQuantity    int    `json:"current_quantity"`
	CurrentTemperature int    `json:"current_temperature"`
	DueDate            string `json:"due_date"`
	InitialQuantity    int    `json:"initial_quantity"`
	ManufacturingDate  string `json:"manufacturing_date"`
	ManufacturingHour  int    `json:"manufacturing_hour"`
	MinimumTemperature int    `json:"minimum_temperature"`
	ProductId          int    `json:"product_id"`
	SectionId          int    `json:"section_id"`
}

type ReportProduct struct {
	SectionId       int `json:"section_id"`
	SectionNumber   int `json:"section_number"`
	CurrentQuantity int `json:"current_quantity"`
}
