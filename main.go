package main

import (
	request "Lucasfth/go-ass4/grpc/grpc"
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
	arg1, _ := strconv.ParseInt(os.Args[1], 10, 32)
	ownPort := int32(arg1) + 5000

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := &peer{
		id:            ownPort,
		requestAmount: 0,
		peers:         make(map[int32]request.RequestServiceClient),
		response:      make(map[int32]int32),
		ctx:           ctx,
	}

	// Create listener tcp on port ownPort
	list, err := net.Listen("tcp", fmt.Sprintf(":%v", ownPort))
	if err != nil {
		log.Fatalf("Failed to listen on port: %v", err)
	}
	grpcServer := grpc.NewServer()
	request.RegisterRequestServiceServer(grpcServer, p)

	go func() {
		if err := grpcServer.Serve(list); err != nil {
			log.Fatalf("failed to server %v", err)
		}
	}()

	for i := 0; i < 3; i++ {
		port := int32(5000) + int32(i)

		if port == ownPort {
			continue
		}

		var conn *grpc.ClientConn
		log.Printf("Trying to dial: %v\n", port)
		conn, err := grpc.Dial(fmt.Sprintf(":%v", port), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("Could not connect: %s", err)
		}
		defer conn.Close()
		c := request.NewRequestServiceClient(conn)
		p.peers[port] = c
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		p.sendRequestToAll()
		numberOfPeers := len(p.response)

		for i := 0; i < numberOfPeers; i++ {
			if p.response[int32(i)] == 0 {
				break
			}
			if numberOfPeers == i {
				if p.decideControl() {
					p.criticalSection()
				}
				p.clearResonse()
			}
		}
	}
}

type peer struct {
	request.UnimplementedRequestServiceServer
	id            int32
	requestAmount int32
	peers         map[int32]request.RequestServiceClient
	response      map[int32]int32
	ctx           context.Context
}

func (p *peer) Request(ctx context.Context, req *request.Request) (*request.Reply, error) {
	p.requestAmount++

	rep := &request.Reply{Id: p.id, RequestAmount: p.requestAmount}
	return rep, nil
}

func (p *peer) sendRequestToAll() {
	request := &request.Request{Id: p.id, RequestAmount: p.requestAmount}
	for _, peer := range p.peers {
		reply, err := peer.Request(p.ctx, request)
		if err != nil {
			log.Fatalf("Could not send request: %v", err)
		}
		log.Printf("Got reply from id %v: %v\n", p.id, reply.RequestAmount)
		p.response[reply.Id] = reply.RequestAmount
	}
}

func (p *peer) criticalSection() {
	log.Printf("%v is now pilot", p.id)
}

func (p *peer) decideControl() bool {
	numberOfPeers := len(p.response)
	for i := 0; i < numberOfPeers; i++ {
		if p.response[int32(i)] > p.requestAmount {
			return false
		} else if p.response[int32(i+1)] == p.requestAmount {
			if p.id < int32(i+1) {
				return false
			}
		}
	}
	return true
}

func (p *peer) clearResonse() {
	numberOfPeers := len(p.response)
	for i := 0; i < numberOfPeers; i++ {
		p.response[int32(i)] = 0
	}
}
