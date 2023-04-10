package domain

type Carry struct {
	ID           int    `json:"id"`
	CID          string `json:"cid"`
	Company_name string `json:"company_name"`
	Address      string `json:"address"`
	Telephone    string `json:"telephone"`
	Locality_id  int    `json:"locality_id"`
}
