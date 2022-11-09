package main

import (
	"context"
	request "github.com/Lucasfth/peer-to-peer/grpc"
	
)

func main(){

}

type peer struct{
	request.UnimplementedRequestServiceServer
	id int32
	amountOfRequest int32
	
}

func (p *peer) request(ctx context.Context)(){
	
}