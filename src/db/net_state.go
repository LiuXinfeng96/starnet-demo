package db

const NETSTATE_TABLE_NAME = "net_state"

type NetState struct {
	GeneralField
	BlockChainField
	SatelliteId      string
	SatelliteName    string
	OrbitId          string
	NetworkSegment   string
	NetworkState     string
	NetworkBandwidth string
}

func (n *NetState) TableName() string {
	return NETSTATE_TABLE_NAME
}

func init() {
	netState := new(NetState)
	TableSlice = append(TableSlice, &netState)
}
