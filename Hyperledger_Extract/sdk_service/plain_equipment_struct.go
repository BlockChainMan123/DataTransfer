package sdk_service

type Input_Plain_equipment struct {
	DocType string `json:"doc_type"`
	//ciphertext
	ID string `json:"id"`

	Address  string `json:"address"`  //address of requestor
	Pk     string `json:"pk"`     //private key of requestor
	Request       string `json:"request"`       //request information

}

type Plain_equipment struct {
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
	Pk     string `json:"Pk"`     //public key of requestor
	Request       string `json:"request"`       //request information
}


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

