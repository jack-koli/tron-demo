package account

import (
	"fmt"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	key := GenerateKey()

	addr := AddressFromPrv(key)
	fmt.Println("key: ", key)
	fmt.Println("addr: ", addr)
}

