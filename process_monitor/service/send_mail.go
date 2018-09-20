package service

import (
	"gopkg.in/gomail.v2"
	"github.com/sirupsen/logrus"
	"sub_account_service/process_monitor/config"
)

func SendMail(sub string,msg string,receivers []string)  {
	sender := config.GetConfigInstance().EmailSender
	d := gomail.NewDialer(sender.Host, sender.Port, sender.Username, sender.Pwd)
	s,err := d.Dial()
	if err != nil {
		logrus.WithError(err).Errorln("init mail server failed")
		return
	}
	from := sender.Username

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", receivers[0])
	m.SetHeader("Subject", sub)
	m.SetBody("text/plain", msg)

	err=s.Send(from, receivers, m)
	if err!=nil{
		logrus.WithError(err).Errorln("send mail error")
		return
	}
}
