package models

type RequestMail struct { //
	Sender		 string		`json:"sender" binding:"required,email"`
	SenderName	 string		`json:"sender_name"`
	Receiver 	 string 	`json:"receiver" binding:"required,email"`
	ReceiverName string		`json:"receiver_name"  binding:"required"`
	Subject 	 string 	`json:"subject" binding:"required"`
	Title 		 string 	`json:"title"`
	Footer 		 string		`json:"footer"`
}
