package sdk

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"fmt"
	"sub_account_service/fabric_server/model"
)

// QueryHello query the chaincode to get the state of hello
func (setup *FabricSetup) QueryAccount(req model.ReqSetAccount) (string, error) {
	// Prepare arguments
	var args [][]byte
	args = append(args, []byte("key1"))

	request := channel.Request{
		ChaincodeID: "demo",
		Fcn:         "query",
		Args:        args,
	}
	response, err := setup.client.Query(request)

	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}

	return string(response.Payload), nil
}
