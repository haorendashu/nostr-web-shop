package dtos

type ApiOrderDto struct {
	Seller         string `json:"seller"`
	Code           string `json:"code"`
	OrderProductID string `json:"orderProductId"`
	Buyer          string `json:"buyer"`
	Num            int    `json:"num"`
	Comment        string `json:"comment"`
	PayStatus      int    `json:"payStatus"`
	PaidTime       int    `json:"paidTime"`
}
