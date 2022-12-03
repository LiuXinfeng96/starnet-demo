package handlers

import (
	"starnet-demo/src/db"
	"starnet-demo/src/models"
	"starnet-demo/src/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ExecAddNetState(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := checkTheAccessPermission(c, db.EXEC); err != nil {
			WithoutPermissionJSONResp(err.Error(), c)
			return
		}

		var req models.AddNetStateReq
		if err := c.ShouldBindJSON(&req); err != nil {
			ParamsTypeErrorJSONResp(err.Error(), c)
			return
		}

		err := isStringRequiredParamsEmpty(req.SatelliteId, req.SatelliteName,
			req.OrbitId, req.NetworkSegment, req.NetworkState, req.NetworkBandwidth)
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}

		// 1. 故障信息上星座链
		// 2. 入库

		netState := &db.NetState{
			SatelliteId:      req.SatelliteId,
			SatelliteName:    req.SatelliteName,
			OrbitId:          req.OrbitId,
			NetworkSegment:   req.NetworkSegment,
			NetworkState:     req.NetworkState,
			NetworkBandwidth: req.NetworkBandwidth,
		}

		err = s.InsertOneObjertToDB(netState)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		SuccessfulJSONResp("", c)
	}
}

func ExecGetNetStateList(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := checkTheAccessPermission(c, db.EXEC); err != nil {
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
			ModelStruct: new(db.NetState),
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

		resp := make([]*models.NetStateInfo, 0)

		for sqlRows.Next() {
			var netState db.NetState
			err := s.ScanRows(sqlRows, &netState)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.NetStateInfo{
				SatelliteId:      netState.SatelliteId,
				SatelliteName:    netState.SatelliteName,
				OrbitId:          netState.OrbitId,
				NetworkSegment:   netState.NetworkSegment,
				NetworkState:     netState.NetworkState,
				NetworkBandwidth: netState.NetworkBandwidth,
				BaseRespInfo: models.BaseRespInfo{
					Id:        netState.Id,
					LastTime:  netState.LastTime,
					ChainTime: netState.ChainTime,
				},
			})
		}

		SuccessfulJSONResp(resp, c)
	}
}

func TraceGetNetState(s *services.Server) gin.HandlerFunc {
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

		model := new(db.NetState)

		sqlRows, err := s.QueryObjectsByCondition(model, "satellite_id", satelliteId)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}
		defer sqlRows.Close()

		resp := make([]*models.NetStateHistoryInfo, 0)

		for sqlRows.Next() {
			var netState db.NetState
			err := s.ScanRows(sqlRows, &netState)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.NetStateHistoryInfo{
				NetStateInfo: models.NetStateInfo{
					SatelliteId:      netState.SatelliteId,
					SatelliteName:    netState.SatelliteName,
					OrbitId:          netState.OrbitId,
					NetworkSegment:   netState.NetworkSegment,
					NetworkState:     netState.NetworkState,
					NetworkBandwidth: netState.NetworkBandwidth,
					BaseRespInfo: models.BaseRespInfo{
						Id:        netState.Id,
						LastTime:  netState.LastTime,
						ChainTime: netState.ChainTime,
					},
				},

				HistoryRespInfo: models.HistoryRespInfo{
					ChainId:     netState.ChainId,
					BlockHeight: netState.BlockHeight,
					TxId:        netState.TxId,
				},
			})
		}

		SuccessfulJSONResp(resp, c)
	}
}
