package account

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jack-koli/tron-protocol/api"
	"github.com/jack-koli/tron-protocol/common/base58"
	"github.com/jack-koli/tron-protocol/common/crypto"
	"github.com/jack-koli/tron-protocol/common/hexutil"
	"github.com/jack-koli/tron-protocol/core/contract"
	"math/big"
)

// 这个结构目前没有用到 只是记录Trc20合约调用对应转换结果
var mapFunctionTcc20 = map[string]string{
	"a9059cbb": "transfer(address,uint256)",
	"70a08231": "balanceOf(address)",
}

// 转账合约燃烧 trx数量 单位 sun 默认0.5trx 转账一笔大概消耗能量 0.26trx
// 目前 trx 价格 0.026 美元 = 0.17720856 , 每次 0.046 人民币
var feelimit int64 = 500000

func TransferTrx(ownerKey *ecdsa.PrivateKey, toAddr string, amount int64) (result *api.Return, err error) {
	transferContract := new (contract.TransferContract)
	transferContract.OwnerAddress = crypto.PubkeyToAddress(ownerKey.PublicKey).Bytes()
	transferContract.ToAddress, _ = base58.DecodeCheck(toAddr)
	transferContract.Amount = amount
	transferEx, err := walletClient.CreateTransaction2(context.Background(), transferContract)
	if err != nil {
		fmt.Println("fail to create transaction")
		return
	}

	if !transferEx.GetResult().Result {
		err = fmt.Errorf("can not create transaction:1 %s", transferEx.String())
		return
	}

	transaction := transferEx.GetTransaction()
	if transaction == nil || len(transaction.GetRawData().Contract) == 0 {
		err = fmt.Errorf("can not create tranaction:2")
		return
	}
	hash, err := SignTransaction(transaction, ownerKey)
	if err != nil {
		fmt.Printf("fail to sign transaction : %s\n", err.Error())
		return
	}
	fmt.Printf("txID: %s\n", hexutil.Encode(hash))
	result, err = walletClient.BroadcastTransaction(context.Background(),  transaction)
	return
}

func TransferTrc20(ownerKey *ecdsa.PrivateKey, contractAddr ,toAddr string, amount int64) (txID string, err error) {
	transferContract := new(contract.TriggerSmartContract)
	transferContract.OwnerAddress = crypto.PubkeyToAddress(ownerKey.PublicKey).Bytes()
	transferContract.ContractAddress, _ = base58.DecodeCheck(contractAddr)
	transferContract.Data = trc20tranferData(toAddr, amount)
	transferEx, err := walletClient.TriggerConstantContract(context.Background(), transferContract)
	if err != nil {
		fmt.Printf("can not create transferTrc20Ex err:%v\n", err)
		return
	}

	transferTransaction := transferEx.GetTransaction()
	if transferTransaction == nil || len(transferTransaction.GetRawData().GetContract()) == 0 {
		err = fmt.Errorf("can not create transfer trc20 transaction : %s", transferEx.String())
		return
	}

	transferTransaction.RawData.FeeLimit = feelimit

	hash, err := SignTransaction(transferTransaction, ownerKey)
	txID = hexutil.Encode(hash)
	result, err := walletClient.BroadcastTransaction(context.Background(), transferTransaction)
	if err != nil {
		fmt.Printf("transfertrc20 broadcast err:%v\n", err)
		return
	}
	if !result.Result {
		err = fmt.Errorf("api get false, msg: %s", result.String())
	}

	return
}

func trc20tranferData(to string, amount int64) (data []byte) {
	methodID, _ := hexutil.Decode("a9059cbb")
	addr, _ := base58.DecodeCheck(to)
	paddedAddress := common.LeftPadBytes(addr[1:], 32)
	amountBig := new(big.Int).SetInt64(amount)
	paddedAmount := common.LeftPadBytes(amountBig.Bytes(), 32)
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)
	return
}
