package main

import (
	pb "BSDISYS1KU-En-GO-gruppe_Hand-in_5/biddybidderpb"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.NewClient("127.0.0.2:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	defer conn.Close()

	client := pb.NewAuctionClient(conn)
	_, err = client.Ping(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	message, err := client.StartFunction(context.Background(), &pb.Time{Time: 0})
	log.Println(message)
}
