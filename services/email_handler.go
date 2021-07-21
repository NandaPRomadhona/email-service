package services

import (
	"fmt"
	"github.com/email-service/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/smtp"
	"os"
)

var Host string
var Port string
var User string
var Pass string
var Auth smtp.Auth
var from string

func NewEmailReq(from string, to []string, subject string, body string) *Request {
	return &Request{
		From: from,
		To: to,
		Subject: subject,
		Body: body,
	}
}

func (r *Request) SendEmail() (bool, error) {
	Host = os.Getenv("EMAIL_HOST")
	Port = os.Getenv("EMAIL_PORT")
	User = os.Getenv("EMAIL_USER")
	Pass = os.Getenv("EMAIL_PASS")
	Auth = smtp.PlainAuth("", User, Pass, Host)

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: "+ r.Subject + "\n"
	msg := []byte(subject+mime+"\n"+r.Body)
	addr := Host+":"+Port

	if err := smtp.SendMail(addr, Auth, r.From, r.To, msg); err != nil{
		return false, err
	}
	return true, nil
}

type UserValReqEmail struct{
	models.RequestMail
	Body string `json:"body" binding:"required"`
	URL  string `json:"url" binding:"required"`
}

func UserValidation(c *gin.Context)  {
	var req UserValReqEmail
	// validation for request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	from = req.Sender
	r := NewEmailReq(from, []string{req.Receiver}, req.Subject, req.Body)
	err := r.ParseTemplate("templates/user-validation.html", req)
	if err == nil{
		ok, err := r.SendEmail()
		if err != nil{
			fmt.Println("Error Send Mail: ", err)
			c.JSON(http.StatusBadGateway, gin.H{
				"status":  "FAILED",
				"message": err,
			})
		}else{
			fmt.Println(ok)
			// Return response
			c.JSON(http.StatusOK, gin.H{
				"status": "SUCCESS",
				"message": "Email Sent!!!",
			})
		}
	}else {
		fmt.Println("Error Load Template: ", err)
		c.JSON(http.StatusBadGateway, gin.H{
			"status": "FAILED",
			"message": err,
		})
	}
}

type ReceiptReqEmail struct {
	models.RequestMail
	models.Order
}

func Receipt(c *gin.Context)  {
	var req ReceiptReqEmail
	// validation for request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	from = req.Sender
	r := NewEmailReq(from, []string{req.Receiver}, req.Subject, req.Title)
	err := r.ParseTemplate("templates/receipt.html", req)
	if err == nil{
		ok, err := r.SendEmail()
		if err != nil{
			fmt.Println("Error Send Mail: ", err)
			c.JSON(http.StatusBadGateway, gin.H{
				"status":  "FAILED",
				"message": err,
			})
		}else{
			fmt.Println(ok)
			// Return response
			c.JSON(http.StatusOK, gin.H{
				"status": "SUCCESS",
				"message": "Email Sent!!!",
			})
		}
	}else {
		fmt.Println("Error Load Template: ", err)
		c.JSON(http.StatusBadGateway, gin.H{
			"status": "FAILED",
			"message": err,
		})
	}
}

type ReminderReqEmail struct{
	models.RequestMail
	models.Reminder

}

func Reminder(c *gin.Context)  {
	var req ReminderReqEmail
	// validation for request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	from = req.Sender
	r := NewEmailReq(from, []string{req.Receiver}, req.Subject, req.Title)
	err := r.ParseTemplate("templates/reminder.html", req)
	if err == nil{
		ok, err := r.SendEmail()
		if err != nil{
			fmt.Println("Error Send Mail: ", err)
			c.JSON(http.StatusBadGateway, gin.H{
				"status":  "FAILED",
				"message": err,
			})
		}else{
			fmt.Println(ok)
			// Return response
			c.JSON(http.StatusOK, gin.H{
				"status": "SUCCESS",
				"message": "Email Sent!!!",
			})
		}
	}else {
		fmt.Println("Error Load Template: ", err)
		c.JSON(http.StatusBadGateway, gin.H{
			"status": "FAILED",
			"message": err,
		})
	}
	
}

type SendGiftReqEmail struct {
	models.RequestMail
	models.SendGift
}

func SendGift(c *gin.Context)  {

}

type ResetPassReqEmail struct {
	models.RequestMail
	models.ResetPassword
}

func ResetPassword(c *gin.Context)  {
	
}


