package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/carousell/chope-assignment/proto"
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
}

func testLogin(c pb.UserLoginClient) {
	req := new(pb.LogInRequest)
	req.Username = ""
	req.Password = ""
	resp, err := c.Login(ctx, req)
	if err != nil {
		fmt.Println("err-heree", err)
	}
	fmt.Println("respt", resp.Token)
}
