package main

import (
	"context"
	pb "dbclient/customer"
	"dbclient/interfaces"
	"dbclient/models"
	"fmt"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type CustomerServiceServer struct {
	mu        sync.Mutex
	BankingService interfaces.ICustomer
	pb.UnimplementedCustomerServiceServer
}



func InitCustomerController(mu sync.Mutex, profileService interfaces.ICustomer,server) CustomerServiceServer {
	return CustomerServiceServer{profileService}
}

func (s *CustomerServiceServer) AddCustomer(ctx context.Context, req *pb.Customer) (*pb.CustomerResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	req.Name="billa"
	req.Balance=10000.0
	// s.Customers[CustomerId] = req
	// return &pb.CustomerResponse{Id: CustomerId}, nil
	var customer *models.Customer
	customer.Name=req.Name
	customer.Balance=float64(req.Balance)
	newProfile, err := s.BankingService.AddCustomer(customer)

}

func (s *CustomerServiceServer) GetCustomers(ctx context.Context, req *pb.Empty) (*pb.CustomerList, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	Customers := make([]*pb.Customer, 0, len(s.Customers))
	for _, Customer := range s.Customers {
		Customers = append(Customers, Customer)
	}
	return &pb.CustomerList{Customers: Customers}, nil
}

func generateID() string {
	return "CustomerID"
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Println("failed to listen :%v", err)
		return
	}
	server := grpc.NewServer()
	pb.RegisterCustomerServiceServer(server, &CustomerServiceServer{
		Customers: make(map[string]*pb.Customer),
	})
	fmt.Println("server listening on :50051")
	if err := server.Serve(lis); err != nil {
		fmt.Println("Failed to serve: %v", err)
	}
}
