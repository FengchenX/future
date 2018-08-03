package myeth

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"sync"
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"math/big"
	"fmt"
	"strings"
)

type element struct {
	c chan *big.Int
}

var payAddressNonceMap = sync.Map{}
func getNonce(auth *bind.TransactOpts, client *ethclient.Client) big.Int{
	addr := strings.ToLower(auth.From.String())
	elePointer,ok := payAddressNonceMap.Load(addr)
	if !ok {
		initPayAddressNonce(addr, auth, client)
		elePointer,_ = payAddressNonceMap.Load(addr)
	}
	ele := elePointer.(*element)
	nonce := <-ele.c
	auth.Nonce = nonce
	return  *nonce
}

// write nonce to channel
// if block chain success,write nonce + 1 else write current nonce
func UpdateNonce(author *bind.TransactOpts, client *ethclient.Client, err error, args ...interface{}) {
	addr := strings.ToLower(author.From.String())
	elePointer,ok := payAddressNonceMap.Load(addr)
	glog.Infoln(addr, "begin update nonce:",args, "current nonce:",author.Nonce, "err:",err)
	if ok {
		ele := elePointer.(*element)
		updateNonce := author.Nonce.Int64()
		if err == nil {
			updateNonce = author.Nonce.Int64() + 1
			ele.c <- big.NewInt(updateNonce)
		} else {//failed，write current nonce
			// failed try connect eth to get the latest nonce
			// prevent another process attach eth then let the nonce changed
			ctx, cancel := context.WithTimeout(context.Background(), HTTPTIMEOUT)
			nonce, err := client.PendingNonceAt(ctx, author.From)
			defer cancel()
			if err != nil { // get nonce err,write current nonce
				ele.c <- author.Nonce
				glog.Errorln("nonce get error :", err)
			} else {
				ele.c <- big.NewInt(int64(nonce))
			}
		}
		glog.Infoln(addr, "update nonce end:",args,"update nonce:",updateNonce)
	}
}

// init nonce per pay address
var initPayAddressNonceLock = sync.Mutex{}
func initPayAddressNonce(payAddress string, auth *bind.TransactOpts, client *ethclient.Client){
	initPayAddressNonceLock.Lock()
	fmt.Println("get nonce addr:",payAddress)
	defer func() {
		if err := recover(); err != nil {
		}
		initPayAddressNonceLock.Unlock()
	}()
	if _,ok := payAddressNonceMap.Load(payAddress); !ok { //if not exist
		ctx, cancel := context.WithTimeout(context.Background(), HTTPTIMEOUT)
		nonce, err := client.PendingNonceAt(ctx, auth.From)
		defer cancel()
		if err != nil {
			glog.Errorln("nonce init error :", err)
			return
		}
		ele := &element{
			c: make(chan *big.Int,1),
		}
		glog.Infoln("init nonce success,latest nonce is :",nonce)
		payAddressNonceMap.Store(payAddress, ele)
		ele.c <- big.NewInt(int64(nonce))
	}
}