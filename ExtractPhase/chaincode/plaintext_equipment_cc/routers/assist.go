/*
@Description:辅助功能路由，供其他路由调用
注:供其他路由调用，因此函数采用小写
*/

package routers

import (
	"bytes"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// ===========================================================================================
// constructQueryResponseFromIterator constructs a JSON array containing query results from
// a given result iterator
// ===========================================================================================
func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	bArrayMemberAlreadyWritten := false

	buffer.WriteString(`{"resultsList":[`)

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
}

// =======================================================================
// getQueryResultForQueryString: executes the passed in query string.
// Result set is built and returned as a byte array containing the JSON results.
// =======================================================================
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return nil, err
	}

	total, err := computeTotalForrich(stub, queryString)
	bufferWithPaginationInfo := addPaginationMetadataToQueryResults(buffer, total)

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", bufferWithPaginationInfo.String())

	return buffer.Bytes(), nil
}

// =========================================================================================
// getQueryResultForQueryStringWithPagination executes the passed in query string with
// pagination info. However, GetQueryResultWithPagination chaincode interface only support
// the bookmark instead of page number. We transform it to support page number query.
// =========================================================================================
func getQueryResultForQueryStringWithPagination(stub shim.ChaincodeStubInterface, queryString string, pageSize int32, start int) ([]byte, error) {
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

	buffer, err := constructQueryResponseFromIterator(iterator)
	if err != nil {
		return nil, err
	}

	total, err := computeTotalForrich(stub, queryString)
	bufferWithPaginationInfo := addPaginationMetadataToQueryResults(buffer, total)

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", bufferWithPaginationInfo.String())

	return buffer.Bytes(), nil
}

// ===========================================================================================
// addPaginationMetadataToQueryResults adds QueryResponseMetadata
//
// ===========================================================================================
func addPaginationMetadataToQueryResults(buffer *bytes.Buffer, total int) *bytes.Buffer {
	buffer.WriteString(fmt.Sprintf(",\"total\":%v}", total))
	/*
	buffer.WriteString(",\"ResponseMetadata\":{\"RecordsCount\":")
	buffer.WriteString("\"")
	buffer.WriteString(fmt.Sprintf("%v", responseMetadata.FetchedRecordsCount))
	buffer.WriteString("\"")
	buffer.WriteString(", \"Bookmark\":")
	buffer.WriteString("\"")
	buffer.WriteString(responseMetadata.Bookmark)
	buffer.WriteString("\"}}")
*/
	return buffer
}

// 富查询的总数计算
// ===========================================================================================
// computeTotalForrich executes to return total that is the total number of query results. However,
// the computeTotal function is not excellent for the reason of computeTotal query iterators
// from start to end. We have not seen there is FABRIC chaincode interface to release querying
// the number of query results up to now.
// ===========================================================================================
func computeTotalForrich(stub shim.ChaincodeStubInterface, queryString string) (int, error) {
	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return 0, err
	}
	defer resultsIterator.Close()

	total := 0
	for resultsIterator.HasNext() {
		_, err := resultsIterator.Next()
		if err != nil {
			return 0, err
		}
		total++
	}
	// total is the total number of query results

	return total, nil
}

// 范围查询的总数计算
// ===========================================================================================
// computeTotalForrange executes to return total that is the total number of query results. However,
// the computeTotal function is not excellent for the reason of computeTotal query iterators
// from start to end. We have not seen there is FABRIC chaincode interface to release querying
// the number of query results up to now.
// ===========================================================================================
func computeTotalForrange(stub shim.ChaincodeStubInterface, startkey string, endkey string) (int, error) {
	resultsIterator, err := stub.GetStateByRange(startkey, endkey)
	if err != nil {
		return 0, err
	}
	defer resultsIterator.Close()

	total := 0
	for resultsIterator.HasNext() {
		_, err := resultsIterator.Next()
		if err != nil {
			return 0, err
		}
		total++
	}
	// total is the total number of query results

	return total, nil
}