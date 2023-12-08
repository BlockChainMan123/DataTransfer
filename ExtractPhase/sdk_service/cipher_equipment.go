package sdk_service

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"strconv"
)

func (s *ServiceSetup) zsgcCipherEquipmentByIDAndResponse(plan Input_Cipher_equipment, FunctionPath string, OperatingContract string) ([]byte, error) {
	eventID := "eventAddPlan"
	reg, nofifier := regitserEvent(s.Client, s.ChaincodeID2, eventID)
	defer s.Client.UnregisterChaincodeEvent(reg)

	b, err := json.Marshal(plan)
	if err != nil {
		return []byte{0x00}, fmt.Errorf("指定的Plans_Equipment序列化时发生错误")
	}

	req := channel.Request{
		ChaincodeID: s.ChaincodeID2,
		Fcn:         "ZSGCCipherEquipment",
		Args:        [][]byte{b, []byte(eventID), []byte(FunctionPath), []byte(OperatingContract)},
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

func (s *ServiceSetup) ZSGCCipherEquipment_SDK(plan Input_Cipher_equipment, FunctionPath string, OperatingContract string) (interface{}, error) {
	result, err := s.zsgcCipherEquipmentByIDAndResponse(plan, FunctionPath, OperatingContract)
	if err != nil {
		return nil, err
	}
	var pla Cipher_equipment
	err = json.Unmarshal(result, &pla)
	if err != nil {
		return nil, err
	}
	return pla, nil

}

func (s *ServiceSetup) QueryCipherEquipmentByHash_SDK(ContractInformation_OnlyHash string) (interface{}, error) {

	result, err := s.QueryCipherEquipmentByHashAndResponse(ContractInformation_OnlyHash)
	if err != nil {
		return nil, err
	}
	var list PlanListResponse
	err = json.Unmarshal(result, &list)
	if err != nil {
		return nil, err
	}
	if list.Total == 0 {
		return nil, fmt.Errorf("the hash is not exist!")

	}

	return list.Resultslist, nil
}

func (s *ServiceSetup) QueryCipherEquipmentByHashAndResponse(BlockInformation_BlockCode string) ([]byte, error) {
	req := channel.Request{
		ChaincodeID: s.ChaincodeID2,
		Fcn:         "QueryCipherEquipmentByHash",
		Args:        [][]byte{[]byte(BlockInformation_BlockCode)},
	}

	respone, err := s.Client.Query(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return respone.Payload, nil

}

func (s *ServiceSetup) GetHistoryCipherWithPagination_SDK(planid string, pageSize int32, start int) (interface{}, error) {

	result, err := s.GetHistoryCipherWithPaginationAndResponse(planid, pageSize, start)
	if err != nil {
		return nil, err
	}

	var list PlanListResponse
	err = json.Unmarshal(result, &list)
	if err != nil {
		return nil, err
	}
	if list.Total == 0 {
		return nil, fmt.Errorf("the history id is not exist!")
	}

	return list, nil
}

func (s *ServiceSetup) GetHistoryCipherWithPaginationAndResponse(planid string, pageSize int32, start int) ([]byte, error) {
	// return type of FormatInt is string
	pstr := strconv.FormatInt(int64(pageSize), 10)
	sstr := strconv.FormatInt(int64(start), 10)

	req := channel.Request{
		ChaincodeID: s.ChaincodeID2,
		Fcn:         "GetHistoryCipherWithPagination",
		Args:        [][]byte{[]byte(planid), []byte(pstr), []byte(sstr)},
	}

	respone, err := s.Client.Query(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return respone.Payload, nil

}
