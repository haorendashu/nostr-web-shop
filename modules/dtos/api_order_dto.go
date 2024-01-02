package dtos

type ApiOrderDto struct {
	Seller         string `json:"seller"`
	Code           string `json:"code"`
	OrderProductID string `json:"orderProductId"`
	Buyer          string `json:"buyer"`
	Num            int    `json:"num"`
	Comment        string `json:"comment"`
	PaidTime       int    `json:"paidTime"`
}
