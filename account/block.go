package account

import (
	"context"
	"github.com/jack-koli/tron-protocol/api"
	"github.com/jack-koli/tron-protocol/core"
	log "github.com/sirupsen/logrus"
)

func GetNowBlock() (block *core.Block,err error) {
	block,err = walletClient.GetNowBlock(context.Background(), new(api.EmptyMessage))
	if err != nil {
		return
	}
	log.Info("now block number",block.BlockHeader.RawData.Number)
	return
}

func GetBlock(number int64) (block *core.Block, err error) {
	block, err = walletClient.GetBlockByNum(context.Background(), &api.NumberMessage{Num: number})
	return
}