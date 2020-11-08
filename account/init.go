package account

import (
	"context"
	"github.com/jack-koli/tron-protocol/core"
	"github.com/mr-tron/base58"
	log "github.com/sirupsen/logrus"
	"github.com/jack-koli/tron-protocol/api"
	"google.golang.org/grpc"
)

var (
	walletClient api.WalletClient
	networkClient api.NetworkClient
)

func init()  {
	log.Info("account.Init")
	conn, err := grpc.Dial("grpc.shasta.trongrid.io:50051", grpc.WithInsecure())
	if err != nil {
		return
	}

	walletClient = api.NewWalletClient(conn)
	networkClient = api.NewNetworkClient(conn)

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



	_, err = walletClient.GetAccount(context.Background(), acc)
	if err != nil {
		log.Errorf("can not get addr info :%s", addr)
		return
	}

	log.Infof("account : %s, balance :%d , ", base58.Encode(acc.Address), acc.GetBalance() )
	// https://shasta.tronscan.org/#/address/TVXTGkZHdrvDQ4TnbLcMCX6TFGsZD1FHn1

	return nil
}
