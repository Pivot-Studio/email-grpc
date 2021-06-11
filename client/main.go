package main

import (
	"bytes"
	"context"
	"emailservice/logrus"
	pb "emailservice/pivotstudio/email"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"html/template"
	"log"
	"time"
)

//const (
//	address     = "localhost:50051"
//)

func main() {
	// 测试参数放在命令行中
	address := flag.String("addr", "localhost:50051", "地址")
	email := flag.String("email", "root@yourmail", "邮箱")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logrus.Log.WithFields(map[string]interface{}{
			"connect error": err,
		}).Fatal()
		//log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEmailServiceClient(conn)

	//设置服务链接时间
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	t, err := template.ParseFiles("reset_email.html")

	if err != nil {
		logrus.Log.WithFields(map[string]interface{}{
			"ParseFile error": err,
		}).Panic()
		//log.Panic(err)
	}
	buffer := new(bytes.Buffer)
	var data interface{}
	if err = t.Execute(buffer, data); err != nil {
		log.Panic(err)
	}

	fmt.Println(*email)
	r, err := c.SendEmail(ctx, &pb.SendEmailInfo{ReceiveEmail: *email, Title: "GRPC practice",
		Content: buffer.String()})
	if err != nil {
		logrus.Log.WithFields(map[string]interface{}{
			"context error": err,
		}).Fatal()
		//log.Fatalf("could not greet: %v", err)
	}
	logrus.Log.WithFields(map[string]interface{}{
		"Greeting info": r.GetMessage(),
	}).Info()
	//log.Printf("Greeting: %s", r.GetMessage())
}
