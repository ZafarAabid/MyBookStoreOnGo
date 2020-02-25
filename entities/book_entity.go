package entities


type Book struct{
	Id int64 `json:"id"`
	Author string `json:"author"`
	BookName string `json:"title"`
	Image string `json:"image"`
	Quantity int `json:"quantity"`
	Price int `json:"price"`
	Description string `json:"intro"`
}

