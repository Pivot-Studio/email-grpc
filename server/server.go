package server

import (
	"context"
	"emailservice/conf"
	"emailservice/logrus"
	pb "emailservice/pivotstudio/email"
	"gopkg.in/gomail.v2"
)

// Server is used to implement helloworld.GreeterServer.
type Server struct {
	pb.UnimplementedEmailServiceServer
}

var emailHost, emailFrom, emailFromPassword string
var emailPort int

func init() {
	config := conf.ReadSettingsFromFile("Config.json")
	emailHost = config.EmailSenderSettings.Servername
	emailFrom = config.EmailSenderSettings.Email
	emailFromPassword = config.EmailSenderSettings.Password
	emailPort = config.EmailSenderSettings.Port
}

func sentEmail(emailTo string, cc string, title string, content string) (err error) {
	m := gomail.NewMessage()
	// 发邮件的地址
	m.SetHeader("From", emailFrom)
	// 给谁发送，支持多个账号
	m.SetHeader("To", emailTo)
	// 抄送谁
	if len(cc) > 0 {
		m.SetAddressHeader("Cc", cc, "Dan")
	}
	// 邮件标题
	m.SetHeader("Subject", title)

	m.SetBody("text/html", content)
	// 附件
	//m.Attach("/home/Alex/lolcat.jpg")
	// stmp服务，端口号，发送邮件账号，发送账号密码
	d := gomail.NewDialer(emailHost, emailPort, emailFrom, emailFromPassword)
	// Send the email to Bob, Cora and Dan.
	if err = d.DialAndSend(m); err != nil {
		logrus.Log.WithFields(map[string]interface{}{
			"Dial error": err,
		})
		return err
	} else {
		logrus.Log.WithFields(map[string]interface{}{
			"Host": emailHost,
			"Port": emailPort,
			"From": emailFrom,
			"To":   emailTo,
		}).Info()
		return nil
	}
}

// SendEmail implements helloworld.GreeterServer
func (s *Server) SendEmail(ctx context.Context, in *pb.SendEmailInfo) (*pb.ResponseInfo, error) {
	err := sentEmail(in.ReceiveEmail, in.Cc, in.Title, in.Content)
	if err != nil {
		logrus.Log.WithFields(map[string]interface{}{
			"send error": err,
			"statuscode": 500,
		}).Error()
		return &pb.ResponseInfo{StatuCode: 500, Message: "Failed to send email"}, err
	} else {
		logrus.Log.WithFields(map[string]interface{}{
			"statuscode": 200,
		}).Info()
		return &pb.ResponseInfo{StatuCode: 200, Message: "Send email successfully"}, nil
	}
}
