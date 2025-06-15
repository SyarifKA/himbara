package dtos

type User struct {
	PhoneNumber string `form:"phoneNumber"`
	// Bank        string `form:"bank"`
}

type PurchaseOrder struct {
	PhoneNumber   string `form:"phoneNumber"`
	PaymentChanel string `form:"paymentChanel"`
	ProductId     string `form:"productId"`
	ProductName   string `form:"productName"`
	Amount        int64  `form:"amount"`
	Status        string `gorm:"default:'waiting'"`
}

type CheckoutPayload struct {
	ProductId     string
	ProductName   string
	PaymentChanel string
	Amount        int64
}
