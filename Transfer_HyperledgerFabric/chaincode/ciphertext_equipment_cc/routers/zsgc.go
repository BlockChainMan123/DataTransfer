package routers

import (
	"encoding/json"
	"fmt"

	"github.com/Transfer_HyperledgerFabric/chaincode/ciphertext_equipment_cc/lib"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"strconv"
	"time"
)

const DOC_TYPE = "record_equipment_object"

func putZSGCCipherEquipment(stub shim.ChaincodeStubInterface, input_asset lib.Input_Cipher_equipment, FunctionPath string, OperatingContract string) ([]byte, bool) {

	var asset lib.Cipher_equipment

	asset.DocType = DOC_TYPE

	total, err := computeTotalForrange(stub, "", "")
	asset.AutoID = strconv.Itoa(5000 - total)
	asset.ContractInformation_OperatingContract = OperatingContract
	asset.ContractInformation_FunctionPath = FunctionPath
	asset.ContractInformation_TimeStamp = lib.BJtime(time.Now().Format("2006-01-02 15:04:05"))
	asset.ContractInformation_OnlyHash = lib.GetMD5String([]byte(string(asset.AutoID)))
	asset.BlockInformation_TimeStamp = lib.BJtimeShort(time.Now().Format("2006-01-02"))
	asset.BlockInformation_ChainPlace = "cipher_chain"
	block := total / 10
	block = block + 1
	asset.BlockInformation_BlockPlace = strconv.Itoa(block)
	asset.BlockInformation_BlockCode = lib.GetMD5String([]byte(string(block)))



	asset.ID = input_asset.ID
	asset.Address = input_asset.Address
	asset.Sk = input_asset.Sk
	asset.Request = input_asset.Request


	b, err := json.Marshal(asset)
	if err != nil {
		return nil, false
	}

	// save information.
	err = stub.PutState(asset.AutoID, b)
	if err != nil {
		return nil, false
	}

	return b, true
}

func ZSGCCipherEquipment(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) < 1 {
		return shim.Error("The number of args is incorrect.")
	}

	var input_cipher lib.Input_Cipher_equipment
	err := json.Unmarshal([]byte(args[0]), &input_cipher)
	if err != nil {
		return shim.Error("Error occur when unmarshal the bytes information to struct.")
	}

	addAsset, bl := putZSGCCipherEquipment(stub, input_cipher, args[2], args[3])
	if !bl {
		return shim.Error("Error occur when put asset equipment information in.")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	//return shim.Success([]byte("Save asset equipment information success."))
	return shim.Success(addAsset)
}

func QueryCipherEquipmentByHash(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("The number of args is incorrect")
	}

	ContractInformation_OnlyHash := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\", \"contract_information_only_hash\":\"%s\"}}", DOC_TYPE, ContractInformation_OnlyHash)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	if queryResults == nil {
		return shim.Error("the hash is not exist!")
	}
	return shim.Success(queryResults)
}
