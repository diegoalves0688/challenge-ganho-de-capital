package domain

type InputLine struct {
	Operation string  `json:"operation"`
	UnitCost  float64 `json:"unit-cost"`
	Quantity  int64   `json:"quantity"`
}
