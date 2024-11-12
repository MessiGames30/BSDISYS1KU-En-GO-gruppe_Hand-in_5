package main

import (
	pb "BSDISYS1KU-En-GO-gruppe_Hand-in_5/biddybidderpb"
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"google.golang.org/grpc"
)

func main() {
	var lis net.Listener
	var err error
	address := 2
	var addrString string

	// server stuff
	for {
		addrString = "127.0.0." + strconv.Itoa(address) + ":50051"
		lis, err = net.Listen("tcp", addrString)
		if err != nil {
			address++
			continue
		}
		break
	}

	grpcServer := grpc.NewServer()

	s := &server{
		address: address,
	}
	pb.RegisterConsensusServer(grpcServer, s)

	fmt.Println("Server is running on address", addrString)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	client := pb.NewBiddyBidderClient(grpcServer)
	scanner := bufio.NewScanner(os.Stdin)

	// Handle sending messages
	for scanner.Scan() {
		text := scanner.Text()

		if _, err := strconv.Atoi(text); err == nil {
			client.Bid(strconv.Atoi(text))
		}

		if text == "quit" {
			break
		} else {
			fmt.Printf("Doesnt look like a number.\n")
		}

	}

	// Leave the chat
	leaveResp, err := client.LeaveChat(context.Background(), &pb.Participant{Name: name, Timestamp: lamportTime})
	if err != nil {
		log.Fatalf("Could not leave: %v", err)
	}
	clientTick(leaveResp.Timestamp)
	log.Println(leaveResp.Message)

	//

}
