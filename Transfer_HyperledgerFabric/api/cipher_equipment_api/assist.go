package cipher_equipment_api

import (
	"github.com/Transfer_HyperledgerFabric/sdk_service"
	"strconv"
)

type InputCipherList struct {
	ContractInformation_FunctionPath      string                            `json:"contract_information_function_path"`
	ContractInformation_OperatingContract string                            `json:"contract_information_operating_contract"`
	Value                                 sdk_service.Input_Cipher_equipment `json:"value"`
}



type Cipher_ID struct {
	ID string `json:"id"`
}

type HashString struct {
	ContractInformation_OnlyHash string `json:"contract_information_only_hash"`
}

type IdPageString struct {
	ID       string `json:"id"`
	PageSize string `json:"pagesize"`
	Start    string `json:"start"`
}

type RichStringWithPage struct {

	PlanName   string `json:"planname"`
	ParentPlanID   string `json:"parentplanid"`
	TypeID        string `json:"typeid"`
	FunctionID  string `json:"functionid"`

	PageSize string `json:"pagesize"`
	Start    string `json:"start"`

}

func PageSizeAndStart(page_size string, start string) (int32, int) {

	PageSize := int64(10)
	Start := int64(1)

	if (len(page_size) >= 1) {
		temp, err := strconv.ParseInt(page_size, 10, 32)
		if err == nil {
			PageSize = temp
		}
	}
	if (len(start) >= 1) {
		temp, err := strconv.ParseInt(start, 10, 64)
		if err == nil {
			Start = temp
		}
	}

	// ensure PageSize bigger than 1
	if PageSize < 1 {
		PageSize = 1
	}
	// ensure Start bigger than 1
	if Start < 1 {
		Start = 1
	}
	return int32(PageSize), int(Start)
}
