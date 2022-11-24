package db

const OPERATION_TABLE_NAME = "operation"

type Operation struct {
	GeneralField
	BlockChainField
	Operator        string
	IsAutomatic     bool
	OperationTime   int64
	OperationRecord string `gorm:"type:longtext"`
	SatelliteId     string
	SatelliteName   string
}

func (o *Operation) TableName() string {
	return OPERATION_TABLE_NAME
}

func init() {
	operation := new(Operation)
	SchemaSlice = append(SchemaSlice, &operation)
}
