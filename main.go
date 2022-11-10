package main

import (
	request "Lucasfth/go-ass4/grpc/grpc"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
)
func main() {
	log.SetFlags(log.Lmicroseconds)
	arg1, _ := strconv.ParseInt(os.Args[1], 10, 32)
	ownPort := int32(arg1) + 5000

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := &peer{
		id:            ownPort,
		requestAmount: 0,
		isPiloting:    false,
		peers:         make(map[int32]request.RequestServiceClient),
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
		port := int32(5000) + int32(i + 1)

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

	//scanner := bufio.NewScanner(os.Stdin)
	for  {
		responses, shouldTry := p.sendRequestToAll()

		if !shouldTry{
			continue
		}

		for i := 0; i < 4; i++ {
			if  i == 3 {
				p.criticalSection()
				p.requestAmount = 0
			}

			if (int32(i) + 5001 == p.id) {
				continue
			}
			if (responses[int32(i) + 5001] > p.requestAmount) {
				break
			} else if (responses[int32(i) + 5001] == p.requestAmount && int32(i) + 5001 > p.id) {
				break
			}
		}
	}
}

type peer struct {
	request.UnimplementedRequestServiceServer
	id            int32
	requestAmount int32
	isPiloting	  bool
	peers         map[int32]request.RequestServiceClient
	ctx           context.Context
}

func (p *peer) Request(ctx context.Context, req *request.Request) (*request.Reply, error) {
	id := req.Id
	reqAmount := p.requestAmount
	
	rep := &request.Reply{Id: id, RequestAmount: reqAmount, IsPiloting: p.isPiloting}
	return rep, nil
}

func (p *peer) sendRequestToAll() (map[int32]int32 , bool){
	response := make(map[int32]int32)
	p.requestAmount++
	request := &request.Request{Id: p.id, RequestAmount: p.requestAmount}
	for id, peer := range p.peers {
		reply, err := peer.Request(p.ctx, request)
		if err != nil {
			log.Fatalf("Could not send request: %v", err)
		}
		log.Printf("Got reply from id %v: %v: %v\n", id, reply.RequestAmount, reply.IsPiloting)
		if reply.IsPiloting {
			time.Sleep(2 * time.Second)
			return make(map[int32]int32) , false
		}
		response[reply.Id] = reply.RequestAmount
	}
	return response , true
	
}

func (p *peer) criticalSection() {
	p.isPiloting = true
	log.Printf("%v is now pilot 	-----------------------", p.id)
	time.Sleep(4 * time.Second)
	log.Printf("%v it not pilot 	-----------", p.id)
	p.isPiloting = false
	time.Sleep(2 * time.Second)
}
