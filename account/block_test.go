package account

import (
	"fmt"
	"testing"
)

func TestGetNowBlock(t *testing.T) {
	block,err := GetNowBlock()
	fmt.Println("err:", err)
	fmt.Println("block number:", block.BlockHeader.RawData.Number)
	fmt.Println("block txns:", len(block.Transactions))
}

func TestGetBlock(t *testing.T) {
	// block number: 24857101
	//block txns: 53
	block, err := GetBlock(24857101)
	fmt.Println("err:", err)
	fmt.Println("block number:", block.BlockHeader.RawData.Number)
	fmt.Println("block txns:", len(block.Transactions))
}