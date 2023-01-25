package domain

type InputOperation struct {
	Type     string  `json:"operation"`
	UnitCost float64 `json:"unit-cost"`
	Quantity float64 `json:"quantity"`
}
