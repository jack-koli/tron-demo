package account

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/jack-koli/tron-protocol/api"
	"github.com/jack-koli/tron-protocol/common/base58"
	"github.com/jack-koli/tron-protocol/common/crypto"
	"github.com/jack-koli/tron-protocol/core"
	"github.com/jack-koli/tron-protocol/core/contract"
	log "github.com/sirupsen/logrus"

)

func CreateAccount(ownerKey *ecdsa.PrivateKey, accountAddress string) (result *api.Return, err error) {
	accountContract := new(contract.AccountCreateContract)
	accountContract.OwnerAddress = crypto.PubkeyToAddress(ownerKey.PublicKey).Bytes()
	accountContract.AccountAddress, err = base58.DecodeCheck(accountAddress) // base58 2 byte
	if err != nil {
		log.Errorf("base58.DecodeCheck err :%v", err)
		return
	}
	extension, err := walletClient.CreateAccount2(context.Background(), accountContract)
	if err != nil {
		log.Fatalf("create account error : %v", err)
		return
	}
	log.Println("result:", extension.String())
	accountTransaction := extension.GetTransaction()
	log.Println("accountTransaction:", accountTransaction)

	if accountTransaction == nil {
		log.Fatalf("create account error: invalid transaction : 1")
		return
	}
	rawData := accountTransaction.GetRawData()
	if nil == rawData || len(rawData.Contract) == 0 {
		log.Fatalf("create account error: invalid transaction rawData: %v", rawData)
		return
	}

	SignTransaction(accountTransaction, ownerKey)
	result, err = walletClient.BroadcastTransaction(context.Background(), accountTransaction)
	if err != nil {
		log.Fatalf("create account broadcast error: %v", err)
		return
	}
	fmt.Printf("result: %s\n", result.String())
	return
	//walletClient.CreateAccount(context.Background(),)
}

func UpdateAccount(ownerKey *ecdsa.PrivateKey, accountName string) (rtn *api.Return, err error) {
	accountUpdateContract := new(contract.AccountUpdateContract)
	accountUpdateContract.AccountName = []byte(accountName)
	accountUpdateContract.OwnerAddress = crypto.PubkeyToAddress(ownerKey.PublicKey).Bytes()

	extension, err:= walletClient.UpdateAccount2(context.Background(), accountUpdateContract)

	if err != nil {
		log.Errorf("UpdateAccountTransactionErr:%s", err.Error())
		return nil, err
	}

	accountUpdateTransaction := extension.GetTransaction()
	if nil == accountUpdateTransaction {
		err = fmt.Errorf("can not create update transaction: %s", extension.String())
		return
	}
	_, err = SignTransaction(accountUpdateTransaction, ownerKey)
	if err != nil {
		return
	}

	result, err := walletClient.BroadcastTransaction(context.Background(), accountUpdateTransaction)
	if err != nil {
		log.Errorf("can not broadcastTransaction(accountUpdate) err: %s", err.Error())
	}

	return result, err
}

func GetAccountInfo(addr string) (account *core.Account, err error) {
	acc := new (core.Account)
	acc.Address, err = base58.Decode(addr)
	if err != nil {
		return nil, fmt.Errorf("tron addr(%s) is invalid", addr)
	}

	account, err = walletClient.GetAccount(context.Background(), acc)
	fmt.Println("acc balance:", acc.GetBalance())
	fmt.Println("account", account)
	if err != nil {
		return
	}
	fmt.Printf("balance (%s) : %d\n", addr, account.GetBalance())

	return account, err
}
