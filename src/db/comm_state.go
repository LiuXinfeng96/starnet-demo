package db

const COMMSTATE_TABLE_NAME = "comm_state"

type CommState struct {
	GeneralField
	BlockChainField
	SatelliteId   string
	SatelliteName string
	OrbitId       string
	CommState     string
	CommBandwidth string
	CommDelay     string
	LinkLoad      string
}

func (c *CommState) TableName() string {
	return COMMSTATE_TABLE_NAME
}

func init() {
	commState := new(CommState)
	SchemaSlice = append(SchemaSlice, &commState)
}
