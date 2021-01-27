package main

import (
	"context"
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
	r, err := c.SendEmail(ctx, &pb.SendEmailInfo{ReceiveEmail: "1743432766@qq.com",Title: "GRPC practice",
		Content:"TIhs is a email sent by grpc server",Cc: "1720648723@qq.com"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
