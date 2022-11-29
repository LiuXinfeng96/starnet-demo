package models

type TokenReq struct {
	Token string `token`
}

type PageReq struct {
	PageSize int32 `pageSize`
	Page     int32 `page`
}

type QueryReq struct {
	SortType         string `json: "sortType"`
	SearchConditions string `json: "searchConditions"`
}

type QueryListReq struct {
	QueryReq
	PageReq
	TokenReq
}

type RegisterReq struct {
	UserName     string `json: "userName"`
	UserPwd      string `json: "userPwd"`
	UserNickName string `json: "userNickName"`
	UserPhoneNum string `json: "userPhoneNum"`
	UserEmail    string `json: "userEmail"`
}

type LoginReq struct {
	UserName string `json: "userName"`
	UserPwd  string `json: "userPwd"`
}

type UserReq struct {
	UserName string `json: "userName"`
	TokenReq
}

type AddDebirsReq struct {
	DebrisId   string  `json: "debrisId"`
	DebrisName string  `json: "debrisName"`
	Angle      float64 `json: "angle"`
	Speed      float64 `json: "speed"`
	Height     float64 `json: "height"`
	Volume     float64 `json: "volume"`
	Type       string  `json: "type"`
	TokenReq
}

type GetDebrisReq struct {
	DebrisId string `json: "debrisId"`
	TokenReq
}

type AddInstructionReq struct {
	InstructtionId      string `json: "instructtionId"`
	ExecInstructionTime int64  `json: "execInstructionTime"`
	DebrisId            string `json: "debrisId"`
	DebrisName          string `json: "debrisName"`
	SatelliteId         string `json: "satelliteId"`
	SatelliteName       string `json: "satelliteName"`
	InstructionContent  string `json: "instructionContent"`
	TokenReq
}

type GetInstructionReq struct {
	InstructionId string `json: "instructionId"`
	TokenReq
}

type GetInstructionResultReq struct {
	Id int32 `json: "id"`
	TokenReq
}

type AddOrbitReq struct {
	OrbitId                string  `json: "orbitId"`
	OrbitType              string  `json: "orbitType"`
	OrbitEccentricity      float64 `json: "orbitEccentricity"`
	OrbitSemiMajorAxis     float64 `json: "orbitSemiMajorAxis"`
	OrbitAngle             float64 `json: "orbitAngle"`
	AscendingNodeLongitude float64 `json: "ascendingNodeLongitude"`
	Perigee                float64 `json: "perigee"`
	TokenReq
}

type AddConstellationReq struct {
	ConstellationId    string `json: "constellationId"`
	ConstellationName  string `json: "constellationName"`
	SatelliteTotalNum  int32  `json: "satelliteTotalNum"`
	SatelliteUpNum     int32  `json: "satelliteUpNum"`
	SatelliteDownNum   int32  `json: "satelliteDownNum"`
	SatelliteLinkState string `json: "satelliteLinkState"`
	TokenReq
}

type GetConstellationReq struct {
	ConstellationId string `json: "constellationId"`
	TokenReq
}

type GetOperationReq struct {
	SatelliteId string `json: "satelliteId"`
	TokenReq
}

type AddSatelliteState struct {
	SatelliteId   string  `json: "satelliteId"`
	SatelliteName string  `json: "satelliteName"`
	OrbitId       string  `json: "orbitId"`
	MeanAnomaly   float64 `json: "meanAnomaly"`
	Speed         float64 `json: "speed"`
	RunState      string  `json: "runState"`
	TokenReq
}

type GetSatelliteStateReq struct {
	ConstellationId string `json: "constellationId"`
	TokenReq
}

type AddControlsReq struct {
	SatelliteId          string  `json: "satelliteId"`
	SatelliteName        string  `json: "satelliteName"`
	SatelliteAttitude    string  `json: "satelliteAttitude"`
	SatelliteTemperature float64 `json: "satelliteTemperature"`
	SatellitePower       float64 `json: "satellitePower"`
	TokenReq
}

type GetControlsReq struct {
	SatelliteId string `json: "satelliteId"`
	TokenReq
}

type AddFaultReq struct {
	SatelliteId      string `json: "satelliteId"`
	SatelliteName    string `json: "satelliteName"`
	OrbitId          string `json: "orbitId"`
	FaultType        string `json: "faultType"`
	FaultDescription string `json: "faultDescription"`
	FaultTime        int64  `json: "faultTime"`
	RepairState      string `json: "repairState"`
	TokenReq
}

type GetFaultReq struct {
	SatelliteId string `json: "satelliteId"`
	TokenReq
}

type AddNetStateReq struct {
	SatelliteId      string  `json: "satelliteId"`
	SatelliteName    string  `json: "satelliteName"`
	OrbitId          string  `json: "orbitId"`
	NetworkSegment   string  `json: "networkSegment"`
	NetworkState     string  `json: "networkState"`
	NetworkBandwidth float32 `json: "networkBandwidth"`
	TokenReq
}

type GetNetStateReq struct {
	SatelliteId string `json: "satelliteId"`
	TokenReq
}

type AddCommStateReq struct {
	SatelliteId   string `json: "satelliteId"`
	SatelliteName string `json: "satelliteName"`
	OrbitId       string `json: "orbitId"`
	CommState     string `json: "commState"`
	CommPort      string `json: "commPort"`
	CommBandwidth string `json: "commBandwidth"`
	CommDelay     string `json: "commDelay"`
	LinkLoad      string `json: "linkLoad"`
	TokenReq
}

type GetCommStateReq struct {
	SatelliteId string `json: "satelliteId"`
	TokenReq
}
