package entities

type TycketLot struct {
	Number         int     `json:"number,omitempty"`
	TycketOptionId string  `json:"tycket_option_id,omitempty"`
	TycketPrice    float64 `json:"tycket_price,omitempty"`
	TycketAmount   int     `json:"tycket_amount,omitempty"`
}

func NewTycketLot(number int, tycketOptionId string, tycketPrice float64, tycketAmount int) *TycketLot {
	return &TycketLot{
		Number:         number,
		TycketOptionId: tycketOptionId,
		TycketPrice:    tycketPrice, // TODO: validate positive numbers
		TycketAmount:   tycketAmount,
	}
}
