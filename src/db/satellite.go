package db

type SatelliteRunState int32

const (
	NORMAL SatelliteRunState = iota + 1
	DEVIATE
	CLOSED
)

const (
	NORMAL_STR  = "正常"
	DEVIATE_STR = "偏轨"
	CLOSED_STR  = "关闭"
)

const SATELLITE_TABLE_NAME = "satellite"

var SatelliteRunStateName = map[SatelliteRunState]string{
	NORMAL:  NORMAL_STR,
	DEVIATE: DEVIATE_STR,
	CLOSED:  CLOSED_STR,
}

var SatelliteRunStateValue = map[string]SatelliteRunState{
	NORMAL_STR:  NORMAL,
	DEVIATE_STR: DEVIATE,
	CLOSED_STR:  CLOSED,
}

type Satellite struct {
	GeneralField
	BlockChainField
	SatelliteId   string `gorm:"uniqueIndex"`
	SatelliteName string
	OrbitId       string
	MeanAnomaly   float64
	Speed         float64
	RunState      SatelliteRunState
}

func (s *Satellite) TableName() string {
	return SATELLITE_TABLE_NAME
}

func init() {
	satellite := new(Satellite)
	SchemaSlice = append(SchemaSlice, &satellite)
}
