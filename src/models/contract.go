package models

type ContractResp struct {
	BlockHeight int32  `json:"block_height"`
	TxId        string `json:"tx_id"`
}

type InstructionResp struct {
	InstructionId      string `json:"instruction_id"`
	InstructionType    string `json:"instruction_type"`
	InstructionContent string `json:"instruction_content"`
	DebrisId           string `json:"debris_id"`
	DebrisName         string `json:"debris_name"`
	SatelliteId        string `json:"satellite_id"`
	SatelliteName      string `json:"satellite_name"`
	Treaten            string `json:"treaten"`
}

type InstructionContractResp struct {
	Instructions []*InstructionResp `json:"instructions"`
	ContractResp
}
