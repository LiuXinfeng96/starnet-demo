package db

const FAULT_TABLE_NAME = "fault"

type FaultType int32

const (
	AType FaultType = iota + 1
	BType
	CType
	DType
	EType
)

const (
	AType_STR = "A类"
	BType_STR = "B类"
	CType_STR = "C类"
	DType_STR = "D类"
	EType_STR = "E类"
)

var FaultTypeName = map[FaultType]string{
	AType: AType_STR,
	BType: BType_STR,
	CType: CType_STR,
	DType: DType_STR,
	EType: EType_STR,
}

var FaultTypeValue = map[string]FaultType{
	AType_STR: AType,
	BType_STR: BType,
	CType_STR: CType,
	DType_STR: DType,
	EType_STR: EType,
}

type Fault struct {
	GeneralField
	BlockChainField
	SatelliteId      string `gorm:"index"`
	SatelliteName    string `gorm:"index"`
	OrbitId          string
	FaultType        FaultType
	FaultTime        int64
	FaultDescription string `gorm:"type:longtext"`
	RepairState      State
}

func (f *Fault) TableName() string {
	return FAULT_TABLE_NAME
}

func init() {
	fault := new(Fault)
	TableSlice = append(TableSlice, &fault)
}
