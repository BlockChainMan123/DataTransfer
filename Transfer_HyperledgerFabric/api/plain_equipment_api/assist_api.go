package plain_equipment_api

import "github.com/Transfer_HyperledgerFabric/sdk_service"

type Plain_ID struct {
	ID string `json:"id"`
}

type HashString struct {
	ContractInformation_OnlyHash string `json:"contract_information_only_hash"`
}



type InputPlainList struct {
	ID                                    string `json:"id"`
	ContractInformation_FunctionPath      string `json:"contract_information_function_path"`
	ContractInformation_OperatingContract string `json:"contract_information_operating_contract"`

	Value                                 sdk_service.Input_Plain_equipment `json:"value"`
}

type IdPageString struct {
	ID       string `json:"id"`
	PageSize string `json:"pagesize"`
	Start    string `json:"start"`
}


