package main

import (
	"context"
	"fmt"
	user "food-delivery-service/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

type server struct{}

func (s *server) GetUserByIds(ctx context.Context, request *user.UserRequest) (*user.UserResponse, error) {
	log.Println(request.UserIds)

	return &user.UserResponse{
		Users: []*user.User{
			{
				Id:        "1",
				FirstName: "Viet",
				LastName:  "Tran",
				Role:      "user",
			},
		},
	}, nil
}

func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Error %v", err)
	}

	fmt.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()

	user.RegisterUserServiceServer(s, &server{})

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalln(err)
		}
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:50051",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = user.RegisterUserServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":3000",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:3000")
	log.Fatalln(gwServer.ListenAndServe())
}