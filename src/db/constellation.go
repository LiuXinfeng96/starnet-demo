package db

const CONSTELLATION_TABLE_NAME = "constellation"

type Constellation struct {
	GeneralField
	BlockChainField
	ConstellationId    string `gorm:"index"`
	ConstellationName  string `gorm:"index"`
	SatelliteTotalNum  int32
	SatelliteUpNum     int32
	SatelliteDownNum   int32
	SatelliteLinkState string
}

func (c *Constellation) TableName() string {
	return CONSTELLATION_TABLE_NAME
}

func init() {
	constellation := new(Constellation)
	TableSlice = append(TableSlice, &constellation)
}
