package services

import (
	"bytes"
	"html/template"
)

type Request struct{
	From 	string
	To 		[]string
	Subject string
	Body 	string
}

func (r *Request) ParseTemplate (fileName string, data interface{}) error {
	t, err := template.ParseFiles(fileName)
	if err != nil{
		return  err
	}
	buf := new(bytes.Buffer)

	if err = t.Execute(buf, data); err != nil{
		return err
	}
	r.Body = buf.String()
	return nil
}

