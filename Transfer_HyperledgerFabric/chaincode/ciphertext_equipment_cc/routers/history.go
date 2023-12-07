package routers

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

func GetHistoryCipherWithPagination(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) < 3 {
		return shim.Error("The number of args is incorrect not 3.")
	}

	assetID := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\", \"id\":\"%s\"}}", DOC_TYPE, assetID)

	//return type of ParseInt is int64
	pageSize, err := strconv.ParseInt(args[1], 10, 32)
	if err != nil {
		return shim.Error(err.Error())
	}

	//return type of ParseInt is int64
	start, err := strconv.ParseInt(args[2], 10, 64)
	if err != nil {
		return shim.Error(err.Error())
	}

	queryResults, err := getQueryResultForQueryStringWithPagination(stub, queryString, int32(pageSize), int(start))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)

}

/*
func historyGetQueryResultForQueryStringWithPagination(stub shim.ChaincodeStubInterface, queryString string, pageSize int32, start int) ([]byte, error) {
	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)
	bookmark := ""
	var iterator shim.StateQueryIteratorInterface
	for currentPage := 0; currentPage < start; currentPage++ {
		resultsIterator, responseMetadata, err := stub.GetQueryResultWithPagination(queryString, pageSize, bookmark)
		iterator = resultsIterator
		if err != nil {
			return nil, err
		}
		defer resultsIterator.Close()

		bookmark = responseMetadata.Bookmark
	}

	buffer, err := historyConstructQueryResponseFromIterator(iterator)
	if err != nil {
		return nil, err
	}

	total, err := computeTotalForrich(stub, queryString)
	bufferWithPaginationInfo := addPaginationMetadataToQueryResults(buffer, total)

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", bufferWithPaginationInfo.String())

	return buffer.Bytes(), nil
}

func historyConstructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	bArrayMemberAlreadyWritten := false

	buffer.WriteString(`{"historyList":[`)

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString(string(queryResponse.Value))
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString(`]`)

	fmt.Print("Query result: %s", buffer.String())
	return &buffer, nil
}*/
