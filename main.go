package main

import (
	"context"
	"fmt"
	"main/database"
	"main/proto"
	"net"
	"os"

	"google.golang.org/grpc"
)

type server struct{}

var db database.Database

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err) // The port may be on use
	}
	srv := grpc.NewServer()
	databaseImplementation := os.Args[1]
	db, err = database.Factory(databaseImplementation)
	if err != nil {
		panic(err)
	}
	proto.RegisterBasicServiceServer(srv, &server{})
	fmt.Println("Prepare to serve")
	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}
func (s *server) Set(ctx context.Context, in *proto.SetRequest) (*proto.ServerResponse, error) {
	value, err := db.Set(in.GetKey(), in.GetValue())
	return generateResponse(value, err)
}
func (s *server) Get(ctx context.Context, in *proto.GetRequest) (*proto.ServerResponse, error) {
	value, err := db.Get(in.GetKey())
	return generateResponse(value, err)
}
func (s *server) Delete(ctx context.Context, in *proto.DeleteRequest) (*proto.ServerResponse, error) {
	value, err := db.Delete(in.GetKey())
	return generateResponse(value, err)
}
func generateResponse(value string, err error) (*proto.ServerResponse, error) {
	if err != nil {
		return &proto.ServerResponse{Success: false, Value: value, Error: err.Error()}, nil
	}
	return &proto.ServerResponse{Success: true, Value: value, Error: ""}, nil
}
