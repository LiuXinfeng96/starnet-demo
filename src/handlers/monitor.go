package handlers

import (
	"starnet-demo/src/db"
	"starnet-demo/src/models"
	"starnet-demo/src/services"

	"github.com/gin-gonic/gin"
)

func MonitorGetAllState(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 查询所有卫星最新的故障状态是否异常
		// 查询所有卫星最新的在轨状态是否异常
		// 查询所有星座最新的状态是否异常
		// 查询所有卫星最新的网络状态是否异常
		// 查询所有卫星最新的通信状态是否异常

		var resp models.AllState

		fault := new(db.Fault)
		sqlRows, err := s.QueryOneObjectWithLatest(fault, "satellite_id")
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		defer sqlRows.Close()

		for sqlRows.Next() {
			var f db.Fault
			err := s.ScanRows(sqlRows, &f)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			if f.RepairState == db.WRONG {
				resp.FaultState = true
				break
			}
		}

		satellite := new(db.Satellite)
		sqlRows, err = s.QueryOneObjectWithLatest(satellite, "satellite_id")
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		defer sqlRows.Close()

		for sqlRows.Next() {
			var sl db.Satellite
			err := s.ScanRows(sqlRows, &sl)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			if sl.RunState == db.WRONG {
				resp.SatelliteState = true
				break
			}
		}

		constellation := new(db.Constellation)
		sqlRows, err = s.QueryOneObjectWithLatest(constellation, "satellite_id")
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		defer sqlRows.Close()

		for sqlRows.Next() {
			var conl db.Constellation
			err := s.ScanRows(sqlRows, &conl)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			if conl.SatelliteLinkState == db.WRONG {
				resp.ConstellationState = true
				break
			}
		}

		netState := new(db.NetState)
		sqlRows, err = s.QueryOneObjectWithLatest(netState, "satellite_id")
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		defer sqlRows.Close()

		for sqlRows.Next() {
			var ns db.NetState
			err := s.ScanRows(sqlRows, &ns)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			if ns.NetworkState == db.WRONG {
				resp.NetState = true
				break
			}
		}

		commState := new(db.CommState)
		sqlRows, err = s.QueryOneObjectWithLatest(commState, "satellite_id")
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		defer sqlRows.Close()

		for sqlRows.Next() {
			var cs db.CommState
			err := s.ScanRows(sqlRows, &cs)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			if cs.CommState == db.WRONG {
				resp.CommState = true
				break
			}
		}

		SuccessfulJSONResp(&resp, c)
	}
}

func MonitorGetTableCount(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {

		var resp models.ChainDataNum
		err := s.GetTableDataCount(&db.Debris{}, &resp.DebrisNum)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		err = s.GetTableDataCount(&db.Instruction{}, &resp.InstructionNum)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		err = s.GetTableDataCount(&db.Satellite{}, &resp.SatelliteStateNum)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		err = s.GetTableDataCount(&db.Orbit{}, &resp.OrbitNum)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		err = s.GetTableDataCount(&db.Constellation{}, &resp.ConstellationNum)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		err = s.GetTableDataCount(&db.Control{}, &resp.ControlNum)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		err = s.GetTableDataCount(&db.Fault{}, &resp.FaultNum)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		err = s.GetTableDataCount(&db.NetState{}, &resp.NetStateNum)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		err = s.GetTableDataCount(&db.CommState{}, &resp.CommStateNum)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		err = s.GetTableDataCount(&db.Operation{}, &resp.OperationNum)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		err = s.GetTableDataCount(&db.LoginLog{}, &resp.LoginLogNum)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		SuccessfulJSONResp(&resp, c)
	}

}

func MonitorGetFaultInfo(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		var resp models.FaultTypeInfo

		gormDb := s.GetGormObject()

		err := gormDb.Model(&db.Fault{}).Where("fault_type = ?", db.AType).
			Count(&resp.AType).Error
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		err = gormDb.Model(&db.Fault{}).Where("fault_type = ?", db.BType).
			Count(&resp.BType).Error
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		err = gormDb.Model(&db.Fault{}).Where("fault_type = ?", db.CType).
			Count(&resp.CType).Error
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		err = gormDb.Model(&db.Fault{}).Where("fault_type = ?", db.DType).
			Count(&resp.DType).Error
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		err = gormDb.Model(&db.Fault{}).Where("fault_type = ?", db.EType).
			Count(&resp.EType).Error
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		SuccessfulJSONResp(&resp, c)
	}
}

func MonitorGetEarlWarning(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		var resp []models.EarlyWarningInfo

		gormDB := s.GetGormObject()
		err := gormDB.Model(&db.Instruction{}).
			Select("instruction.debris_id, instruction.debris_name, instruction.debris_name.threaten, "+
				"instruction.satellite_name, instruction.satellite_id, debris.speed, debris.height").
			Joins("inner join debris on debris.debris_id = instruction.debris_id").
			Where("treaten = ? AND treaten = ? AND exec_state = ?", db.LOW, db.HIGH, db.NOTEXEC).
			Order("last_time desc").Limit(20).Find(&resp).Error
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		SuccessfulJSONResp(resp, c)

	}
}
