package dtos

type OrderAddDto struct {
	Comment string
	Skus    []*OrderAddProductDto
}

type OrderAddProductDto struct {
	DetailId string
	Num      int
}
