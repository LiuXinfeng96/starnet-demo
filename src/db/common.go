package db

var TableSlice = make([]interface{}, 0)

type GeneralField struct {
	Id       int32 `gorm:"primaryKey;autoIncrement"`
	LastTime int64 `gorm:"autoCreateTime;index"`
}

type BlockChainField struct {
	BlockHeight int32
	ChainId     string
	TxId        string
	ChainTime   int64
}
