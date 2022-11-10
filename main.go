package main

import (
	request "Lucasfth/go-ass4/grpc/grpc"
	"context"
)

func main() {

}

type peer struct {
	request.UnimplementedRequestServiceServer
	id              int32
	amountOfRequest int32
	peers           map[int32]request.RequestServiceClient
	ctx             context.Context
}

func (p *peer) request(ctx context.Context, req *request.Request) (*request.Request, error) {
	id := req.Id
	amount := req.RequestAmount
}
