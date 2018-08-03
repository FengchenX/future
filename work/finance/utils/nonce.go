package utils

/*
import (
	"context"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/golang/glog"

	"sub_account_service/blockchain/lib"
)

// Gobal var of All the nonce in api server
var MyNonce *Nonce

type Nonce struct {
	m    map[string]uint64
	Lock sync.RWMutex
}

func NewNonceMap() {
	MyNonce = &Nonce{
		m: make(map[string]uint64),
	}
}

// Change the value of the local nonce map,so we can write
// to the block chain node concurrency.
func (this *Nonce) Calc(auth *bind.TransactOpts, client *ethclient.Client) int64 {
	this.Lock.Lock()
	defer this.Lock.Unlock()

	user_address := lib.Byte20ToStr(auth.From)

	// If nonce of corresponding user_address is not exist
	// get it from the block_chain by rpc func. And reflash
	// the MyNonce map.
	if _, ok := this.m[user_address]; !ok {
		nonce, err := client.PendingNonceAt(context.TODO(), auth.From)
		if err != nil {
			glog.Errorln("Nonce Calc error :", err)
			return -1
		}
		this.m[user_address] = nonce
		auth.Nonce = big.NewInt(int64(nonce))
		this.m[user_address]++
		return int64(nonce)
	} else {
		// Else just add is just fine.
		nonce := this.m[user_address]
		this.m[user_address]++
		auth.Nonce = big.NewInt(int64(nonce))
		return int64(nonce)
	}
}
*/
