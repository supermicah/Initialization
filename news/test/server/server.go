package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/supermicah/Protobufs/news/micah/wiki/news"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedSearchServiceServer
}

// Search implements helloworld.GreeterServer
func (s *server) Search(ctx context.Context, in *pb.NewsRequest) (*pb.NewsResponse, error) {
	log.Printf("Received: %v", in.GetName())
	newsList := make([]*pb.News, 0)
	newsList = append(newsList, &pb.News{
		Id:         1,
		Title:      "test",
		Keyword:    "test",
		Desc:       "test",
		HappenTime: &timestamppb.Timestamp{Seconds: 1, Nanos: 1},
		State:      0,
		CreateTime: &timestamppb.Timestamp{Seconds: 2, Nanos: 2},
		UpdateTime: &timestamppb.Timestamp{Seconds: 3, Nanos: 3},
	})
	return &pb.NewsResponse{NewsList: newsList}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSearchServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
