package services

import "github.com/Pivot-Studio/Authorization-Template/pkg/ATutil"

type EmailSetting struct {
	Username 	string
	Password 	string
	Host 	 	string
}
func SendEmail(){
	ReadEnailSettings()
}
func ReadEnailSettings(){
	email_config:=ATutil.ReadSettingsFromFile("Config.json")
	email_username:=email_config.EmailSenderSettings.Email
	email_password:=email_config.EmailSenderSettings.Password
	email_host:=email_config.EmailSenderSettings.Servername
	a:=EmailSetting{
		Username:email_username,
		Password: email_password,
		Host: email_host,
	}
	
}

