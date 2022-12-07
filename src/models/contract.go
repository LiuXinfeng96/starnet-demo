package models

type ContractResp struct {
	BlockHeight int32 `json:"blockHeight"`
	ChainTime   int64 `json:"chainTime"`
}

type InstructionContractResp struct {
	InstructionId      string `json:"instructionId"`
	InstructionContent string `json:"instructionContent"`
	GenInstructionTime int64  `json:"genInstructionTime"`
	DebrisId           string `json:"debrisId"`
	DebrisName         string `json:"debrisName"`
	SatelliteId        string `json:"satelliteId"`
	SatelliteName      string `json:"satelliteName"`
	ContractResp
}
