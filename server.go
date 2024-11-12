package main

import (
	pb "Chitty-Chat_HW3_V2/chittychatpb" // Import the generated protobuf package
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"sync"
)

type server struct {
	pb.UnimplementedChittyChatServer
	participants map[string]int64
	messages     []*pb.BroadcastMessage
	clients      map[string]chan *pb.BroadcastMessage // Channel for each client to send messages
	mu           sync.Mutex
	lamportTime  int64
}

// Start the server
func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	s := &server{
		participants: make(map[string]int64),
		messages:     []*pb.BroadcastMessage{},
		clients:      make(map[string]chan *pb.BroadcastMessage), // Initialize the map for client channels
	}
	pb.RegisterChittyChatServer(grpcServer, s)

	fmt.Println("Server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// PublishMessage - allows clients to send a message
func (s *server) PublishMessage(ctx context.Context, msg *pb.ChatMessage) (*pb.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Increment the logical clock for each new message
	serverTick(msg.Timestamp, s)

	// Create a broadcast message with the logical timestamp
	broadcast := &pb.BroadcastMessage{
		Participant: msg.Participant,
		Message:     msg.Message,
		Timestamp:   s.lamportTime,
	}
	s.messages = append(s.messages, broadcast)

	// Log the message
	log.Printf("Message from %s: %s", msg.Participant, msg.Message)

	// Broadcast the message to all connected clients
	for name, ch := range s.clients {
		serverTick(msg.Timestamp, s)
		broadcast := &pb.BroadcastMessage{
			Participant: msg.Participant,
			Message:     msg.Message,
			Timestamp:   s.lamportTime,
		}
		log.Printf("Sending message to client %s", name)
		ch <- broadcast
	}

	return &pb.Empty{}, nil
}

// BroadcastMessages - stream messages to connected clients
func (s *server) BroadcastMessages(_ *pb.Empty, stream pb.ChittyChat_BroadcastMessagesServer) error {
	clientID := fmt.Sprintf("client-%d", len(s.clients)+1) // Unique ID for each client

	// Create a message channel for the client
	msgCh := make(chan *pb.BroadcastMessage, 100)
	s.mu.Lock()
	s.clients[clientID] = msgCh // Store client stream channel
	s.mu.Unlock()

	// Start a goroutine to send messages from the channel to the client
	go func() {
		for msg := range msgCh {
			if err := stream.Send(msg); err != nil {
				if err == io.EOF || err.Error() == "transport is closing" {
					log.Printf("Client %s disconnected", clientID)
					break
				}
				log.Printf("Error sending message to client %s: %v", clientID, err)
			}
		}
	}()

	// Wait for the client to disconnect (block until done)
	<-stream.Context().Done()

	// Remove client from active clients
	s.mu.Lock()
	delete(s.clients, clientID)
	s.mu.Unlock()

	log.Printf("Client %s disconnected", clientID)
	close(msgCh) // Close the message channel
	return nil
}

// JoinChat - client joins and gets a welcome message with a timestamp
func (s *server) JoinChat(ctx context.Context, p *pb.Participant) (*pb.JoinLeaveResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Increment logical clock on join
	serverTick(p.Timestamp, s)

	// Log the join event
	message := fmt.Sprintf("Participant %s joined Chitty-Chat at Lamport time %d", p.Name, s.lamportTime)
	log.Println(message)

	// Broadcast the join event to all clients
	for name, ch := range s.clients {
		// Increasing lamport time
		serverTick(p.Timestamp, s)
		// Create a broadcast message for the join event
		broadcast := &pb.BroadcastMessage{
			Participant: p.Name,
			Message:     "joined the chat",
			Timestamp:   s.lamportTime,
		}
		log.Printf("Notifying client %s about new participant %s", name, p.Name)
		ch <- broadcast
	}

	return &pb.JoinLeaveResponse{
		Message:   message,
		Timestamp: s.lamportTime,
	}, nil
}

// LeaveChat - client leaves and gets a leave message
func (s *server) LeaveChat(ctx context.Context, p *pb.Participant) (*pb.JoinLeaveResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Increment logical clock on leave
	serverTick(p.Timestamp, s)
	delete(s.participants, p.Name)

	// Log the leave event
	message := fmt.Sprintf("Participant %s left Chitty-Chat at Lamport time %d", p.Name, s.lamportTime)
	log.Println(message)

	// Broadcast the leave event to all clients
	for name, ch := range s.clients {
		serverTick(p.Timestamp, s)
		// Create a broadcast message for the leave event
		broadcast := &pb.BroadcastMessage{
			Participant: p.Name,
			Message:     "left the chat",
			Timestamp:   s.lamportTime,
		}
		log.Printf("Notifying client %s about participant %s leaving", name, p.Name)
		ch <- broadcast
	}

	return &pb.JoinLeaveResponse{
		Message:   message,
		Timestamp: s.lamportTime,
	}, nil
}

func serverTick(recivedTime int64, s *server) int64 {
	tempTime := s.lamportTime
	if recivedTime > s.lamportTime {
		s.lamportTime = recivedTime
	}
	s.lamportTime++
	log.Printf("Server lamport time from: %d to %d", tempTime, s.lamportTime)
	return s.lamportTime
}
