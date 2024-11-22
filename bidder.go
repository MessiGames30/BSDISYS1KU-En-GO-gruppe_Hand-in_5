package main

import (
	pb "BSDISYS1KU-En-GO-gruppe_Hand-in_5/biddybidderpb"
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	address := 2
	var addrString string
	bidderId := rand.Intn(10000)
	var client pb.AuctionClient

	// server stuff
	for {
		addrString = "127.0.0." + strconv.Itoa(address) + ":50051"
		conn, err := grpc.NewClient(addrString, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Failed to connect: %v", err)
			continue
		}
		defer conn.Close()
		client = pb.NewAuctionClient(conn)
		break
	}

	fmt.Println("started as bidder with id: " + strconv.Itoa(bidderId))
	scanner := bufio.NewScanner(os.Stdin)

	// Handle sending messages
	for scanner.Scan() {
		text := scanner.Text()
		if newBid, err := strconv.Atoi(text); err == nil {
			result, _ := client.Bid(context.Background(), &pb.Bid{
				BidderId: int64(bidderId),
				Amount:   int64(newBid),
			})
			fmt.Println(result.Status)
			continue
		}

		if text == "status" {
			auction, _ := client.Result(context.Background(), &pb.Empty{})
			if auction.timeLeft <= 0 {
				fmt.Println("Auction is over, the winning bid")
			}
			fmt.Println(auction.HighestBidder, "is winning with bid", auction.CurrentBid)
			fmt.Println("there is ")
		} else if text == "quit" {
			break
		} else {
			fmt.Printf("Doesnt look like a number.\n")
		}

	}

}
