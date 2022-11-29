package db

const FAULT_TABLE_NAME = "fault"

type Fault struct {
	GeneralField
	BlockChainField
	SatelliteId      string
	SatelliteName    string
	OrbitId          string
	FaultType        string
	FaultTime        int64
	FaultDescription string `gorm:"type:longtext"`
	RepairState      string
}

func (f *Fault) TableName() string {
	return FAULT_TABLE_NAME
}

func init() {
	fault := new(Fault)
	TableSlice = append(TableSlice, &fault)
}
