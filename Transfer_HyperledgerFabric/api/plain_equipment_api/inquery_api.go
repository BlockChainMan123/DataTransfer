package plain_equipment_api

import (
	"encoding/json"
	"github.com/Transfer_HyperledgerFabric/sdk_service"
	"strconv"
)

func QueryPlainEquipmentByHash_API(setup sdk_service.ServiceSetup, hashstring string) (interface{}, error) {
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

func GetHistoryPlainWithPagination_API(setup sdk_service.ServiceSetup, idpagestring string) (interface{}, error) {
	var input IdPageString
	err := json.Unmarshal([]byte(idpagestring), &input)
	if err != nil {
		return "", err
	}
	pagesize, start := PageSizeAndStart(input.PageSize, input.Start)
	results, err := setup.GetHistoryPlainWithPagination_SDK(input.ID, pagesize, start)
	if err != nil {
		return "", err
	}

	return results, nil
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