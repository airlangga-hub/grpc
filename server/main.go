package main

import (
	"context"
	"errors"
	"log"
	"net"

	pb "github.com/airlangga-hub/grpc/proto"
	"google.golang.org/grpc"
)

type Person struct {
	ID				int32
	Name			string
	Email			string
	PhoneNumber		string
}

var nextID int32 = 1
var persons = map[int32]*Person{}

type server struct {
	pb.UnimplementedPersonServiceServer
}

func (s *server) Create(ctx context.Context, in *pb.CreatePersonRequest) (*pb.PersonProfileResponse, error) {
	
	if in.Name == "" || in.Email == "" || in.PhoneNumber == "" {
		return &pb.PersonProfileResponse{}, errors.New("fields missing!")
	}
	
	person := Person{
		ID: nextID,
		Name: in.Name,
		Email: in.Email,
		PhoneNumber: in.PhoneNumber,
	}
	
	persons[person.ID] = &person
	nextID += 1
	
	return &pb.PersonProfileResponse{
		Id: person.ID,
		Name: person.Name,
		Email: person.Email,
		PhoneNumber: person.PhoneNumber,
	}, nil
}

func (s *server) Read(ctx context.Context, in *pb.SinglePersonRequest) (*pb.PersonProfileResponse, error) {
	
	person, ok := persons[in.Id]
	if !ok {
		return &pb.PersonProfileResponse{}, errors.New("person not found!")
	}
	
	return &pb.PersonProfileResponse{
		Id: person.ID,
		Name: person.Name,
		Email: person.Email,
		PhoneNumber: person.PhoneNumber,
	}, nil
}

func (s *server) Update(ctx context.Context, in *pb.UpdatePersonRequest) (*pb.SuccessResponse, error) {
	
	if in.Name == "" || in.Email == "" || in.PhoneNumber == "" {
		return &pb.SuccessResponse{Response: "fields missing!"}, errors.New("fields missing!")
	}
	
	person, ok := persons[in.Id]
	if !ok {
		return &pb.SuccessResponse{Response: "person not found!"}, errors.New("person not found!")
	}
	
	person.Name = in.Name
	person.Email = in.Email
	person.PhoneNumber = in.PhoneNumber
	
	return &pb.SuccessResponse{Response: "update success!"}, nil
}

func (s *server) Delete(ctx context.Context, in *pb.SinglePersonRequest) (*pb.SuccessResponse, error) {
	
	person, ok := persons[in.Id]
	if !ok || person.ID == 0 {
		return &pb.SuccessResponse{Response: "person not found!"}, errors.New("person not found!")
	}
	
	delete(persons, person.ID)
	
	return &pb.SuccessResponse{Response: "delete success!"}, nil
}

func main() {
	
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	s := grpc.NewServer()
	pb.RegisterPersonServiceServer(s, &server{})
	
	log.Println("gRPC server listening at: ", lis.Addr())
	
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}