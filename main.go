package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/apolyeti/simple-kvs/kvs"
	pb "github.com/apolyeti/simple-kvs/proto"
)

type server struct {
	pb.UnimplementedKvsServer
	store *kvs.Store
}

func (s *server) Set(ctx context.Context, in *pb.SetRequest) (*pb.SetResponse, error) {
	s.store.Set(in.Key, int(in.Value))
	return &pb.SetResponse{}, nil
}

func (s *server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	value, err := s.store.Get(in.Key)
	if err != nil {
		return nil, err
	}
	return &pb.GetResponse{Value: int32(value)}, nil
}

func (s *server) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	s.store.Delete(in.Key)
	return &pb.DeleteResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	store := kvs.New()
	pb.RegisterKvsServer(s, &server{store: store})
	log.Printf("Starting server on %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
