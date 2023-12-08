/*
@Description:链码初始化和调用
*/

package main

import (
	"fmt"
	"github.com/Transfer_HyperledgerFabric/chaincode/plaintext_equipment_cc/routers"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type Plain_Equipment_Chaincode struct {
}

func (t *Plain_Equipment_Chaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t *Plain_Equipment_Chaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	funcName, args := stub.GetFunctionAndParameters()
	switch funcName {
	case "ZSGCPlainEquipment":
		return routers.ZSGCPlainEquipment(stub, args)
	case "GetAllPlainEquipmentWithPagination":
		return routers.GetAllPlainEquipmentWithPagination(stub, args)
	case "QueryPlainEquipmentByHash":
		return routers.QueryPlainEquipmentByHash(stub, args)
	case "GetHistoryPlainWithPagination":
		return routers.GetHistoryPlainWithPagination(stub, args)

	default:
		return shim.Error(fmt.Sprintf("没有该功能: %s", funcName))
	}
}

func main() {
	err := shim.Start(new(Plain_Equipment_Chaincode))
	if err != nil {
		fmt.Printf("启动Assets_Equipment_Chaincode时发生错误: %s", err)
	}
}
