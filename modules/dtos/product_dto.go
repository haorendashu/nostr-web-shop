package dtos

type ProductDto struct {
	Id        string
	Pubkey    string
	UpdatedAt int64
	CreatedAt int64
	Status    int
	Name      string
	Imgs      string // images join ,
	Des       string
	Content   string // html content
	Price     int    // milisats, sats num * 1000
	Lnwallet  string

	Skus []*ProductDetailDto
}

type ProductDetailDto struct {
	Code  string
	Name  string
	Price int // milisats, sats num * 1000
	Stock int
}
