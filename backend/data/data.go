package data

//fields of User

type User struct {
	UserId   int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Usertype string `json:"usertype"`
}

//fields of customer

type Product struct {
	ProductID     int     `json:"product_id"`
	ProductName   string  `json:"name"`
	Price         float64 `json:"price"`
	Specification string  `json:"specification"`
	Image         string  `json:"image"`
	ImagePath     string  `json:"image_path"`
}

//to store the extracted data of login

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
