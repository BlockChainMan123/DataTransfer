/*
@Description:
*/

package sdk_service

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/gopackager"
	//dzh:
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/cauthdsl"
)

const ChaincodeVersion = "1.0"

//初始化SDK
func SetupSDK(configFile string, initialized bool) (*fabsdk.FabricSDK, error) {

	if initialized {
		return nil, fmt.Errorf("Fabric SDK已被实例化")
	}

	sdk, err := fabsdk.New(config.FromFile(configFile))

	if err != nil {
		return nil, fmt.Errorf("实例化SDK失败")
	}

	fmt.Println("Fabric SDK初始化成功")
	return sdk, nil
}

//创建通道并将指定的peers加入
func CreatChannel(sdk *fabsdk.FabricSDK, inf *InitInfo) error {
	clientContext := sdk.Context(fabsdk.WithUser(inf.OrgAdmin), fabsdk.WithOrg(inf.OrgName))
	if clientContext == nil {
		return fmt.Errorf("根据指定的组织名称与管理员创建资源管理客户端Context失败")
	}

	// New returns a resource management client instance.
	resMgmtClient, err := resmgmt.New(clientContext)
	if err != nil {
		return fmt.Errorf("根据指定的资源管理客户端Context创建通道管理客户端失败: %v", err)
	}

	// New creates a new Client instance
	mspClient, err := mspclient.New(sdk.Context(), mspclient.WithOrg(inf.OrgName))
	if err != nil {
		return fmt.Errorf("根据指定的 OrgName 创建 Org MSP 客户端实例失败: %v", err)
	}

	//  Returns: signing identity
	adminIdentity, err := mspClient.GetSigningIdentity(inf.OrgAdmin)
	if err != nil {
		return fmt.Errorf("获取指定id的签名标识失败: %v", err)
	}

	// SaveChannelRequest holds parameters for save channel request
	channelReq := resmgmt.SaveChannelRequest{ChannelID: inf.ChannelID, ChannelConfigPath: inf.ChannelConfig, SigningIdentities: []msp.SigningIdentity{adminIdentity}}
	// save channel response with transaction ID
	_, err = resMgmtClient.SaveChannel(channelReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint(inf.OrdererOrgName))
	if err != nil {
		return fmt.Errorf("创建应用通道失败: %v", err)
	}

	fmt.Println("通道已成功创建,")

	inf.OrgResMgmt = resMgmtClient

	// allows for peers to join existing channel with optional custom options (specific peers, filtered peers). If peer(s) are not specified in options it will default to all peers that belong to client's MSP.
	err = inf.OrgResMgmt.JoinChannel(inf.ChannelID, resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint(inf.OrdererOrgName))
	if err != nil {
		return fmt.Errorf("Peers加入通道失败: %v", err)
	}

	fmt.Println("peers 已成功加入通道.")
	return nil
}

func NewChannelContext(inf *InitInfo, sdk *fabsdk.FabricSDK) (*channel.Client, error) {
	clientChannelContext := sdk.ChannelContext(inf.ChannelID, fabsdk.WithUser(inf.UserName), fabsdk.WithOrg(inf.OrgName))
	// returns a Client instance. Channel client can query chaincode, execute chaincode and register/unregister for chaincode events on specific channel.
	channelClient, err := channel.New(clientChannelContext)
	if err != nil {
		return nil, fmt.Errorf("创建应用通道客户端失败: %v", err)
	}

	fmt.Println("通道客户端创建成功，可以利用此客户端调用链码进行查询或执行事务.")

	return channelClient, nil
}
//lm: remove InstallAndInstantiateCC & add InstallAndInstantiateCC, InstallAndInstantiateCCWithOne
/*
func InstallAndInstantiateCC(sdk *fabsdk.FabricSDK, inf *InitInfo) (*channel.Client, error) {
	fmt.Println("开始安装链码......")
	// creates new go lang chaincode package
	ccPkg, err := gopackager.NewCCPackage(inf.ChaincodePath, inf.ChaincodeGoPath)
	if err != nil {
		return nil, fmt.Errorf("创建链码包失败: %v", err)
	}

	// contains install chaincode request parameters
	installCCReq := resmgmt.InstallCCRequest{Name: inf.ChaincodeID, Path: inf.ChaincodePath, Version: ChaincodeVersion, Package: ccPkg}
	// allows administrators to install chaincode onto the filesystem of a peer
	_, err = inf.OrgResMgmt.InstallCC(installCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		return nil, fmt.Errorf("安装链码失败: %v", err)
	}
	fmt.Println("指定的链码安装成功")

	fmt.Println("开始实例化链码......")
	//  returns a policy that requires one valid
	ccPolicy := cauthdsl.SignedByAnyMember([]string{"org1.asset.com"})

	instantiateCCReq := resmgmt.InstantiateCCRequest{Name: inf.ChaincodeID, Path: inf.ChaincodePath, Version: ChaincodeVersion, Args: [][]byte{[]byte("init")}, Policy: ccPolicy}
	// instantiates chaincode with optional custom options (specific peers, filtered peers, timeout). If peer(s) are not specified
	_, err = inf.OrgResMgmt.InstantiateCC(inf.ChannelID, instantiateCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		return nil, fmt.Errorf("实例化链码失败: %v", err)
	}

	fmt.Println("链码实例化成功")
	channelContext, _ := NewChannelContext(inf, sdk)
	return channelContext, nil
}*/

func InstallAndInstantiateCC(sdk *fabsdk.FabricSDK, inf *InitInfo) (*channel.Client, error) {
	fmt.Println("安装第一条链码:")
	err := InstallAndInstantiateCCWithOne(sdk, inf, inf.ChaincodeID, inf.ChaincodePath)
	if err != nil {
		return nil, fmt.Errorf("第一条链码安装和实例化失败: %v", err)
	}

	fmt.Println("安装第一条链码成功, 开始安装第二条链码:")
	err = InstallAndInstantiateCCWithOne(sdk, inf, inf.ChaincodeID2, inf.ChaincodePath2)
	if err != nil {
		return nil, fmt.Errorf("第二条链码安装和实例化失败: %v", err)
	}


	channelContext, _ := NewChannelContext(inf, sdk)
	return channelContext, nil
}

func InstallAndInstantiateCCWithOne(sdk *fabsdk.FabricSDK, inf *InitInfo, ChaincodeID string, ChaincodePath string) error {
	fmt.Println("开始安装链码......")
	// creates new go lang chaincode package
	ccPkg, err := gopackager.NewCCPackage(ChaincodePath, inf.ChaincodeGoPath)
	if err != nil {
		return fmt.Errorf("创建链码包失败: %v", err)
	}

	// contains install chaincode request parameters
	installCCReq := resmgmt.InstallCCRequest{Name: ChaincodeID, Path: ChaincodePath, Version: ChaincodeVersion, Package: ccPkg}
	// allows administrators to install chaincode onto the filesystem of a peer
	_, err = inf.OrgResMgmt.InstallCC(installCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		return fmt.Errorf("安装链码失败: %v", err)
	}
	fmt.Println("指定的链码安装成功")

	fmt.Println("开始实例化链码......")
	//  returns a policy that requires one valid
	ccPolicy := cauthdsl.SignedByAnyMember([]string{"org1.asset.com"})
	//ccPolicy := cauthdsl.SignedByAnyMember([]string{os.Getenv("PEER_ORG_NAME")})

	instantiateCCReq := resmgmt.InstantiateCCRequest{Name: ChaincodeID, Path: ChaincodePath, Version: ChaincodeVersion, Args: [][]byte{[]byte("init")}, Policy: ccPolicy}
	// instantiates chaincode with optional custom options (specific peers, filtered peers, timeout). If peer(s) are not specified
	_, err = inf.OrgResMgmt.InstantiateCC(inf.ChannelID, instantiateCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		return fmt.Errorf("实例化链码失败: %v", err)
	}

	fmt.Println("链码实例化成功")
	return nil
}
