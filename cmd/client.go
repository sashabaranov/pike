package main

import (
	"context"
	"github.com/sashabaranov/pike/backend"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer conn.Close()

	client := backend.NewBackendClient(conn)

	ctx := context.Background()

	createResponse, err := client.CreateUser(
		ctx,
		&backend.CreateUserRequest{
			Profile: &backend.UserProfile{
				Name:         "Sasha",
				PasswordHash: "",
				Age:          28,
			},
		},
	)
	log.Printf("createResponse=%v error=%v", createResponse, err)

	getResponse, err := client.GetUser(
		ctx,
		&backend.GetUserRequest{
			Id: createResponse.Profile.Id,
		},
	)
	log.Printf("getResponse=%v error=%v", getResponse, err)

	getResponse.Profile.Name = "Someone"
	updateResponse, err := client.UpdateUser(
		ctx,
		&backend.UpdateUserRequest{
			UpdatedProfile: getResponse.Profile,
		},
	)
	log.Printf("updateResponse=%v error=%v", updateResponse, err)

	deleteResponse, err := client.DeleteUser(
		ctx,
		&backend.DeleteUserRequest{
			Id: createResponse.Profile.Id,
		},
	)
	log.Printf("deleteResponse=%v error=%v", deleteResponse, err)
}
