package main

import (
	"context"
	"github.com/jack-koli/tron-protocol/api"
	"google.golang.org/grpc"
	"log"
)

func main()  {
	//grpc.shasta.trongrid.io:50051
	addr := "grpc.shasta.trongrid.io:50051"
	conn, err := grpc.Dial( addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect to %s, err:%v" , addr, err)
	}
	client := api.NewWalletClient(conn)

	witnessList, err := client.ListWitnesses(context.Background(), new(api.EmptyMessage))
	if err != nil {
		log.Fatalf("can not get witness err :%v", err)
	}

	for idx, item := range witnessList.GetWitnesses() {
		log.Printf("witness %d is %x\n", idx, item.Address)
	}
}
