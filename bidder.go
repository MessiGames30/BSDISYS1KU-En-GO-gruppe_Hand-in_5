package main

import (
	pb "BSDISYS1KU-En-GO-gruppe_Hand-in_5/biddybidderpb"
	"bufio"
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	bidderId := rand.Intn(10000)

	client := *getConnection()

	fmt.Println("started as bidder with id: " + strconv.Itoa(bidderId))
	scanner := bufio.NewScanner(os.Stdin)

	// Handle sending messages
	for scanner.Scan() {
		text := scanner.Text()
		if newBid, err := strconv.Atoi(text); err == nil {
			for {
				result, err2 := client.Bid(context.Background(), &pb.Bid{
					BidderId: int64(bidderId),
					Amount:   int64(newBid),
				})
				if err2 != nil {
					client = *getConnection()
					continue
				}

				fmt.Println(result.Status)
				break
			}
			if getStatus(client) {
				break
			}
			continue
		}

		if text == "status" {
			if getStatus(client) {
				break
			}
			continue
		}

		if text == "quit" {
			break
		}

		fmt.Printf("Doesnt look like a number.\n")
	}
}

func getStatus(client pb.AuctionClient) bool {
	auction, _ := client.Result(context.Background(), &pb.Empty{})
	if auction.Timeleft < 0 {
		fmt.Println("Auction is over, the winning bid was", auction.HighestBid, "made by", auction.HighestBidder)
		return true
	}
	fmt.Println(auction.HighestBidder, "is winning with bid", auction.HighestBid)
	fmt.Println("there is", auction.Timeleft, "units of time left")
	return false
}

func getConnection() *pb.AuctionClient {
	address := 2
	var addrString string
	var client pb.AuctionClient
	// server stuff
	for {
		addrString = "127.0.0." + strconv.Itoa(address) + ":50051"
		conn, err := grpc.NewClient(addrString, grpc.WithTransportCredentials(insecure.NewCredentials()))
		client = pb.NewAuctionClient(conn)
		_, err = client.Ping(context.Background(), &pb.Empty{})
		if err != nil {
			address++
			continue
		}
		break
	}
	fmt.Println("Connected to server", address)
	return &client
}
