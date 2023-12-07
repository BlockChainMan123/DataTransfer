/*
@Description:SDK调用链码，对单个资产的增查删改
*/

package sdk_service

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"strconv"
)

// ===========================================================================
// ZSGCPlainEquipment SDK
// ===========================================================================

func (s *ServiceSetup) zsgcPlainEquipmentByIDAndResponse(asset Input_Plain_equipment, StatusCode string, FunctionPath string, OperatingContract string) ([]byte, error) {
	eventID := "eventAddAsset"
	reg, nofifier := regitserEvent(s.Client, s.ChaincodeID, eventID)
	defer s.Client.UnregisterChaincodeEvent(reg)

	b, err := json.Marshal(asset)
	if err != nil {
		return []byte{0x00}, fmt.Errorf("指定的Plain_Equipment序列化时发生错误")
	}

	req := channel.Request{
		ChaincodeID: s.ChaincodeID,
		Fcn:         "ZSGCPlainEquipment",
		Args:        [][]byte{b, []byte(eventID), []byte(StatusCode), []byte(FunctionPath), []byte(OperatingContract)},
	}
	response, err := s.Client.Execute(req)
	if err != nil {
		return []byte{0x00}, nil
	}

	err = eventResult(nofifier, eventID)
	if err != nil {
		return []byte{0x00}, nil
	}

	//return string(response.TransactionID), nil
	return response.Payload, nil

}

func (s *ServiceSetup) ZSGCPlainEquipment_SDK(asset Input_Plain_equipment, StatusCode string, FunctionPath string, OperatingContract string) (interface{}, error) {
	result, err := s.zsgcPlainEquipmentByIDAndResponse(asset, StatusCode, FunctionPath, OperatingContract)
	if err != nil {
		return nil, err
	}
	var ass Plain_equipment
	err = json.Unmarshal(result, &ass)
	if err != nil {
		return nil, err
	}
	return ass, nil

}

// ===========================================================================
// QueryPlainEquipmentByHash SDK
// ===========================================================================
func (s *ServiceSetup) QueryPlainEquipmentByHash_SDK(ContractInformation_OnlyHash string) (interface{}, error) {

	result, err := s.QueryPlainEquipmentByHashAndResponse(ContractInformation_OnlyHash)
	if err != nil {
		return nil, err
	}
	var list ListResponse
	err = json.Unmarshal(result, &list)
	if err != nil {
		return nil, err
	}
	if list.Total == 0 {
		return nil, fmt.Errorf("the hash is not exist!")

	}

	return list.Resultslist, nil
}

func (s *ServiceSetup) QueryPlainEquipmentByHashAndResponse(BlockInformation_BlockCode string) ([]byte, error) {
	req := channel.Request{
		ChaincodeID: s.ChaincodeID,
		Fcn:         "QueryPlainEquipmentByHash",
		Args:        [][]byte{[]byte(BlockInformation_BlockCode)},
	}

	respone, err := s.Client.Query(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return respone.Payload, nil

}

func (s *ServiceSetup) GetHistoryPlainWithPagination_SDK(assetid string, pageSize int32, start int) (interface{}, error) {

	result, err := s.GetHistoryPlainWithPaginationAndResponse(assetid, pageSize, start)
	if err != nil {
		return nil, err
	}

	var list ListResponse
	err = json.Unmarshal(result, &list)
	if err != nil {
		return nil, err
	}
	if list.Total == 0 {
		return nil, fmt.Errorf("the history id is not exist!")
	}

	return list, nil
}

func (s *ServiceSetup) GetHistoryPlainWithPaginationAndResponse(assetid string, pageSize int32, start int) ([]byte, error) {
	// return type of FormatInt is string
	pstr := strconv.FormatInt(int64(pageSize), 10)
	sstr := strconv.FormatInt(int64(start), 10)

	req := channel.Request{
		ChaincodeID: s.ChaincodeID,
		Fcn:         "GetHistoryPlainWithPagination",
		Args:        [][]byte{[]byte(assetid), []byte(pstr), []byte(sstr)},
	}

	respone, err := s.Client.Query(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return respone.Payload, nil

}
