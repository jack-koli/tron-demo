package account

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"testing"
)

func TestSendTrx(t *testing.T) {
//key:  5edef98cf81df796028ec09b56ad484319de75f5dab3b87fe81d06adfc603755
//addr:  TUCfXm7opcwaQvXAUvLVhSLFqiUnXpc4K2
	ownerKeyStr := "5edef98cf81df796028ec09b56ad484319de75f5dab3b87fe81d06adfc603755"
	privKey, err := crypto.HexToECDSA(ownerKeyStr)
	if err != nil {
		t.Error("can not init privKey")
	}
	toAddr := "TZ9zDpP2ZsmfVYYTeJgWPKtHzKuzmjTUfU"
	result, err := TransferTrx(privKey, toAddr, 500)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("result: %t \n", result.Result)
	fmt.Println("result string:", result.String())
}

func TestTransferTrc20(t *testing.T) {
	//key:  5edef98cf81df796028ec09b56ad484319de75f5dab3b87fe81d06adfc603755
	//addr:  TUCfXm7opcwaQvXAUvLVhSLFqiUnXpc4K2
	ownerKeyStr := "5edef98cf81df796028ec09b56ad484319de75f5dab3b87fe81d06adfc603755"
	privKey, err := crypto.HexToECDSA(ownerKeyStr)
	if err != nil {
		t.Error("can not init privKey")
	}
	contractAddr := "TF17BgPaZYbz8oxbjhriubPDsA7ArKoLX3"
	toAddr := "TZ9zDpP2ZsmfVYYTeJgWPKtHzKuzmjTUfU"
	txID, err := TransferTrc20(privKey, contractAddr, toAddr, 120000000000000000)

	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("txID", txID)
}