package handlers

import (
	"starnet-demo/src/contract"
	"starnet-demo/src/db"
	"starnet-demo/src/models"
	"starnet-demo/src/services"
	"strconv"

	"chainmaker.org/chainmaker/pb-go/v2/common"
	"github.com/gin-gonic/gin"
)

func ExecAddSatelliteState(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := checkTheAccessPermission(c, db.EXEC); err != nil {
			WithoutPermissionJSONResp(err.Error(), c)
			return
		}

		var req models.AddSatelliteState
		if err := c.ShouldBindJSON(&req); err != nil {
			ParamsTypeErrorJSONResp(err.Error(), c)
			return
		}

		err := isStringRequiredParamsEmpty(req.SatelliteId, req.SatelliteName,
			req.OrbitId, req.RunState)
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}

		// 1. 卫星在轨状态上星座链
		// 2. 入库

		runState, ok := db.SatelliteRunStateValue[req.RunState]
		if !ok {
			ParamsValueJSONResp("run state type not as expected", c)
			return
		}
		satellite := &db.Satellite{
			SatelliteId:   req.SatelliteId,
			SatelliteName: req.SatelliteName,
			OrbitId:       req.OrbitId,
			RunState:      runState,
			MeanAnomaly:   req.MeanAnomaly,
			Speed:         req.Speed,
			BlockChainField: db.BlockChainField{
				ChainId: s.GetExecChainId(),
			},
		}

		token, ok1 := c.Get("token")
		claims, ok2 := token.(*services.MyClaims)
		if !ok1 || !ok2 {
			ServerErrorJSONResp("get the token from context failed", c)
			return
		}
		client, err := s.GetSdkClient(claims.Name + s.GetExecChainId())
		if err != nil {
			NotInChainJSONResp(err.Error(), c)
			return
		}

		kvs := contract.SatelliteConvert(satellite)

		chainResp, err := client.InvokeContract(s.GetExecContractName(),
			contract.EXEC_CONTRACT_FUNC_NAME_PUT_SATELLITE, "", kvs, -1, true)
		if err != nil {
			PutChainFailJSONResp(err.Error(), c)
			return
		}
		if chainResp.Code != common.TxStatusCode_SUCCESS {
			PutChainFailJSONResp(chainResp.Message, c)
			return
		}

		satellite.BlockChainField, err = GetBlockChainFiledFromResp(chainResp.ContractResult)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		go func() {
			masterClient, err := s.GetSdkClient(claims.Name + s.GetMasterChainId())
			if err != nil {
				s.GetSuLogger().Warn(err)
				return
			}
			resp, err := masterClient.InvokeContract(s.GetMasterContractName(),
				contract.MASTER_CONTRACT_FUNC_NAME_PUT_SATELLITE, "", kvs, -1, true)
			if err != nil {
				s.GetSuLogger().Warn(err)
				return
			}
			if resp.Code != common.TxStatusCode_SUCCESS {
				s.GetSuLogger().Warnf("invoke contract failed: tx status code: [%d], msg: [%s]\n",
					resp.Code, resp.Message)
				return
			}
		}()

		err = s.InsertOneObjertToDB(satellite)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		SuccessfulJSONResp("", c)
	}
}

func ExecGetSatelliteStateList(s *services.Server) gin.HandlerFunc {
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
			ModelStruct: new(db.Satellite),
			Page:        int32(page),
			PageSize:    int32(pageSize),
			SortType:    sortType,
			SearchInput: searchInput,
			SearchIndex: make([]string, 0),
		}

		if len(searchInput) != 0 {
			params.SearchIndex = append(params.SearchIndex, "satellite_id")
			params.SearchIndex = append(params.SearchIndex, "satellite_name")
			params.SearchIndex = append(params.SearchIndex, "orbit_id")
		}

		sqlRows, err := s.QueryObjectsWithPage(params)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		defer sqlRows.Close()

		resp := make([]*models.SatelliteStateInfo, 0)

		for sqlRows.Next() {
			var satellite db.Satellite
			err := s.ScanRows(sqlRows, &satellite)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.SatelliteStateInfo{
				SatelliteId:   satellite.SatelliteId,
				SatelliteName: satellite.SatelliteName,
				OrbitId:       satellite.OrbitId,
				RunState:      db.SatelliteRunStateName[satellite.RunState],
				MeanAnomaly:   satellite.MeanAnomaly,
				Speed:         satellite.Speed,
				BaseRespInfo: models.BaseRespInfo{
					Id:        satellite.Id,
					LastTime:  satellite.LastTime,
					ChainTime: satellite.ChainTime,
				},
			})
		}

		SuccessfulJSONRespWithPage(resp, len(resp), c)
	}
}

func TraceGetSatelliteState(s *services.Server) gin.HandlerFunc {
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

		model := new(db.Satellite)

		sqlRows, err := s.QueryObjectsByCondition(model, "satellite_id", satelliteId)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}
		defer sqlRows.Close()

		resp := make([]*models.SatelliteStateHistoryInfo, 0)

		for sqlRows.Next() {
			var satellite db.Satellite
			err := s.ScanRows(sqlRows, &satellite)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.SatelliteStateHistoryInfo{
				SatelliteStateInfo: models.SatelliteStateInfo{
					SatelliteId:   satellite.SatelliteId,
					SatelliteName: satellite.SatelliteName,
					OrbitId:       satellite.OrbitId,
					RunState:      db.SatelliteRunStateName[satellite.RunState],
					MeanAnomaly:   satellite.MeanAnomaly,
					Speed:         satellite.Speed,
					BaseRespInfo: models.BaseRespInfo{
						Id:        satellite.Id,
						LastTime:  satellite.LastTime,
						ChainTime: satellite.ChainTime,
					},
				},

				HistoryRespInfo: models.HistoryRespInfo{
					ChainId:     satellite.ChainId,
					BlockHeight: satellite.BlockHeight,
					TxId:        satellite.TxId,
				},
			})
		}

		SuccessfulJSONRespWithPage(resp, len(resp), c)
	}
}

func TraceGetSatelliteStateList(s *services.Server) gin.HandlerFunc {
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
			ModelStruct: new(db.Satellite),
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
			params.SearchIndex = append(params.SearchIndex, "orbit_id")
		}

		sqlRows, err := s.QueryLatestObjectsWithPage(params)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		defer sqlRows.Close()

		resp := make([]*models.SatelliteStateInfo, 0)

		for sqlRows.Next() {
			var satellite db.Satellite
			err := s.ScanRows(sqlRows, &satellite)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.SatelliteStateInfo{
				SatelliteId:   satellite.SatelliteId,
				SatelliteName: satellite.SatelliteName,
				OrbitId:       satellite.OrbitId,
				RunState:      db.SatelliteRunStateName[satellite.RunState],
				MeanAnomaly:   satellite.MeanAnomaly,
				Speed:         satellite.Speed,
				BaseRespInfo: models.BaseRespInfo{
					Id:        satellite.Id,
					LastTime:  satellite.LastTime,
					ChainTime: satellite.ChainTime,
				},
			})
		}

		SuccessfulJSONRespWithPage(resp, len(resp), c)
	}
}
