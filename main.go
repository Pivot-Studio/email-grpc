package main

import (
	"emailservice/logrus"
	pb "emailservice/pivotstudio/email"
	"emailservice/server"
	"google.golang.org/grpc"
	"net"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		logrus.Log.WithFields(map[string]interface{}{
			"Listening error": err,
		}).Fatal()
		//log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEmailServiceServer(s, &server.Server{})
	if err := s.Serve(lis); err != nil {
		logrus.Log.WithFields(map[string]interface{}{
			"serve error": err,
		})
		//log.Fatalf("failed to serve: %v", err)
	}
}
