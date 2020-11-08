package account

import (
	"testing"
)

func TestGetAccount(t *testing.T)  {


	addrs := []string{"TFysCB929XGezbnyumoFScyevjDggu3BPq","TZ5dPxnxd4rRZb7nudcorifD9zfxi2NSRY", "TVXTGkZHdrvDQ4TnbLcMCX6TFGsZD1FHn1"}
	for _, addr := range addrs {
		Info(addr)
	}

}
