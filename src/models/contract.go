package models

type ContractResp struct {
	BlockHeight int32  `json:"blockHeight"`
	ChainTime   int64  `json:"chainTime"`
	TxId        string `json:"txId"`
}

type InstructionResp struct {
	InstructionId      string `json:"instructionId"`
	InstructionType    string `json:"instructionType"`
	InstructionContent string `json:"instructionContent"`
	DebrisId           string `json:"debrisId"`
	DebrisName         string `json:"debrisName"`
	SatelliteId        string `json:"satelliteId"`
	SatelliteName      string `json:"satelliteName"`
	Treaten            string `json:"treaten"`
}

type InstructionContractResp struct {
	Instructions []*InstructionResp `json:"instructions"`
	ContractResp
}
