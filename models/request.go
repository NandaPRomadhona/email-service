package models

import (
	"time"
)

type RequestMail struct { //
	Sender		 string		`json:"sender" binding:"required,email"`
	SenderName	 string		`json:"sender_name"`
	Receiver 	 string 	`json:"receiver" binding:"required,email"`
	ReceiverName string		`json:"receiver_name"  binding:"required"`
	Subject 	 string 	`json:"subject" binding:"required"`
	Title 		 string 	`json:"title"`
	Footer 		 string		`json:"footer"`
}

type Item struct {
	ID 		int32	`json:"id"`
	Name 	string 	`json:"name"`
	Price 	float32	`json:"price"`
	ImagePath 	string	`json:"image_path"`
}

type Order struct{
	OrderNumber 	int32		`json:"order_number"`
	TransactionDate	time.Time	`json:"transaction_date"`
	PaymentMethod	string		`json:"payment_method"`
	Items			[]Item		`json:"items"`
	CurrencyCode	string		`json:"currency_code"`
	OrderTotal		float32		`json:"order_total"`
	Discount		float32		`json:"discount"`
	GrandTotal 		float32		`json:"grand_total"`
}

type Reminder struct{
	Body 	string	`json:"body"`
	InfoURL string	`json:"info_url"`
	Footer	string	`json:"footer"`
	Items	[]Item	`json:"items"`
}

type SendGift struct {
	Body 	string	`json:"body"`
	Items 	[]Item 	`json:"items"`
	Code 	string	`json:"code"`
}

type ResetPassword struct {
	Body 	string	`json:"body"`
	URL 	string	`json:"url"`
}
