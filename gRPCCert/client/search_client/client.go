package main

import (
	"context"
	"log"

	pb "github.com/jie1090/GoTutorial/gRPCCert/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const PORT = "9002"

func main() {
	c, err := credentials.NewClientTLSFromFile("../../conf/server.pem", "gRPCCert")
	if err != nil {
		log.Fatalf("credentials.NewClientTLSFromFile err: %v", err)
	}
	conn, err := grpc.Dial(":"+PORT, grpc.WithTransportCredentials(c))
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()
	client := pb.NewSearchServiceClient(conn)
	resp, err := client.Search(context.Background(), &pb.SearchRequest{
		Request: "gRPC",
	})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}
	log.Printf("resp: %s", resp.GetResponse())
}
