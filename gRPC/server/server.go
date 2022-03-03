package main

import (
	"context"
	"log"
	"net"

	pb "github.com/jie1090/GoTutorial/gRPC/proto"
	"google.golang.org/grpc"
)

type SearchService struct{}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

const PORT = "9002"

func main() {
	server := grpc.NewServer()
	pb.RegisterSearchServiceServer(server, &SearchService{})
	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	server.Serve(lis)
}
