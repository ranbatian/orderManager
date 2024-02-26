package common

type OrderType struct {
	Id          uint   `json:"id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	OrderKindId uint   `json:"orderKindId,omitempty"`
}
