package main

import (
	"context"
	"log"
	"net"
	
	pb "github.com/airlangga-hub/grpc/coffeeshop_proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCoffeeShopServer
}

func (s *server) GetMenu(menuRequest *pb.MenuRequest, srv pb.CoffeeShop_GetMenuServer) error {
	
	items := []*pb.Item{
		{
			Id: 1,
			Name: "Black Coffee",
		},
		{
			Id: 2,
			Name: "Americano",
		},
		{
			Id: 3,
			Name: "Vanilla Latte",
		},
	}
	
	for i, _ := range items {
		srv.Send(&pb.Menu{
			Items: items[:i+1],
		})
	}
	
	return nil
}