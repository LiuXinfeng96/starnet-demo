package models

type StandardResp struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type StandardRespWithPage struct {
	StandardResp
	Total int32 `json:"total"`
}

const (
	RESP_CODE_SUCCESS = 200

	RESP_CODE_SERVER_ERROR = 500

	RESP_CODE_PARAMS_TYPE_ERROR = 101

	RESP_CODE_PARAMS_MISSING = 102

	RESP_CODE_PARAMS_FORMAT_ERROR = 103

	RESP_CODE_PARAMS_VALUE_ERROR = 104

	RESP_CODE_WITHOUT_PERMISSION = 301
)

const (
	RESP_MSG_SERVER_ERROR = "服务出错"

	RESP_MSG_PARAMS_TYPE_ERROR = "错误的参数类型"

	RESP_MSG_PARAMS_MISSING = "缺少必要参数"

	RESP_MSG_PARAMS_FORMAT_ERROR = "参数不符合规范"

	RESP_MSG_PARAMS_VALUE_ERROR = "参数值域不符合预期"

	RESP_MSG_WITHOUT_PERMISSION = "无权限访问"
)

type UserInfo struct {
	Id           int32  `json:"id"`
	UserName     string `json:"userName"`
	UserNickName string `json:"userNickName"`
	UserRole     string `json:"userRole"`
	UserPhoneNum string `json:"userPhoneNum"`
	UserEmail    string `json:"userEmail"`
	Expires      int32  `json:"expires"`
}

type LoginInfo struct {
	UserNickName string `json:"userNickName"`
	UserRole     string `json:"userRole"`
	Expires      int32  `json:"expires"`
	Token        string `json:"token"`
}

type BaseRespInfo struct {
	Id        int32 `json:"id"`
	LastTime  int64 `json:"lastTime"`
	ChainTime int64 `json:"chainTime"`
}

type HistoryRespInfo struct {
	ChainId     string `json:"chainId"`
	BlockHeight int32  `json:"blockHeight"`
	TxId        string `json:"txId"`
}

type DebrisInfo struct {
	DebrisId   string  `json:"debrisId"`
	DebrisName string  `json:"debrisName"`
	Angle      float64 `json:"angle"`
	Speed      float64 `json:"speed"`
	Height     float64 `json:"height"`
	Volume     float64 `json:"volume"`
	Type       string  `json:"type"`
	BaseRespInfo
}

type DebrisHistoryInfo struct {
	DebrisInfo
	HistoryRespInfo
}

type InstructionInfo struct {
	BaseRespInfo
	InstructtionId      string `json:"instructtionId"`
	InstructtionType    string `json:"instructtionType"`
	ExecInstructionTime int64  `json:"execInstructionTime"`
	InstructionSource   string `json:"instructionSource"`
	DebrisId            string `json:"debrisId"`
	DebrisName          string `json:"debrisName"`
	SatelliteId         string `json:"satelliteId"`
	SatelliteName       string `json:"satelliteName"`
	InstructionContent  string `json:"instructionContent"`
	GenInstructionTime  int64  `json:"genInstructionTime"`
}

type InstructionDetails struct {
	InstructionInfo
	ExecState string `json:"execState"`
	HistoryRespInfo
}

type InstructionResultInfo struct {
	InstructionInfo
	ExecState string `json:"execState"`
}

type OrbitInfo struct {
	BaseRespInfo
	OrbitId                string  `json:"orbitId"`
	OrbitType              string  `json:"orbitType"`
	OrbitEccentricity      float64 `json:"orbitEccentricity"`
	OrbitSemiMajorAxis     float64 `json:"orbitSemiMajorAxis"`
	OrbitAngle             float64 `json:"orbitAngle"`
	AscendingNodeLongitude float64 `json:"ascendingNodeLongitude"`
	Perigee                float64 `json:"perigee"`
}

type ConstellationInfo struct {
	BaseRespInfo
	ConstellationId    string `json:"constellationId"`
	ConstellationName  string `json:"constellationName"`
	SatelliteTotalNum  int32  `json:"satelliteTotalNum"`
	SatelliteUpNum     int32  `json:"satelliteUpNum"`
	SatelliteDownNum   int32  `json:"satelliteDownNum"`
	SatelliteLinkState string `json:"satelliteLinkState"`
}

type ConstellationHistoryInfo struct {
	ConstellationInfo
	HistoryRespInfo
}

type OperationInfo struct {
	BaseRespInfo
	Operator        string `json:"operator"`
	OperationTime   int64  `json:"operationTime"`
	OperationIp     string `json:"operationIp"`
	OperationRecord string `json:"operationRecord"`
	SatelliteId     string `json:"satelliteId"`
	SatelliteName   string `json:"satelliteName"`
}

type OperationHistoryInfo struct {
	OperationInfo
	HistoryRespInfo
}

type LoginLogInfo struct {
	BaseRespInfo
	UserName  string `json:"userName"`
	LoginTime int64  `json:"loginTime"`
	LoginIp   string `json:"loginIp"`
}

type SatelliteStateInfo struct {
	BaseRespInfo
	SatelliteId   string  `json:"satelliteId"`
	SatelliteName string  `json:"satelliteName"`
	OrbitId       string  `json:"orbitId"`
	MeanAnomaly   float64 `json:"meanAnomaly"`
	Speed         float64 `json:"speed"`
	RunState      string  `json:"runState"`
}

type SatelliteStateHistoryInfo struct {
	SatelliteStateInfo
	HistoryRespInfo
}

type ControlsInfo struct {
	BaseRespInfo
	SatelliteId          string  `json:"satelliteId"`
	SatelliteName        string  `json:"satelliteName"`
	SatelliteAttitude    string  `json:"satelliteAttitude"`
	SatelliteTemperature float64 `json:"satelliteTemperature"`
	SatellitePower       float64 `json:"satellitePower"`
}

type ControlsHistoryInfo struct {
	ControlsInfo
	HistoryRespInfo
}

type FaultInfo struct {
	BaseRespInfo
	SatelliteId      string `json:"satelliteId"`
	SatelliteName    string `json:"satelliteName"`
	OrbitId          string `json:"orbitId"`
	FaultType        string `json:"faultType"`
	FaultDescription string `json:"faultDescription"`
	FaultTime        int64  `json:"faultTime"`
	RepairState      string `json:"repairState"`
}

type FaultHistoryInfo struct {
	FaultInfo
	HistoryRespInfo
}

type NetStateInfo struct {
	BaseRespInfo
	SatelliteId      string  `json:"satelliteId"`
	SatelliteName    string  `json:"satelliteName"`
	OrbitId          string  `json:"orbitId"`
	NetworkSegment   string  `json:"networkSegment"`
	NetworkState     string  `json:"networkState"`
	NetworkBandwidth float32 `json:"networkBandwidth"`
}

type NetStateHistoryInfo struct {
	NetStateInfo
	HistoryRespInfo
}

type CommstateInfo struct {
	BaseRespInfo
	SatelliteId   string `json:"satelliteId"`
	SatelliteName string `json:"satelliteName"`
	OrbitId       string `json:"orbitId"`
	CommState     string `json:"commState"`
	CommPort      string `json:"commPort"`
	CommBandwidth string `json:"commBandwidth"`
	CommDelay     string `json:"commDelay"`
	LinkLoad      string `json:"linkLoad"`
}

type CommstateHistoryInfo struct {
	CommstateInfo
	HistoryRespInfo
}
