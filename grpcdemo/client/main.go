package main

import (
	"flag"
	"fmt"
	"grpcdemo/pb"
	"io"
	"log"
	"os"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

const port = ":9000"

func main() {

	option := flag.Int("o", 1, "command to run")
	flag.Parse()

	creds, err := credentials.NewClientTLSFromFile("cert.pem", "")
	if err != nil {
		log.Fatal(err)
	}

	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}

	conn, err := grpc.Dial("localhost"+port, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewEmployeeServiceClient(conn)

	switch *option {
	case 1:
		SendMetadata(client)
	case 2:
		GetByBadgeNumber(client)
	case 3:
		GetAll(client)
	case 4:
		AddPhoto(client)
	case 5:
		SaveAll(client)
	}
}

//SaveAll comments
func SaveAll(client pb.EmployeeServiceClient) {
	employees := []pb.Employee{
		pb.Employee{
			BadgeNumber:         124,
			FirstName:           "testFn4",
			LastName:            "testLn4",
			VacationAccrualRate: 1.6,
			VacationAccrued:     2.3,
		},
		pb.Employee{
			BadgeNumber:         125,
			FirstName:           "testFn5",
			LastName:            "testLn5",
			VacationAccrualRate: 1.5,
			VacationAccrued:     2.5,
		},
	}

	reqStream, err := client.SaveAll(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	doneChan := make(chan struct{})
	go func() {
		for {
			res, err := reqStream.Recv()
			if err == io.EOF {
				doneChan <- struct{}{}
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("res :", res.Employee)
		}
	}()

	for _, e := range employees {
		err := reqStream.Send(&pb.EmployeeRequest{Employee: &e})
		if err != nil {
			log.Fatal(err)
		}
	}
	reqStream.CloseSend()
	<-doneChan
}

//AddPhoto comments
func AddPhoto(client pb.EmployeeServiceClient) {
	f, err := os.Open("test.JPG")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	md := metadata.New(map[string]string{"badgenumber": "2080"})
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, md)

	reqStream, err := client.AddPhoto(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for {
		chunk := make([]byte, 64*1024)
		n, err := f.Read(chunk)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if n < len(chunk) {
			chunk = chunk[:n]
		}

		reqStream.Send(&pb.AddPhotoRequest{Data: chunk})
	}

	err = reqStream.CloseSend()
	if err != nil {
		log.Fatal(err)
	}

	res, err := reqStream.Recv()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("res: ", res.IsOk)
}

//GetAll comments
func GetAll(client pb.EmployeeServiceClient) {
	resStream, err := client.GetAll(context.Background(), &pb.GetAllRequest{})
	if err != nil {
		log.Fatal(err)
	}

	for {
		res, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res)
	}
}

//GetByBadgeNumber comments
func GetByBadgeNumber(client pb.EmployeeServiceClient) {
	res, err := client.GetByBadgeNumber(context.Background(), &pb.GetByBadgeNumberRequest{BadgeNumber: 2080})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("res :", res)
	fmt.Println("res :", res.Employee)
}

//SendMetadata test comment
func SendMetadata(client pb.EmployeeServiceClient) {
	md := metadata.MD{}
	md["user"] = []string{"testuser"}
	md["password"] = []string{"passcode1"}
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, md)
	client.GetByBadgeNumber(ctx, &pb.GetByBadgeNumberRequest{})
}
