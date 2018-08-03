package api

import (
	// "sub_account_service/app_server_v2/db"
	// "sync"
	// "time"

	// "github.com/golang/glog"
)

// const (
// 	SUBWAY_MONTH = 2
// 	SUBWAY_DATE  = 1
// 	FRISTDAY     = 1
// 	TIMEOUT      = 60 * time.Second
// 	ADAY         = 24 * time.Hour
// 	RESETPENDING = 20 * time.Second
// 	RESETTIMES   = "10"
// )

// //Rfs 不晓得思昊用意
// var Rfs *rfs

// type rfs struct {
// 	m map[int][]*db.Reflash
// 	l sync.RWMutex
// }

// //InitRfs 不晓得作用
// func InitRfs() {
// 	Rfs = &rfs{
// 		m: make(map[int][]*db.Reflash),
// 	}
// }

// //Add 不晓得作用
// func (rs *rfs) Add(i int, r *db.Reflash) {
// 	rs.l.Lock()
// 	defer rs.l.Unlock()

// 	glog.Infoln("add", i, r)
// 	if _, exist := rs.m[i]; exist {
// 		rs.m[i] = append(rs.m[i], r)
// 	} else {
// 		rs.m[i] = []*db.Reflash{r}
// 	}
// 	if dberr := db.DbClient.Client.Create(r); dberr.Error != nil {
// 		glog.Errorln("dberr", dberr.Error)
// 	}
// }

// func (rs *rfs) Get(key int) []*db.Reflash {
// 	rs.l.Lock()
// 	defer rs.l.Unlock()

// 	if val, exist := rs.m[key]; exist {
// 		return val
// 	}
// 	return nil
// }
