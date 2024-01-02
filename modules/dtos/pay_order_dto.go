package dtos

type PayOrderDto struct {
	Oid        string
	Price      int
	Pr         string
	VerifyUrl  string
	CreatedAt  int64
	PayStatus  int
	ExpireTime int64
}
