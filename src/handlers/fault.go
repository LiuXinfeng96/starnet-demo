package handlers

import (
	"starnet-demo/src/db"
	"starnet-demo/src/models"
	"starnet-demo/src/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ExecAddFault(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := checkTheAccessPermission(c, db.EXEC); err != nil {
			WithoutPermissionJSONResp(err.Error(), c)
			return
		}

		var req models.AddFaultReq
		if err := c.ShouldBindJSON(&req); err != nil {
			ParamsTypeErrorJSONResp(err.Error(), c)
			return
		}

		err := isStringRequiredParamsEmpty(req.SatelliteId, req.SatelliteName,
			req.OrbitId, req.FaultType, req.FaultDescription, req.RepairState)
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}

		// 1. 故障信息上星座链
		// 2. 入库

		fault := &db.Fault{
			SatelliteId:      req.SatelliteId,
			SatelliteName:    req.SatelliteName,
			OrbitId:          req.OrbitId,
			FaultType:        req.FaultType,
			FaultDescription: req.FaultDescription,
			FaultTime:        req.FaultTime,
			RepairState:      req.RepairState,
		}

		err = s.InsertOneObjertToDB(fault)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		SuccessfulJSONResp("", c)
	}
}

func ExecGetFaultList(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := checkTheAccessPermission(c, db.EXEC, db.TRACE); err != nil {
			WithoutPermissionJSONResp(err.Error(), c)
			return
		}

		pageStr := c.Query("page")
		pageSizeStr := c.Query("pageSize")
		sortTypeStr := c.Query("sortType")
		searchInput := c.Query("searchConditions")

		err := isStringRequiredParamsEmpty(pageSizeStr, pageStr)
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}

		page, err := strconv.Atoi(pageStr)
		if err != nil {
			ParamsTypeErrorJSONResp(err.Error(), c)
			return
		}

		pageSize, err := strconv.Atoi(pageSizeStr)
		if err != nil {
			ParamsTypeErrorJSONResp(err.Error(), c)
			return
		}

		sortType, ok := services.SortTypeValue[sortTypeStr]
		if !ok {
			sortType = services.SORTTYPE_TIME
		}

		params := &services.QueryObjectsParams{
			ModelStruct: new(db.Fault),
			Page:        int32(page),
			PageSize:    int32(pageSize),
			SortType:    sortType,
			SearchInput: searchInput,
			SearchIndex: make([]string, 0),
		}

		if len(searchInput) != 0 {
			params.SearchIndex = append(params.SearchIndex, "satellite_id")
			params.SearchIndex = append(params.SearchIndex, "satellite_name")
		}

		sqlRows, err := s.QueryObjectsWithPage(params)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		defer sqlRows.Close()

		resp := make([]*models.FaultInfo, 0)

		for sqlRows.Next() {
			var fault db.Fault
			err := s.ScanRows(sqlRows, &fault)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.FaultInfo{
				SatelliteId:      fault.SatelliteId,
				SatelliteName:    fault.SatelliteName,
				OrbitId:          fault.OrbitId,
				FaultType:        fault.FaultType,
				FaultDescription: fault.FaultDescription,
				FaultTime:        fault.FaultTime,
				RepairState:      fault.RepairState,
				BaseRespInfo: models.BaseRespInfo{
					Id:        fault.Id,
					LastTime:  fault.LastTime,
					ChainTime: fault.ChainTime,
				},
			})
		}

		SuccessfulJSONRespWithPage(resp, len(resp), c)
	}
}

func TraceGetFault(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := checkTheAccessPermission(c, db.TRACE); err != nil {
			WithoutPermissionJSONResp(err.Error(), c)
			return
		}
		satelliteId := c.Query("satelliteId")

		err := isStringRequiredParamsEmpty(satelliteId)
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}

		model := new(db.Fault)

		sqlRows, err := s.QueryObjectsByCondition(model, "satellite_id", satelliteId)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}
		defer sqlRows.Close()

		resp := make([]*models.FaultHistoryInfo, 0)

		for sqlRows.Next() {
			var fault db.Fault
			err := s.ScanRows(sqlRows, &fault)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.FaultHistoryInfo{
				FaultInfo: models.FaultInfo{
					SatelliteId:      fault.SatelliteId,
					SatelliteName:    fault.SatelliteName,
					OrbitId:          fault.OrbitId,
					FaultType:        fault.FaultType,
					FaultDescription: fault.FaultDescription,
					FaultTime:        fault.FaultTime,
					RepairState:      fault.RepairState,
					BaseRespInfo: models.BaseRespInfo{
						Id:        fault.Id,
						LastTime:  fault.LastTime,
						ChainTime: fault.ChainTime,
					},
				},

				HistoryRespInfo: models.HistoryRespInfo{
					ChainId:     fault.ChainId,
					BlockHeight: fault.BlockHeight,
					TxId:        fault.TxId,
				},
			})
		}

		SuccessfulJSONRespWithPage(resp, len(resp), c)
	}
}

func TraceGetFaultList(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := checkTheAccessPermission(c, db.TRACE); err != nil {
			WithoutPermissionJSONResp(err.Error(), c)
			return
		}

		pageStr := c.Query("page")
		pageSizeStr := c.Query("pageSize")
		sortTypeStr := c.Query("sortType")
		searchInput := c.Query("searchConditions")

		err := isStringRequiredParamsEmpty(pageSizeStr, pageStr)
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}

		page, err := strconv.Atoi(pageStr)
		if err != nil {
			ParamsTypeErrorJSONResp(err.Error(), c)
			return
		}

		pageSize, err := strconv.Atoi(pageSizeStr)
		if err != nil {
			ParamsTypeErrorJSONResp(err.Error(), c)
			return
		}

		sortType, ok := services.SortTypeValue[sortTypeStr]
		if !ok {
			sortType = services.SORTTYPE_TIME
		}

		params := &services.QueryLatestObjectsParams{
			ModelStruct: new(db.Fault),
			Page:        int32(page),
			PageSize:    int32(pageSize),
			SortType:    sortType,
			SearchInput: searchInput,
			SearchIndex: make([]string, 0),
			GroupIndex:  "satellite_id",
		}

		if len(searchInput) != 0 {
			params.SearchIndex = append(params.SearchIndex, "satellite_id")
			params.SearchIndex = append(params.SearchIndex, "satellite_name")
		}

		sqlRows, err := s.QueryLatestObjectsWithPage(params)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		defer sqlRows.Close()

		resp := make([]*models.FaultInfo, 0)

		for sqlRows.Next() {
			var fault db.Fault
			err := s.ScanRows(sqlRows, &fault)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.FaultInfo{
				SatelliteId:      fault.SatelliteId,
				SatelliteName:    fault.SatelliteName,
				OrbitId:          fault.OrbitId,
				FaultType:        fault.FaultType,
				FaultDescription: fault.FaultDescription,
				FaultTime:        fault.FaultTime,
				RepairState:      fault.RepairState,
				BaseRespInfo: models.BaseRespInfo{
					Id:        fault.Id,
					LastTime:  fault.LastTime,
					ChainTime: fault.ChainTime,
				},
			})
		}

		SuccessfulJSONRespWithPage(resp, len(resp), c)
	}
}
