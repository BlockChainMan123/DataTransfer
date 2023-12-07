package plain_equipment_api

import (
	"encoding/json"
	"github.com/Transfer_HyperledgerFabric/sdk_service"
)

func AddPlainEquipment_API(setup sdk_service.ServiceSetup, input_plain_string string) (interface{}, error) {

	var input_asset InputPlainList
	err := json.Unmarshal([]byte(input_plain_string), &input_asset)
	if err != nil {
		return "", err
	}
	var asset sdk_service.Input_Plain_equipment
	asset = input_asset.Value

	result, err := setup.ZSGCPlainEquipment_SDK(asset, input_asset.ID, input_asset.ContractInformation_FunctionPath, input_asset.ContractInformation_OperatingContract)
	if err != nil {
		return "", err
	}

	return result, nil
}

func FindPlainEquipment_API(setup sdk_service.ServiceSetup, input_plain_string string) (interface{}, error) {


	var input_asset InputPlainList
	err := json.Unmarshal([]byte(input_plain_string), &input_asset)
	if err != nil {
		return "", err
	}
	var asset sdk_service.Input_Plain_equipment
	asset = input_asset.Value

	result, err := setup.ZSGCPlainEquipment_SDK(asset, input_asset.ID, input_asset.ContractInformation_FunctionPath, input_asset.ContractInformation_OperatingContract)
	if err != nil {
		return "", err
	}

	return result, nil
}




