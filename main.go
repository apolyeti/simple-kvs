package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

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
	store := kvs.New()

	// HTTP Handlers
	http.HandleFunc("/put", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		valueStr := r.URL.Query().Get("value")

		value, err := strconv.Atoi(valueStr)
		if err != nil {
			http.Error(w, "Invalid value", http.StatusBadRequest)
			return
		}

		store.Set(key, value)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Key-Value pair added: %s = %d", key, value)
	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")

		value, err := store.Get(key)
		if err != nil {
			http.Error(w, "Key not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Value: %d", value)
	})

	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")

		store.Delete(key)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Key-Value pair deleted: %s", key)
	})

	// Start HTTP server in a goroutine
	go func() {
		log.Println("Starting HTTP server on :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("failed to start HTTP server: %v", err)
		}
	}()

	// Start gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterKvsServer(s, &server{store: store})
	log.Printf("Starting gRPC server on %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
