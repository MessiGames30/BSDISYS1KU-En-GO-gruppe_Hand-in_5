package main

import (
	pb "BSDISYS1KU-En-GO-gruppe_Hand-in_5/biddybidderpb"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	pb.UnimplementedAuctionServer
	address       int
	targetAddress int
	started       bool
	client        pb.AuctionClient
}

type Auction struct {
	AuctionId   int       `json:"name"`
	CurrentBid  int       `json:"type"`
	TimeCreated time.Time `json:"Age"`
}

type Auctions struct {
	PrimaryAuction   Auction `json:"PrimaryAuction"`
	SecondaryAuction Auction `json:"SecondaryAuction"`
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
	pb.RegisterAuctionServer(grpcServer, s)

	fmt.Println("Server is running on address", addrString)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

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
	// Read the auction file
	// Check if the bid is higher than the current bid
	// If it is write the new bid to the file
	// If not return an error

	file, err := os.Open("./Auctions/" + strconv.Itoa(int(bid.AuctionId)) + ".json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	jsonData, err := ioutil.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}
	var data Auctions
	err = json.Unmarshal(jsonData, &data)

	if int(bid.Amount) > data.PrimaryAuction.CurrentBid {
		data.PrimaryAuction.CurrentBid = int(bid.Amount)
		jsonData, err = json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}

		_, err = file.Write(jsonData)
		if err != nil {
			log.Fatal(err)
		}

		return &pb.Ack{
			Status: true,
		}, nil
	}

	return &pb.Ack{
		Status: false,
	}, nil
}

func (s *server) OngoingAuctions(ctx context.Context, empty *pb.Empty) (*pb.Auctions, error) {
	// Scan ./Auctions for files
	// Read the files and return the data
	// If no files are found return an error
	files, err := ioutil.ReadDir("./Auctions")
	if err != nil {
		log.Fatal(err)
	}
	auctions := pb.Auctions{}

	for i := 0; i < len(files); i++ {
		file, err := os.Open("./Auctions/" + strconv.Itoa(i) + ".json")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		jsonData, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}

		auction := GetPrimaryAuctionFromJson(jsonData)

		timeLeft := time.Now().Sub(auction.TimeCreated)

		protoAuction := pb.AuctionDetails{
			AuctionID:  int64(auction.AuctionId),
			Timeleft:   timeLeft.Milliseconds(),
			CurrentBid: int64(auction.CurrentBid),
		}

		auctions.AD = append(auctions.AD, &protoAuction)

		if err != nil {
			log.Fatal(err)
		}

	}

	return &auctions, nil
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

	Auctions := Auctions{
		PrimaryAuction:   data,
		SecondaryAuction: Auction{},
	}

	jsonData, err := json.Marshal(Auctions)
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

	jsonData, err := ioutil.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}
	var data Auctions
	err = json.Unmarshal(jsonData, &data)
	data.SecondaryAuction = auction

	jsonBackedUpData, err := json.Marshal(data)

	_, err = file.Write(jsonBackedUpData)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func GetPrimaryAuctionFromJson(data []byte) Auction {
	var auctions Auctions

	err := json.Unmarshal(data, &auctions)

	if err != nil {
		log.Fatal(err)
	}

	return auctions.PrimaryAuction
}
