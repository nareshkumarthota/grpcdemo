package main

import (
	"errors"
	"fmt"
	"grpcdemo/pb"
	"io"
	"log"
	"net"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"golang.org/x/net/context"
)

const port = ":9000"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	creds, err := credentials.NewServerTLSFromFile("cert.pem", "key.pem")
	if err != nil {
		log.Fatal(err)
	}

	opts := []grpc.ServerOption{grpc.Creds(creds)}

	s := grpc.NewServer(opts...)

	pb.RegisterEmployeeServiceServer(s, new(employeeService))

	log.Println("Starting server on port: ", port)

	s.Serve(lis)
}

type employeeService struct{}

func (s *employeeService) GetByBadgeNumber(ctx context.Context, req *pb.GetByBadgeNumberRequest) (*pb.EmployeeResponse, error) {

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		fmt.Printf("metadata  received : %v\n", md)
	}

	for _, e := range employees {
		if e.BadgeNumber == req.BadgeNumber {
			return &pb.EmployeeResponse{Employee: &e}, nil
		}
	}

	return nil, errors.New("Employee not found")
}

func (s *employeeService) GetAll(req *pb.GetAllRequest, stream pb.EmployeeService_GetAllServer) error {
	for _, e := range employees {
		stream.Send(&pb.EmployeeResponse{Employee: &e})
	}
	return nil
}

func (s *employeeService) Save(ctx context.Context, req *pb.EmployeeRequest) (*pb.EmployeeResponse, error) {
	return nil, nil
}

func (s *employeeService) SaveAll(stream pb.EmployeeService_SaveAllServer) error {

	for {
		emp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		employees = append(employees, *emp.Employee)
		stream.Send(&pb.EmployeeResponse{Employee: emp.Employee})
	}

	for _, e := range employees {
		fmt.Println(e)
	}

	return nil
}

func (s *employeeService) AddPhoto(stream pb.EmployeeService_AddPhotoServer) error {

	md, ok := metadata.FromIncomingContext(stream.Context())
	if ok {
		fmt.Printf("recevied photo request for badge number %v \n", md["badgenumber"][0])
	}

	imageData := []byte{}

	for {
		data, err := stream.Recv()
		if err == io.EOF {
			fmt.Printf("filereceived with length %v \n", len(imageData))
			return stream.Send(&pb.AddPhotoResponse{IsOk: true})
		}
		if err != nil {
			return err
		}
		fmt.Printf("Received %v bytes \n", len(data.Data))
		imageData = append(imageData, data.Data...)
	}
}
