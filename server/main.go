package main

import (
	"context"
	pb "emailservice/pivotstudio/email"
	"github.com/Pivot-Studio/Authorization-Template/pkg/ATutil"
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
var email_host,email_from,email_from_password string
func init(){
	config:=ATutil.ReadSettingsFromFile("Config.json")
	email_host= config.EmailSenderSettings.Servername
	email_from=config.EmailSenderSettings.Email
	email_from_password=config.EmailSenderSettings.Password
}

func sentEmail(email_to string,cc string, title string,content string)(err error){
	m := gomail.NewMessage()
	// 发邮件的地址
	m.SetHeader("From", email_from)
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
	d := gomail.NewDialer(email_host,25,email_from,email_from_password)
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