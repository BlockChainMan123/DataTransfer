/*
@Description:链码初始化和调用
*/

package main

import (
	"fmt"

	"github.com/Transfer_HyperledgerFabric/chaincode/ciphertext_equipment_cc/routers"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type Ciphertext_Equipment_Chaincode struct {
}

func (t *Ciphertext_Equipment_Chaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t *Ciphertext_Equipment_Chaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	funcName, args := stub.GetFunctionAndParameters()
	switch funcName {
	case "ZSGCCipherEquipment":
		return routers.ZSGCCipherEquipment(stub, args)
	case "QueryCipherEquipmentByHash":
		return routers.QueryCipherEquipmentByHash(stub, args)
	case "GetHistoryCipherWithPagination":
		return routers.GetHistoryCipherWithPagination(stub, args)

	default:
		return shim.Error(fmt.Sprintf("没有该功能: %s", funcName))
	}
}

func main() {
	err := shim.Start(new(Ciphertext_Equipment_Chaincode))
	if err != nil {
		fmt.Printf("启动Ciphertext_Equipment_Chaincode时发生错误: %s", err)
	}
}
