package domain

type Person struct {
	Name    *string `json:"name"`
	Surname *string `json:"surname"`
	Age     string  `json:"age"`
	Addres  addres  `json:"addres"`
	Phone   *string `json:"phone"`
}

type addres struct {
	Country   *string `json:"country"`
	City      *string `json:"city"`
	Street    *string `json:"street"`
	House     *string `json:"house"`
	Apartment *string `json:"apartment"`
}

func NewPersonStorage() []Person {
	return []Person{}
}
