package account

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"testing"
)

func TestAccountInfo(t *testing.T)  {
	// TZ9zDpP2ZsmfVYYTeJgWPKtHzKuzmjTUfU , 不存在的地址
	// TAfqywW7vAT7Cfz6mPKeULQsq16yjzN6iy test usdt
	// TY15WF5kUQrCrioKMr3J3Ko3HvNzPGwW9j test , 有trx 余额，有冻结的
	// TYT9XYA3cjXs2Zz8Lqh812TJXNbHHGHVkz 正式地址 ,测试网络没有
	// TAUN6FwrnwwmaEqYcckffC7wYmbaS6cBiX 正式网络，很多 usdt ,测试网络也有很多trx
	addrs := []string{"TAUN6FwrnwwmaEqYcckffC7wYmbaS6cBiX","TZ9zDpP2ZsmfVYYTeJgWPKtHzKuzmjTUfU" , "TY15WF5kUQrCrioKMr3J3Ko3HvNzPGwW9j", "TAfqywW7vAT7Cfz6mPKeULQsq16yjzN6iy" ,"TYT9XYA3cjXs2Zz8Lqh812TJXNbHHGHVkz"}
	for _, addr := range addrs {
		acc,err := GetAccountInfo(addr)
		fmt.Println("addr:", addr," err:", err)
		fmt.Println("acc:", acc.GetBalance())
	}


}

func TestCreateAccount(t *testing.T)  {
/*key:  4c20e10d5228fc0bdc0d0769f70b565c67b8e65fef90680ce610f6b1b7d39a5c
addr:  TLQ3gCw9oWahTFQtJkwjq4UtUdVFdoggYr
*/
	//key:  c8b80dc4130ed9d0b58428ea171224ca185dc7a44e2e0e5662f54cb0215c0992
//addr:  TPY6G94vmGB6V7La9En719EfmBh64LYpgi

// 有余额 prvA := "416EBFF57188ED9BA1EDF68CD9D724D974667D6BCB3FE6BD6E4AA7B3F35DBD92"
	// addrA TZ9zDpP2ZsmfVYYTeJgWPKtHzKuzmjTUfU
//key:  efd816f904ec382268a917e098006a4383ed905bd465310adc29350bf3b93539
//addr:  TMu8NjSMqPD3Bd6PRuLeVsijPz2f5Egz2U

	//
	prvA := "416EBFF57188ED9BA1EDF68CD9D724D974667D6BCB3FE6BD6E4AA7B3F35DBD92"
	privKey, err := crypto.HexToECDSA(prvA)
	if err != nil {
		t.Error("wrong key")
		return
	}
	addrA := AddressFromPrv(prvA)
	addr := "TPY6G94vmGB6V7La9En719EfmBh64LYpgi"
	fmt.Println("addrA:", addrA)
	resp, err := CreateAccount(privKey, addr)
	if err != nil {
		fmt.Printf("CreateAccount Error: %v\n", err)
		return
	}
	if resp.Result == true {
		fmt.Printf("created %s\n", addr)
	}
	fmt.Println("message ",string(resp.Message))
	//addrA: TZ9zDpP2ZsmfVYYTeJgWPKtHzKuzmjTUfU
	/**
	PrivB : 5D282BD506752AC164EED18A33AD2AC317DA3084ABD2EEF299599BD3785D8FD8
	addrB: TDdAeYQXHYDEioWy2xfe9ZFMTRPCe9gkdk

	privC: A7298BF2AED729330A1F888A82004174DC8218036CB77FE7BC2F119D75F1B5D3
	addRC: TZ5dPxnxd4rRZb7nudcorifD9zfxi2NSRY
	 */
}

func TestUpdateAccount(t *testing.T) {
	//key:  efd816f904ec382268a917e098006a4383ed905bd465310adc29350bf3b93539
	//addr:  TMu8NjSMqPD3Bd6PRuLeVsijPz2f5Egz2U
	prvA := "efd816f904ec382268a917e098006a4383ed905bd465310adc29350bf3b93539"
	privKey, err := crypto.HexToECDSA(prvA)
	if err != nil {
		t.Error("wrong key")
		return
	}
	_,err = UpdateAccount(privKey, "testNice2")
	if err != nil {
		fmt.Printf("UpdateAccount err:%s", err.Error())
	}
}

func TestGetAccountInfo(t *testing.T) {
	addr1 := "TZ9zDpP2ZsmfVYYTeJgWPKtHzKuzmjTUfU"
	addr2 := "TUCfXm7opcwaQvXAUvLVhSLFqiUnXpc4K2"
	acc1, err1 := GetAccountInfo(addr1)
	if err1 != nil {
		fmt.Printf("err1 :%s\n", err1.Error())
	}

	fmt.Printf("acc1 %v", acc1)

	acc2, err2 := GetAccountInfo(addr2)
	if err2 != nil {
		fmt.Printf("err1 :%s\n", err2.Error())
	}

	fmt.Printf("acc1 %v", acc2)
}