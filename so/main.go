package main

import (
        "encoding/json"
        "github.com/gorilla/mux"
        "log"
        "net/http"
        "TLSSign"
)

type Username struct {
        User string `json:"username"`
}

type Signature struct {
        Sig string `json:"sig"`
}

var pri_key = `-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgC/5UQVqe7raMEk53
G/dYGwsESLFMsFJLuIREl4aNpoShRANCAAQYa4PeKwKdltzuMspdK/cBxyyYFTdm
/zKG99+5R/hiEeTev2/ceXqIZauR/40UCB4/zv2dYHnr9YBK0Z8CFadL
-----END PRIVATE KEY-----`
var your_appID = 1400132287

func PostTLSSign(w http.ResponseWriter, req *http.Request) {

        decoder := json.NewDecoder(req.Body)
        var t Username
        err := decoder.Decode(&t)
        if err != nil {
                panic(err)
        }
        defer req.Body.Close()

        tls_conf :=&TLSSign.TLSSignConf{Identifier:t.User, SDKAppId:your_appID, PriKey:pri_key}

        sig, _ := tls_conf.Sign()
        
        sigjson := Signature{sig}
        js, err := json.Marshal(sigjson)
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        w.Write(js)
}


func main() {
        router := mux.NewRouter()

        router.HandleFunc("/tlssign", PostTLSSign).Methods("POST")
        log.Fatal(http.ListenAndServe(":3000", router))
}