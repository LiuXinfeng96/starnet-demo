package db

const COMMSTATE_TABLE_NAME = "comm_state"

type CommState struct {
	GeneralField
	BlockChainField
	SatelliteId   string `gorm:"index"`
	SatelliteName string `gorm:"index"`
	OrbitId       string
	CommState     string
	CommBandwidth string
	CommDelay     string
	CommPort      string
	LinkLoad      string
}

func (c *CommState) TableName() string {
	return COMMSTATE_TABLE_NAME
}

func init() {
	commState := new(CommState)
	TableSlice = append(TableSlice, &commState)
}
