package main

import (
	"context"
	"log"
	"time"

	pb "github.com/macduyhai/grpcApi/apiproto"

	"google.golang.org/grpc"
)

const (
	name = "Anna"
)

func main() {
	log.Println("Client proto starting...")
	// flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial("0.0.0.0:8899", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	} else {
		log.Println("Connection done")
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	// // Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Add(ctx, &pb.Addrequest{Username: "SimonHD", Password: "123123", Fullname: "Mac Hung", Salary: 12345})
	if err != nil {
		log.Print(r)
		log.Printf("\n%v\n", err)
	}
	log.Printf("Result: %s", r.GetMessage())
}
