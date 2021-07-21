package services

import (
	"fmt"
	"github.com/email-service/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/smtp"
	"os"
)

var Host = os.Getenv("EMAIL_HOST")
var Port = os.Getenv("EMAIL_PORT")
var User = os.Getenv("EMAIL_USER")
var Pass = os.Getenv("EMAIL_PASS")
var Auth = smtp.PlainAuth("", User, Pass, Host)
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
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: "+ r.Subject + "\n"
	msg := []byte(subject+mime+"\n"+r.Body)
	addr := Host+":"+Port

	if err := smtp.SendMail(addr, Auth, from, r.To, msg); err != nil{
		return false, err
	}
	return true, nil
}
type UserValReqEmail struct{
	models.RequestMail
	Body string `json:"body" binding:"required"`
}

func UserValidationEmailHandler(c *gin.Context)  {
	var req UserValReqEmail
	from = req.Sender

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	r := NewEmailReq(from, []string{req.Receiver}, req.Subject, req.Body)
	err := r.ParseTemplate("templates/template.html", req)
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


