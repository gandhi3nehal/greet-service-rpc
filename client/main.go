package main

import (
	"context"
	"log"
	api "github.com/gandhi3nehal/greet-service/api/v1"
	grpc "google.golang.org/grpc"
)

func main(){
	cc, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer cc.Close()

	c := api.NewGreetServiceClient(cc)

	getGreeting("Jack", "us", c)
	getGreeting("Jose", "mx", c)
}

func getGreeting(name, countryCode string, c api.GreetServiceClient) {
	log.Println("creating greeting")

	res, err := c.Greet(context.Background(), &api.GreetRequest{CountryCode: countryCode, UserName: name})

	if err != nil {
		log.Println("error: ", err)
		panic(err)
	}

	log.Println(res.Result)

}

