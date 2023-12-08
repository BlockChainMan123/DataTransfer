package main

import (
	"fmt"
	"github.com/Transfer_HyperledgerFabric/sdk_service"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"os"
	"github.com/Transfer_HyperledgerFabric/web"
	"github.com/Transfer_HyperledgerFabric/web/controller"



)

const (
	configFile  = "config.yaml"
	initialized = false
	PlainCC     = "plaincc"
	CipherCC = "ciphercc"
)


var initInf = &sdk_service.InitInfo{
	ChannelID:      "Transfer_CrossChain",
	ChannelConfig:  os.Getenv("GOPATH") + "/src/github.com/Transfer_HyperledgerFabric/fixtures/artifacts/channel.tx",
	OrgAdmin:       "Admin",
	OrgName:        "Org1",
	OrdererOrgName: "orderer.asset.com",

	ChaincodeID:     PlainCC,
	ChaincodeGoPath: os.Getenv("GOPATH"),
	ChaincodePath:   "github.com/Transfer_HyperledgerFabric/chaincode/plaintext_equipment_cc",

	ChaincodeID2:   CipherCC,
	ChaincodePath2: "github.com/Transfer_HyperledgerFabric/chaincode/ciphertext_equipment_cc",

	UserName:       "User1",
}

func main() {
	sdk, channelClient := sdkStart()


	serviceUp(sdk, channelClient)
	serviceSetup := sdk_service.ServiceSetup{
		ChaincodeID: PlainCC,
		ChaincodeID2: CipherCC,
		Client:       channelClient,
	}
	app := controller.Application{
		SdkSetup: &serviceSetup,
	}
	web.WebStart(app)

}

func sdkStart() (*fabsdk.FabricSDK, *channel.Client) {
	sdk, err := sdk_service.SetupSDK(configFile, initialized)

	if err != nil {
		fmt.Println(err.Error())
	}
	//defer sdk.Close()

	err = sdk_service.CreatChannel(sdk, initInf)
	if err != nil {
		fmt.Println(err.Error())
		return nil, nil
	}

	channelClient, err := sdk_service.InstallAndInstantiateCC(sdk, initInf)
	if err != nil {
		fmt.Println(err.Error())
		return nil, nil
	}

	//fmt.Println(channelClient)

	return sdk, channelClient
}

func serviceUp(sdk *fabsdk.FabricSDK, channelClient *channel.Client) {
	defer sdk.Close()

	serviceSetup := sdk_service.ServiceSetup{
		ChaincodeID: PlainCC,
		ChaincodeID2: CipherCC,
		Client:       channelClient,
	}

	app := controller.Application{
		SdkSetup: &serviceSetup,
	}
	web.WebStart(app)
}
