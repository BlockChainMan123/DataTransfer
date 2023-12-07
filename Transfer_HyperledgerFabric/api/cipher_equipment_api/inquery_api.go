package cipher_equipment_api

import (
	"encoding/json"
	"github.com/Transfer_HyperledgerFabric/sdk_service"
)

func QueryCipherEquipmentByHash_API(setup sdk_service.ServiceSetup, hashstring string) (interface{}, error) {
	var input HashString
	err := json.Unmarshal([]byte(hashstring), &input)
	if err != nil {
		return "", err
	}
	results, err := setup.QueryPlainEquipmentByHash_SDK(input.ContractInformation_OnlyHash)
	if err != nil {
		return "", err
	}

	return results, nil
}

func GetHistoryCipherWithPagination_API(setup sdk_service.ServiceSetup, idpagestring string) (interface{}, error) {
	var input IdPageString
	err := json.Unmarshal([]byte(idpagestring), &input)
	if err != nil {
		return "", err
	}
	pagesize, start := PageSizeAndStart(input.PageSize, input.Start)
	results, err := setup.GetHistoryCipherWithPagination_SDK(input.ID, pagesize, start)
	if err != nil {
		return "", err
	}

	return results, nil
}

func AddCipherEquipment_API(setup sdk_service.ServiceSetup, input_cipher_string string) (interface{}, error) {

	var input_asset InputCipherList
	err := json.Unmarshal([]byte(input_cipher_string), &input_asset)
	if err != nil {
		return "", err
	}
	var asset sdk_service.Input_Cipher_equipment
	asset = input_asset.Value

	result, err := setup.ZSGCCipherEquipment_SDK(asset,input_asset.ContractInformation_FunctionPath,input_asset.ContractInformation_OperatingContract)
	if err != nil {
		return "", err
	}

	return result, nil
}

func FindCipherEquipment_API(setup sdk_service.ServiceSetup, input_cipher_string string) (interface{}, error) {


	var input_asset InputCipherList
	err := json.Unmarshal([]byte(input_cipher_string), &input_asset)
	if err != nil {
		return "", err
	}
	var asset sdk_service.Input_Cipher_equipment
	asset = input_asset.Value

	result, err := setup.ZSGCCipherEquipment_SDK(asset,input_asset.ContractInformation_FunctionPath,input_asset.ContractInformation_OperatingContract)
	if err != nil {
		return "", err
	}

	return result, nil
}


