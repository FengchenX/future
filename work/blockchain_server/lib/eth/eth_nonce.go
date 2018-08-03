package myeth

import (
	"bytes"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc/naming"
	"math/big"
	"sync"
	"time"
)

const REFLASHTIME = 10 * time.Second
const HTTPTIMEOUT = 10 * time.Second

var NonceMap *Nonce

type n struct {
	num        uint64
	createTime time.Time
}

type Nonce struct {
	m    map[string]*n
	Lock sync.RWMutex
}

func NewNonceMap() {
	NonceMap = &Nonce{
		m: make(map[string]*n),
	}
}

// every 10 second recall the ipc to get the correct nonce.
// but whenever the nodes is just pending when we reseting the
// nonce we can hardly get the correct nonce.
func (this *Nonce) Reflash() {
	naming.NewDNSResolver()
	for {
		time.Sleep(REFLASHTIME)
		this.Lock.Lock()
		for k, v := range this.m {
			if time.Now().Sub(v.createTime) >= 10 {
				glog.Infoln("delete", k)
				delete(NonceMap.m, k)
			}
		}
		this.Lock.Unlock()
	}
}

// Change the value of the local nonce map,so we can write
// to the block chain node concurrency.
func (this *Nonce) calc(addr string, auth *bind.TransactOpts, client *ethclient.Client) int64 {
	this.Lock.Lock()
	defer this.Lock.Unlock()

	addr = ParseUserName(addr)
	if addr == "" {
		glog.Errorln("user_address == nil")
		return -1
	}

	glog.Infoln("Nonce Calc", addr)

	// If nonce of corresponding user_address is not exist
	// get it from the block_chain by rpc func. And reflash
	// the MyNonce map.
	if _, ok := this.m[addr]; !ok {
		ctx, cancel := context.WithTimeout(context.Background(), HTTPTIMEOUT)
		nonce, err := client.PendingNonceAt(ctx, auth.From)
		defer cancel()
		if err != nil {
			glog.Errorln("Nonce Calc error :", err)
			return -1
		}
		this.m[addr] = &n{
			num:        nonce,
			createTime: time.Now(),
		}
		auth.Nonce = big.NewInt(int64(nonce))
		this.m[addr].num++
		return int64(nonce)
	} else {
		// Else just add is just fine.
		nonce := this.m[addr].num
		this.m[addr].num++
		auth.Nonce = big.NewInt(int64(nonce))
		return int64(nonce)
	}
}

// when the contract func returns error ,reset the nonce in case of the
// error is call by the wrong nonce
func (this *Nonce) reset22222(key_string string, client *ethclient.Client) int64 {
	this.Lock.Lock()
	defer this.Lock.Unlock()

	mp := make(map[string]interface{})
	json.Unmarshal(bytes.NewBufferString(key_string).Bytes(), &mp)

	glog.Infoln(mp)

	if user_address, exist := mp["address"]; exist {
		u_addr := user_address.(string)
		u_addr = ParseUserName(u_addr)
		ctx, cancel := context.WithTimeout(context.Background(), HTTPTIMEOUT)
		nonce, err := client.PendingNonceAt(ctx, common.HexToAddress(u_addr))
		if err != nil {
			glog.Errorln("Nonce Calc error :", err)
			return -1
		}
		glog.Infoln("Reseting Nonce", nonce, u_addr)
		this.m[u_addr] = &n{
			num:        nonce,
			createTime: time.Now(),
		}
		cancel()
		return 0
	}
	return -1
}

// if the connection is down reset all nonce of the users
func (this *Nonce) resetAll11111(client *ethclient.Client) int64 {

	for addr, _ := range this.m {
		ctx, cancel := context.WithTimeout(context.Background(), HTTPTIMEOUT)
		nonce, err := client.PendingNonceAt(ctx, common.HexToAddress(addr))
		if err != nil {
			glog.Errorln("Nonce Calc error :", err)
			return -1
		}
		glog.Infoln("Reseting Nonce", nonce, addr)
		this.Lock.Lock()
		this.m[addr] = &n{
			num:        nonce,
			createTime: time.Now(),
		}
		cancel()
		this.Lock.Unlock()
		return 0
	}
	return -1
}
