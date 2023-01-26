package domain

type OperationInput struct {
	Type     OperationType `json:"operation"`
	UnitCost float64       `json:"unit-cost"`
	Quantity float64       `json:"quantity"`
}

type OperationType string

const (
	BuyOperationType  OperationType = "buy"
	SellOperationType OperationType = "sell"
)
