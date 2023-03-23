package main

import (
	"context"
	"fmt"
	api "github.com/gandhi3nehal/greet-service/api/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	api.UnimplementedGreetServiceServer
}

func (s server) Greet(ctx context.Context, req *api.GreetRequest) (*api.GreetResponse, error) {

	log.Println("User name: ", req.UserName)
	log.Println("Country code: ", req.CountryCode)

	var greeting string

	switch req.CountryCode {
	case "us":
		greeting = "hello " + req.UserName
	case "mx":
		greeting = "Hola " + req.UserName
	default:
		greeting = "Hola/Hello " + req.UserName
	}

	return &api.GreetResponse{
		Result: greeting,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	fmt.Println("starting server")

	s := grpc.NewServer()
	srv := server{}
	api.RegisterGreetServiceServer(s, &srv)

	if err := s.Serve(listener); err != nil {
		panic(err)
	}

}
