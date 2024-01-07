package dtos

type OrderDto struct {
	Id          string
	Pubkey      string
	CreatedAt   int64
	Status      int
	OrderStatus int
	PayStatus   int
	PaiedTime   int
	Price       int
	Lnwallet    string
	Comment     string
	Seller      string

	Skus []*OrderProductDto
}

type OrderProductDto struct {
	Id       string
	OrderId  string
	Pid      string
	DetailId string
	Seller   string
	Code     string
	Name     string
	Price    int
	Num      int
	Img      string
}
