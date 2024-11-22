package main

import (
	pb "BSDISYS1KU-En-GO-gruppe_Hand-in_5/biddybidderpb"
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	pb.UnimplementedAuctionServer
	address        int
	targetAddress  int
	started        bool
	currentAuction Auction
	client         pb.AuctionClient
	currentTime    int
}

type Auction struct {
	CurrentBid  int
	TimeCreated int
	Duration    int
}

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
		address:     address,
		started:     false,
		currentTime: 0,
	}
	pb.RegisterAuctionServer(grpcServer, s)

	fmt.Println("Server is running on address", addrString)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func (s *server) StartFunction(ctx context.Context, time *pb.time) (*pb.SuccessStart, error) {
	if s.started {
		log.Println("Server already started")
		s.currentAuction = Auction{
			CurrentBid:  0,
			TimeCreated: s.currentTime,
			Duration:    100,
		}
		return &pb.SuccessStart{
			Message: "Server did not start",
		}, nil
	}
	s.started = true
	log.Println("Server " + strconv.Itoa(s.address) + " started")

	s.targetAddress = s.address + 1
	message, client := connectToServer(s.targetAddress)
	if message == nil {
		s.targetAddress = 2
		message, client = connectToServer(s.targetAddress)

	}
	s.client = client

	log.Println("Connect Success to address" + strconv.Itoa(s.targetAddress))

	return &pb.SuccessStart{
		Message: "Server " + strconv.Itoa(s.address) + " started",
	}, nil
}

func connectToServer(address int) (*pb.SuccessStart, pb.AuctionClient) {
	connectAddress := "127.0.0." + strconv.Itoa(address) + ":50051"
	conn, err := grpc.NewClient(connectAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	// defer conn.Close()

	client := pb.NewAuctionClient(conn)
	message, err := client.StartFunction(context.Background(), &pb.Empty{})
	return message, client

}

func (s *server) Bid(ctx context.Context, bid *pb.Bid) (*pb.Ack, error) {
	// Read the auction memory
	// Check if the bid is higher than the current bid
	// If it is write the new bid to the memory
	// If not return an error

	return &pb.Ack{
		Status: false,
	}, nil
}

func (s *server) OngoingAuction(ctx context.Context, empty *pb.Empty) (*pb.AuctionDetails, error) {
	// Scan ./Auctions for files
	// Read the files and return the data
	// If no files are found return an error
	auctions := pb.AuctionDetails{}
	return &auctions, nil
}
