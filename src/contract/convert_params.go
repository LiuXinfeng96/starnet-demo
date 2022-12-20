package contract

import (
	"starnet-demo/src/db"
	"starnet-demo/src/models"
	"strconv"

	"chainmaker.org/chainmaker/pb-go/v2/common"
)

func DebrisConvert(debris *db.Debris) []*common.KeyValuePair {
	params := make([]*common.KeyValuePair, 0)
	params = append(params,
		&common.KeyValuePair{
			Key:   "debrisId",
			Value: []byte(debris.DebrisId),
		}, &common.KeyValuePair{
			Key:   "debrisName",
			Value: []byte(debris.DebrisName),
		}, &common.KeyValuePair{
			Key:   "angle",
			Value: []byte(strconv.FormatFloat(debris.Angle, 'f', 10, 64)),
		}, &common.KeyValuePair{
			Key:   "speed",
			Value: []byte(strconv.FormatFloat(debris.Angle, 'f', 10, 64)),
		}, &common.KeyValuePair{
			Key:   "height",
			Value: []byte(strconv.FormatFloat(debris.Angle, 'f', 10, 64)),
		}, &common.KeyValuePair{
			Key:   "volume",
			Value: []byte(strconv.FormatFloat(debris.Angle, 'f', 10, 64)),
		}, &common.KeyValuePair{
			Key:   "debrisType",
			Value: []byte(db.DebrisTypeName[debris.Type]),
		},
	)
	return params
}

func InstructionConvert(ins *db.Instruction) []*common.KeyValuePair {
	params := make([]*common.KeyValuePair, 0)
	params = append(params,
		&common.KeyValuePair{
			Key:   "instructionId",
			Value: []byte(ins.InstructionId),
		}, &common.KeyValuePair{
			Key:   "instructionType",
			Value: []byte(db.InstructionTypeName[ins.Type]),
		}, &common.KeyValuePair{
			Key:   "instructionSource",
			Value: []byte(ins.InstructionSource),
		}, &common.KeyValuePair{
			Key:   "instructionContent",
			Value: []byte(ins.InstructionContent),
		}, &common.KeyValuePair{
			Key:   "debrisId",
			Value: []byte(ins.DebrisId),
		}, &common.KeyValuePair{
			Key:   "debrisName",
			Value: []byte(ins.DebrisName),
		}, &common.KeyValuePair{
			Key:   "satelliteId",
			Value: []byte(ins.SatelliteId),
		}, &common.KeyValuePair{
			Key:   "satelliteName",
			Value: []byte(ins.SatelliteName),
		}, &common.KeyValuePair{
			Key:   "execResult",
			Value: []byte(db.ExecStateName[ins.ExecState]),
		}, &common.KeyValuePair{
			Key:   "genInstructionTime",
			Value: []byte(strconv.Itoa(int(ins.GenInstructionTime))),
		}, &common.KeyValuePair{
			Key:   "execInstructionTime",
			Value: []byte(strconv.Itoa(int(ins.ExecInstructionTime))),
		}, &common.KeyValuePair{
			Key:   "treaten",
			Value: []byte(db.ThreatDegreeName[ins.Treaten]),
		},
	)
	return params
}

func OrbitConvert(orbit *db.Orbit) []*common.KeyValuePair {
	params := make([]*common.KeyValuePair, 0)
	params = append(params,
		&common.KeyValuePair{
			Key:   "orbitId",
			Value: []byte(orbit.OrbitId),
		}, &common.KeyValuePair{
			Key:   "orbitType",
			Value: []byte(orbit.OrbitType),
		}, &common.KeyValuePair{
			Key:   "orbitSemiMajorAxis",
			Value: []byte(strconv.FormatFloat(orbit.OrbitSemiMajorAxis, 'f', 10, 64)),
		}, &common.KeyValuePair{
			Key:   "orbitEccentricity",
			Value: []byte(strconv.FormatFloat(orbit.OrbitEccentricity, 'f', 10, 64)),
		}, &common.KeyValuePair{
			Key:   "orbitAngle",
			Value: []byte(strconv.FormatFloat(orbit.OrbitAngle, 'f', 10, 64)),
		}, &common.KeyValuePair{
			Key:   "ascendingNodeLongitude",
			Value: []byte(strconv.FormatFloat(orbit.AscendingNodeLongitude, 'f', 10, 64)),
		}, &common.KeyValuePair{
			Key:   "perigee",
			Value: []byte(strconv.FormatFloat(orbit.Perigee, 'f', 10, 64)),
		},
	)
	return params
}

func ConstellationConvert(conl *db.Constellation) []*common.KeyValuePair {
	params := make([]*common.KeyValuePair, 0)
	params = append(params,
		&common.KeyValuePair{
			Key:   "constellationId",
			Value: []byte(conl.ConstellationId),
		}, &common.KeyValuePair{
			Key:   "constellationName",
			Value: []byte(conl.ConstellationName),
		}, &common.KeyValuePair{
			Key:   "satelliteTotalNum",
			Value: []byte(strconv.Itoa(int(conl.SatelliteTotalNum))),
		}, &common.KeyValuePair{
			Key:   "satelliteUpNum",
			Value: []byte(strconv.Itoa(int(conl.SatelliteUpNum))),
		}, &common.KeyValuePair{
			Key:   "satelliteDownNum",
			Value: []byte(strconv.Itoa(int(conl.SatelliteDownNum))),
		}, &common.KeyValuePair{
			Key:   "satelliteLinkState",
			Value: []byte(db.StateName[conl.SatelliteLinkState]),
		},
	)
	return params
}

func LoginLogConvert(ll *db.LoginLog) []*common.KeyValuePair {
	params := make([]*common.KeyValuePair, 0)
	params = append(params,
		&common.KeyValuePair{
			Key:   "userName",
			Value: []byte(ll.UserName),
		}, &common.KeyValuePair{
			Key:   "loginTime",
			Value: []byte(strconv.Itoa(int(ll.LoginTime))),
		}, &common.KeyValuePair{
			Key:   "loginIp",
			Value: []byte(ll.LoginIp),
		},
	)
	return params
}

func OperationConvert(oper *db.Operation) []*common.KeyValuePair {
	params := make([]*common.KeyValuePair, 0)
	params = append(params,
		&common.KeyValuePair{
			Key:   "operator",
			Value: []byte(oper.Operator),
		}, &common.KeyValuePair{
			Key:   "operationTime",
			Value: []byte(strconv.Itoa(int(oper.OperationTime))),
		}, &common.KeyValuePair{
			Key:   "operationRecord",
			Value: []byte(oper.OperationRecord),
		}, &common.KeyValuePair{
			Key:   "satelliteName",
			Value: []byte(oper.SatelliteName),
		}, &common.KeyValuePair{
			Key:   "satelliteId",
			Value: []byte(oper.SatelliteId),
		},
	)
	return params
}

func SatelliteConvert(sate *db.Satellite) []*common.KeyValuePair {
	params := make([]*common.KeyValuePair, 0)
	params = append(params,
		&common.KeyValuePair{
			Key:   "satelliteId",
			Value: []byte(sate.SatelliteId),
		}, &common.KeyValuePair{
			Key:   "satelliteName",
			Value: []byte(sate.SatelliteName),
		}, &common.KeyValuePair{
			Key:   "orbitId",
			Value: []byte(sate.OrbitId),
		}, &common.KeyValuePair{
			Key:   "meanAnomaly",
			Value: []byte(strconv.FormatFloat(sate.MeanAnomaly, 'f', 10, 64)),
		}, &common.KeyValuePair{
			Key:   "speed",
			Value: []byte(strconv.FormatFloat(sate.Speed, 'f', 10, 64)),
		}, &common.KeyValuePair{
			Key:   "runState",
			Value: []byte(db.StateName[sate.RunState]),
		},
	)
	return params
}

func ControlConvert(ctrl *db.Control) []*common.KeyValuePair {
	params := make([]*common.KeyValuePair, 0)
	params = append(params,
		&common.KeyValuePair{
			Key:   "satelliteId",
			Value: []byte(ctrl.SatelliteId),
		}, &common.KeyValuePair{
			Key:   "satelliteName",
			Value: []byte(ctrl.SatelliteName),
		}, &common.KeyValuePair{
			Key:   "satelliteAttitude",
			Value: []byte(ctrl.SatelliteAttitude),
		}, &common.KeyValuePair{
			Key:   "satelliteTemperature",
			Value: []byte(strconv.FormatFloat(ctrl.SatelliteTemperature, 'f', 10, 64)),
		}, &common.KeyValuePair{
			Key:   "satellitePower",
			Value: []byte(ctrl.SatellitePower),
		},
	)
	return params
}

func FaultConvert(fault *db.Fault) []*common.KeyValuePair {
	params := make([]*common.KeyValuePair, 0)
	params = append(params,
		&common.KeyValuePair{
			Key:   "satelliteId",
			Value: []byte(fault.SatelliteId),
		}, &common.KeyValuePair{
			Key:   "satelliteName",
			Value: []byte(fault.SatelliteName),
		}, &common.KeyValuePair{
			Key:   "orbitId",
			Value: []byte(fault.OrbitId),
		}, &common.KeyValuePair{
			Key:   "faultType",
			Value: []byte(db.FaultTypeName[fault.FaultType]),
		}, &common.KeyValuePair{
			Key:   "faultTime",
			Value: []byte(strconv.Itoa(int(fault.FaultTime))),
		}, &common.KeyValuePair{
			Key:   "faultDescription",
			Value: []byte(fault.FaultDescription),
		}, &common.KeyValuePair{
			Key:   "repairState",
			Value: []byte(db.StateName[fault.RepairState]),
		},
	)
	return params
}

func CommStateConvert(cs *db.CommState) []*common.KeyValuePair {
	params := make([]*common.KeyValuePair, 0)
	params = append(params,
		&common.KeyValuePair{
			Key:   "satelliteId",
			Value: []byte(cs.SatelliteId),
		}, &common.KeyValuePair{
			Key:   "satelliteName",
			Value: []byte(cs.SatelliteName),
		}, &common.KeyValuePair{
			Key:   "orbitId",
			Value: []byte(cs.OrbitId),
		}, &common.KeyValuePair{
			Key:   "commState",
			Value: []byte(db.StateName[cs.CommState]),
		}, &common.KeyValuePair{
			Key:   "commPort",
			Value: []byte(cs.CommPort),
		}, &common.KeyValuePair{
			Key:   "commBandwidth",
			Value: []byte(cs.CommBandwidth),
		}, &common.KeyValuePair{
			Key:   "commDelay",
			Value: []byte(cs.CommDelay),
		}, &common.KeyValuePair{
			Key:   "linkLoad",
			Value: []byte(cs.LinkLoad),
		},
	)
	return params
}

func NetStateConvert(ns *db.NetState) []*common.KeyValuePair {
	params := make([]*common.KeyValuePair, 0)
	params = append(params,
		&common.KeyValuePair{
			Key:   "satelliteId",
			Value: []byte(ns.SatelliteId),
		}, &common.KeyValuePair{
			Key:   "satelliteName",
			Value: []byte(ns.SatelliteName),
		}, &common.KeyValuePair{
			Key:   "orbitId",
			Value: []byte(ns.OrbitId),
		}, &common.KeyValuePair{
			Key:   "networkSegment",
			Value: []byte(ns.NetworkSegment),
		}, &common.KeyValuePair{
			Key:   "networkState",
			Value: []byte(db.StateName[ns.NetworkState]),
		}, &common.KeyValuePair{
			Key:   "networkBandwidth",
			Value: []byte(ns.NetworkBandwidth),
		},
	)
	return params
}

func InstructionContractRespConvert(icp *models.InstructionContractResp, source, chainId string) []*db.Instruction {
	models := make([]*db.Instruction, 0)
	for _, v := range icp.Instructions {
		models = append(models, &db.Instruction{
			InstructionId:      v.InstructionId,
			InstructionContent: v.InstructionContent,
			Type:               db.AUTOMATIC,
			InstructionSource:  source,
			DebrisId:           v.DebrisId,
			DebrisName:         v.DebrisName,
			Treaten:            db.ThreatDegreeValue[v.Treaten],
			SatelliteId:        v.SatelliteId,
			SatelliteName:      v.SatelliteName,
			ExecState:          db.NOTEXEC,
			BlockChainField: db.BlockChainField{
				ChainId:     chainId,
				BlockHeight: icp.BlockHeight,
				TxId:        icp.TxId,
				ChainTime:   icp.ChainTime,
			},
		})
	}
	return models
}
