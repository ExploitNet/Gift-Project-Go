// Package implement validation service client implementation
package implement

import (
	"log"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Client for validator service
func CreateValidatorService() ValidatorClient {
	conn, err := grpc.NewClient("159.69.181.5:50001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("не удалось подключиться: %v", err)
	}

	c := NewValidatorClient(conn)
	return c
}
