package account

import (
	"context"
	"github.com/jack-koli/tron-protocol/api"
	"github.com/jack-koli/tron-protocol/core"
	"github.com/mr-tron/base58"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	conn *grpc.ClientConn
	walletClient api.WalletClient
	databaseClient api.DatabaseClient
)

func init()  {
	log.Info("account.Init")
	//rpcHost 正式 grpc.trongrid.io:50051
	// 测试
	//rpcHost := "grpc.shasta.trongrid.io:50051"
	rpcHost := "192.168.8.240:50051"
	var err error
	conn, err = grpc.Dial(rpcHost, grpc.WithInsecure())
	if err != nil {
		log.Errorf("can not grpc.Dial %v", err)
		return
	}

	walletClient = api.NewWalletClient(conn)

	databaseClient = api.NewDatabaseClient(conn)
}

func ListWitnesses() (*api.WitnessList, error) {
	witnessList, err := walletClient.ListWitnesses(context.Background(), new(api.EmptyMessage))

	return witnessList, err
}



func NodeInfo()  {
	info, err := walletClient.GetNodeInfo(context.Background(), new(api.EmptyMessage))
	if err != nil {
		return
	}

	log.Info("node: sync:", info.GetBeginSyncNum())

}

func Info(addr string) (err error) {

	acc := new(core.Account)
	acc.Address, err = base58.Decode(addr)
	if err != nil {
		return err
	}


	//databaseClient.GetBlockByNum()

	_, err = walletClient.GetAccount(context.Background(), acc)
	if err != nil {
		log.Errorf("can not get addr info :%s", addr)
		return
	}

	log.Infof("account : %s, balance :%d , ", base58.Encode(acc.Address), acc.GetBalance() )
	// https://shasta.tronscan.org/#/address/TVXTGkZHdrvDQ4TnbLcMCX6TFGsZD1FHn1

	return nil
}
