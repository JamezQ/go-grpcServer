package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	chatpb "github.com/JamezQ/go-grpcServer/protos/proto"
	pro "github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8086, "The server port")
)

type chatServer struct {
	chatpb.UnimplementedChatServerServer
	mu          sync.Mutex
	chatHistory *chatpb.ChannelHistory
	listeners   map[int]ListenerFunc
	id          int
}

type ListenerFunc func(msg *chatpb.ChannelChat)

func (c *chatServer) GetHistory(ctx context.Context, req *chatpb.Empty) (*chatpb.ChannelHistory, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	historyCopy := pro.Clone(c.chatHistory).(*chatpb.ChannelHistory)
	return historyCopy, nil
}

func (c *chatServer) ListenChannel(ch *chatpb.Channel, stream chatpb.ChatServer_ListenChannelServer) error {

	return nil
}

func (c *chatServer) ConnectChat(stream chatpb.ChatServer_ConnectChatServer) error {
	return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	chatpb.RegisterChatServerServer(grpcServer, &chatServer{})
	log.Println("Server started")
	grpcServer.Serve(lis)
}
