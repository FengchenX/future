package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"sub_account_service/fabric_server/api"
	"sub_account_service/fabric_server/httpServ"
	"sub_account_service/fabric_server/sdk"
)


func main() {

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)

	// Definition of the Fabric SDK properties
	fSetup := sdk.FabricSetup{
		// Network parameters
		OrdererID: "orderer.launch.com",

		// Channel parameters
		ChannelID:     "mychannel",
		ChannelConfig: os.Getenv("GOPATH") + "/src/sub_account_service/fabric_server/artifacts/mychannel.tx",

		// Chaincode parameters
		ChainCodeID: "demo",
		OrgAdmin:    "Admin",
		OrgName:     "org1.launch.com",
		ConfigFile:  os.Getenv("GOPATH") + "/src/sub_account_service/fabric_server/artifacts/config.yaml",
	}

	// Initialization of the Fabric SDK from the previously set properties
	err := fSetup.Initialize()
	if err != nil {
		fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
		//return
	}
	// Close SDK
	defer fSetup.CloseSDK()

	sever := &api.ApiService{
		Fabric: &fSetup,
	}
	httpServ.Serve(sever)
}
