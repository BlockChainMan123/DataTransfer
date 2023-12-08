/*
 * @Description: 定义的数据结构体、常量
 */

package lib

const DOC_TYPE = "cipher_equipment_object"

type Input_Cipher_equipment struct {
	DocType string `json:"doc_type"`

	//ciphertext
	ID string `json:"id"`

	Address  string `json:"address"`  //address of requestor
	Sk     string `json:"sk"`     //private key of requestor
	Request       string `json:"request"`       //request information
}




type Cipher_equipment struct {
	DocType string `json:"doc_type"`

	AutoID string `json:"autoid"` //key is the AutoID

	//contract info
	ContractInformation_FunctionPath      string `json:"contract_information_function_path"`
	ContractInformation_OperatingContract string `json:"contract_information_operating_contract"`
	ContractInformation_TimeStamp         string `json:"contract_information_time_stamp"`
	ContractInformation_OnlyHash          string `json:"contract_information_only_hash"`

	//block info
	BlockInformation_TimeStamp  string `json:"block_information_time_stamp"`
	BlockInformation_BlockPlace string `json:"block_information_block_place"`
	BlockInformation_BlockCode  string `json:"block_information_block_hash"`
	BlockInformation_ChainPlace string `json:"block_information_chain_place"`

	//ciphertext
	ID string `json:"id"`

	Address  string `json:"address"`  //address of requestor
	Sk     string `json:"sk"`     //private key of requestor
	Request       string `json:"request"`       //request information
}


