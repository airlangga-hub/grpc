package main

import (
	"context"
	"fmt"
	pb "github.com/airlangga-hub/grpc/proto"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	
	// connect
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	
	conn, err := grpc.NewClient("localhost:8080", opts...)
	if err != nil {
		log.Fatalln("failed to connect to server: ", err)
	}
	
	defer conn.Close()
	
	client := pb.NewPersonServiceClient(conn)
	
	// timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	
	// CREATE
	fmt.Println("creating new person....")
	createReq := &pb.CreatePersonRequest{
		Name: "John Wick",
		Email: "john@wick.com",
		PhoneNumber: "12345-6789",
	}
	
	createRes, err := client.Create(ctx, createReq)
	if err != nil {
		log.Fatalln("error creating: ", err)
	}
	
	fmt.Printf("person created: %+v\n", createRes)
	
	// READ
	fmt.Println("reading person by ID....")
	readReq := &pb.SinglePersonRequest{Id: createRes.Id}
	readRes, err := client.Read(ctx, readReq)
	if err != nil {
		log.Fatalln("error reading: ", err)
	}
	fmt.Printf("person read: %+v\n", readRes)
	
	// UPDATE
	fmt.Println("updating person....")
	updateReq := &pb.UpdatePersonRequest{
		Id: createRes.Id,
		Name: "Luke Skywalker",
		Email: "luke@skywalker.com",
		PhoneNumber: "9032-82133",
	}
	
	updateRes, err := client.Update(ctx, updateReq)
	if err != nil {
		log.Fatalln("error updating: ", err)
	}
	
	fmt.Printf("person updated: %s\n", updateRes.Response)
	
	// DELETE
	fmt.Println("deleting person....")
	
	deleteReq := &pb.SinglePersonRequest{Id: createRes.Id}
	deleteRes, err := client.Delete(ctx, deleteReq)
	if err != nil {
		log.Fatalln("error deleting: ", err)
	}
	
	fmt.Printf("person deleted: %s\n", deleteRes.Response)
}