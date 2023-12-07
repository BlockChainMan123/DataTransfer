package routers

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

func GetAllAssetEquipment(stub shim.ChaincodeStubInterface) peer.Response {

	startKey := ""
	endKey := ""

	resultsIterator, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return shim.Error(err.Error())
	}

	total, err := computeTotalForrange(stub, startKey, endKey)

	bufferWithPaginationInfo := addPaginationMetadataToQueryResults(buffer, total)

	fmt.Printf("- GetAssetEquipmentByRange queryResult:\n%s\n", bufferWithPaginationInfo.String())

	return shim.Success(buffer.Bytes())
}

func GetAllCipherEquipmentWithPagination(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	startKey := ""
	endKey := ""
	//return type of ParseInt is int64
	pageSize, err := strconv.ParseInt(args[0], 10, 32)
	if err != nil {
		return shim.Error(err.Error())
	}
	//return type of ParseInt is int64
	start, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return shim.Error(err.Error())
	}
	bookmark := ""
	var iterator shim.StateQueryIteratorInterface

	for currentPage := 0; currentPage < int(start); currentPage++ {
		resultsIterator, responseMetadata, err := stub.GetStateByRangeWithPagination(startKey, endKey, int32(pageSize), bookmark)
		iterator = resultsIterator

		if err != nil {
			return shim.Error(err.Error())
		}
		defer resultsIterator.Close()

		bookmark = responseMetadata.Bookmark
	}

	buffer, err := constructQueryResponseFromIterator(iterator)
	if err != nil {
		return shim.Error(err.Error())
	}

	total, err := computeTotalForrange(stub, startKey, endKey)

	bufferWithPaginationInfo := addPaginationMetadataToQueryResults(buffer, total)

	fmt.Printf("- GetAllPlanEquipmentWithPagination queryResult:\n%s\n", bufferWithPaginationInfo.String())

	return shim.Success(buffer.Bytes())
}


