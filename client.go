package main

import (
	pb "Chitty-Chat_HW3_V2/chittychatpb"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var lamportTime int64

func main() {
	const maxMessageLength = 128

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewChittyChatClient(conn)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name: ")
	scanner.Scan()
	name := scanner.Text()

	// Join the chat
	clientTick(0)
	joinResp, err := client.JoinChat(context.Background(), &pb.Participant{Name: name, Timestamp: lamportTime})
	if err != nil {
		log.Fatalf("Could not join: %v", err)
	}
	log.Println(joinResp.Message)
	clientTick(joinResp.Timestamp)

	// Start a goroutine to receive messages from the server
	go func() {
		stream, err := client.BroadcastMessages(context.Background(), &pb.Empty{})
		if err != nil {
			log.Fatalf("Error receiving messages: %v", err)
		}

		for {
			msg, err := stream.Recv()
			if err != nil {
				log.Fatalf("Error receiving message: %v", err)
			}
			clientTick(msg.Timestamp)
			log.Printf("[Broadcast] %s: %s", msg.Participant, msg.Message)
		}
	}()

	// Handle sending messages
	for scanner.Scan() {
		text := scanner.Text()
		if text == "quit" {
			break
		}
		if len(text) > maxMessageLength {
			log.Printf("Message too long, max %d characters\n", maxMessageLength)
			continue
		}

		clientTick(0)
		_, err = client.PublishMessage(context.Background(), &pb.ChatMessage{
			Participant: name,
			Message:     text,
			Timestamp:   lamportTime,
		})
		if err != nil {
			log.Fatalf("Could not publish message: %v", err)
		}
	}

	// Leave the chat
	leaveResp, err := client.LeaveChat(context.Background(), &pb.Participant{Name: name, Timestamp: lamportTime})
	if err != nil {
		log.Fatalf("Could not leave: %v", err)
	}
	clientTick(leaveResp.Timestamp)
	log.Println(leaveResp.Message)
}

func clientTick(recivedTime int64) int64 {
	tempTime := lamportTime
	if recivedTime > lamportTime {
		lamportTime = recivedTime
	}
	lamportTime++
	log.Printf("Client lamport time from: %d to %d", tempTime, lamportTime)
	return lamportTime
}
