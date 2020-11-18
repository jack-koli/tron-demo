package account

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/jack-koli/tron-protocol/api"
	"github.com/jack-koli/tron-protocol/common/base58"
	"github.com/jack-koli/tron-protocol/common/crypto"
	"github.com/jack-koli/tron-protocol/core/contract"
)

func SendTrx(ownerKey *ecdsa.PrivateKey, toAddr string, amount int64) (result *api.Return, err error) {
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
	err = SignTransaction(transaction, ownerKey)
	if err != nil {
		fmt.Printf("fail to sign transaction : %s\n", err.Error())
		return
	}

	result, err = walletClient.BroadcastTransaction(context.Background(),  transaction)
	return
}
