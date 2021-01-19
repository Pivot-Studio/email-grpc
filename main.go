package main

import (
	"github.com/Pivot-Studio/Authorization-Template/pkg/ATutil"
	"gopkg.in/gomail.v2"
)

type EmailSetting_Sender struct {
	Username 	string
	Password 	string
	Host 	 	string
}
type EmailSetting_Receiver struct{
	Username 	string
	header 		string
	Content		string
}
func EmailSender(){
	from_config:=readEnailSettings_Sender()
	//在下一行填入receiver和content
	to_config:=readEnailSettings_Receiver("1720648723@qq.com","nihao","123")
	email_from:=from_config.Username
	email_from_password:=from_config.Password
	email_host:=from_config.Host
	email_to:=to_config.Username
	email_to_header:=to_config.header
	content:=to_config.Content
	sentEmail(email_from,email_from_password,email_host,email_to,email_to_header,content)
}
func readEnailSettings_Sender()EmailSetting_Sender{
	email_config:=ATutil.ReadSettingsFromFile("Config.json")
	email_username:=email_config.EmailSenderSettings.Email
	email_password:=email_config.EmailSenderSettings.Password
	email_host:=email_config.EmailSenderSettings.Servername
	email_settings:=EmailSetting_Sender{
		Username:email_username,
		Password: email_password,
		Host: email_host,
	}
	return email_settings
}
func readEnailSettings_Receiver(receiver string,header string,content string)  EmailSetting_Receiver{
	a:=EmailSetting_Receiver{
		Username: receiver,
		header:header,
		Content: content,
	}
	return a
}
func sentEmail(email_from string,email_from_password string,email_host string,email_to string,email_to_header string,content string){
	m := gomail.NewMessage()
	// 发邮件的地址
	m.SetHeader("From", email_from)
	// 给谁发送，支持多个账号
	m.SetHeader("To", email_to)
	// 抄送谁
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	// 邮件标题
	m.SetHeader("Subject", email_to_header)
	// 邮件正文，支持 html
	//这里等着前端给我写html啦
	m.SetBody("text/html", content)
	// 附件
	//m.Attach("/home/Alex/lolcat.jpg")
	// stmp服务，端口号，发送邮件账号，发送账号密码
	d := gomail.NewDialer(email_host,25,email_from,email_from_password)
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func main(){
	EmailSender()
}