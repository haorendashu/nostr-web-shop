package dtos

type PayOrderDto struct {
	Id         string
	Oid        string
	Price      int
	Pr         string
	VerifyUrl  string
	CreatedAt  int64
	PayStatus  int
	ExpireTime int64
}
