package main

import (
	"context"
	"emailservice/conf"
	pb "emailservice/pivotstudio/email"
	"google.golang.org/grpc"
	"gopkg.in/gomail.v2"
	"log"
	"net"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedEmailServiceServer
}
var emailHost, emailFrom, emailFromPassword string
var emailPort int
func init(){
	config:=conf.ReadSettingsFromFile("Config.json")
	emailHost = config.EmailSenderSettings.Servername
	emailFrom =config.EmailSenderSettings.Email
	emailFromPassword =config.EmailSenderSettings.Password
	emailPort=config.EmailSenderSettings.Port
}

func sentEmail(email_to string,cc string, title string,content string)(err error){
	m := gomail.NewMessage()
	// 发邮件的地址
	m.SetHeader("From", emailFrom)
	// 给谁发送，支持多个账号
	m.SetHeader("To", email_to)
	// 抄送谁
	if len(cc)>0{
		m.SetAddressHeader("Cc", cc, "Dan")
	}
	// 邮件标题
	m.SetHeader("Subject", title)

	m.SetBody("text/html",content)
	// 附件
	//m.Attach("/home/Alex/lolcat.jpg")
	// stmp服务，端口号，发送邮件账号，发送账号密码
	d := gomail.NewDialer(emailHost,emailPort, emailFrom, emailFromPassword)
	// Send the email to Bob, Cora and Dan.
	if err = d.DialAndSend(m); err != nil {
		return err
	}else {
		return nil
	}
}

// SendEmail implements helloworld.GreeterServer
func (s *server) SendEmail(ctx context.Context, in *pb.SendEmailInfo) (*pb.ResponseInfo, error) {
	err:=sentEmail(in.ReceiveEmail,in.Cc,in.Title,in.Content)
	if err!=nil{
		return &pb.ResponseInfo{StatuCode:500,Message: "Failed to end email"}, err
	}else {
		return &pb.ResponseInfo{StatuCode:200,Message: "Send email successfully"},nil
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEmailServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}