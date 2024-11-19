package main

import (
	pb "BSDISYS1KU-En-GO-gruppe_Hand-in_5/biddybidderpb"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	pb.UnimplementedConsensusServer
	address       int
	targetAddress int
	started       bool
	client        pb.ConsensusClient
}

type Auction struct {
	AuctionId   int       `json:"name"`
	CurrentBid  int       `json:"type"`
	TimeCreated time.Time `json:"Age"`
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
		address: address,
	}
	pb.RegisterBiddyBidderServer(grpcServer, s)

	fmt.Println("Server is running on address", addrString)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func (s *server) PassToken(ctx context.Context, empty *pb.Token) (*pb.Empty, error) {
	log.Println("Token recieved")
	go writeToFile(s)

	return &pb.Empty{}, nil
}

func writeToFile(s *server) {
	if rand.Intn(10) == 0 {
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		data := []byte("Client " + strconv.Itoa(s.address) + " wrote at: " + currentTime + "\n")
		file, err := os.OpenFile("CriticalSection.txt", os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		_, err = file.Write(data)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()

		log.Println("Wrote to file")
		time.Sleep(1 * time.Second)
	}
	s.client.PassToken(context.Background(), &pb.Token{})
	log.Println("sent Token")
}

func (s *server) StartFunction(ctx context.Context, empty *pb.Empty) (*pb.SuccessStart, error) {
	if s.started {
		log.Println("Server already started")
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

	if s.address == 2 {
		go writeToFile(s)
	}

	return &pb.SuccessStart{
		Message: "Server " + strconv.Itoa(s.address) + " started",
	}, nil
}

func connectToServer(address int) (*pb.SuccessStart, pb.BiddyBidderClient) {
	connectAddress := "127.0.0." + strconv.Itoa(address) + ":50051"
	conn, err := grpc.NewClient(connectAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	// defer conn.Close()

	client := pb.NewBiddyBidderClient(conn)
	message, err := client.StartFunction(context.Background(), &pb.Empty{})
	return message, client

}

func OngoingAuctions(ey &empty) (&Auction, error) {
	// Scan ./Auctions for files
	// Read the files and return the data
	// If no files are found return an error
	files, err := ioutil.ReadDir("./Auctions")
	if err != nil {
		log.Fatal(err)
	}
	if len(files) == 0 {
		return nil, errors.New("No ongoing auctions")
	}

	

	for _, file := range files {
	
	}

	file, err := os.Open("./Auctions/" + strconv.Itoa(len(files)) + ".json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := Auction{}
	jsonData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Fatal(err)
	}

	return &data, nil
	
}

func createAuctionFile() (int, error) {
	// Scane ./Auctions for files creates
	// If no file is found create a new file with the name 1.json
	// If a file is found create a new file with the name of the last file + 1
	// Write the auction to the file

	files, err := ioutil.ReadDir("./Auctions")
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create("./Auctions/" + strconv.Itoa(len(files)+1) + ".json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := Auction{
		AuctionId:   len(files) + 1,
		CurrentBid:  0,
		TimeCreated: time.Now(),
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatal(err)
	}

	return len(files) + 1, nil // IF PROBLEM WITH AUCTION ID FIX THIS
}

func backupAuction(auction Auction) error {

	id := auction.AuctionId
	fileNumber := 0
	files, err := ioutil.ReadDir("./Auctions")

	if id == len(files) {
		fileNumber = 1
	} else {
		fileNumber = id + 1
	}

	file, err := os.Open("./Auctions/" + strconv.Itoa(fileNumber) + ".json")

	jsonData, err := json.Marshal(auction)
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
