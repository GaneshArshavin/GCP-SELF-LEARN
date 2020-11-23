package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/carousell/gcp-self-study/proto"
	"google.golang.org/grpc"
)

const (
	address     = "0.0.0.0:8080"
	httpAddress = "http://127.0.0.1:8080/"
)

var ctx = context.Background()

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		fmt.Println("err")
	}
	defer conn.Close()
	c := pb.NewUserLoginClient(conn)
	fmt.Printf("c = %+v\n", c)
	testLogin(c)
	//testRegister(c)
}

func testLogin(c pb.UserLoginClient) {
	req := new(pb.LogInRequest)
	req.Username = "Ganesh"
	req.Password = "Welcome135"
	resp, err := c.Login(ctx, req)
	if err != nil {
		fmt.Println("err-heree", err)
	}
	fmt.Println("respt", resp.Token)
}

func testRegister(c pb.UserLoginClient) {
	req := new(pb.RegisterRequest)
	req.Username = "Ganesh"
	req.Password = "Welcome135"
	req.Email = "ganesharshavin@gmail.com"
	req.ApiKey = "asas"
	req.Secret = "qwwqd"
	resp, err := c.Register(ctx, req)
	if err != nil {
		fmt.Println("err-heree", err)
	}
	fmt.Println("respt", resp.Token)
}
