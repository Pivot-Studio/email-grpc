package main

import (
	"bytes"
	"context"
	"html/template"
	"google.golang.org/grpc"
	pb "emailservice/pivotstudio/email"
	"log"
	"time"
)
const (
	address     = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEmailServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()


	t, err := template.ParseFiles("reset_email.html")

	if err != nil {
		log.Panic(err)
	}
	buffer := new(bytes.Buffer)
	var data interface{}
	if err = t.Execute(buffer, data); err != nil {
		log.Panic(err)
	}
	r, err := c.SendEmail(ctx, &pb.SendEmailInfo{ReceiveEmail: "xieyuschen@gmail.com",Title: "GRPC practice",
		Content:buffer.String()})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
