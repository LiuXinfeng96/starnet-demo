package handlers

import (
	"starnet-demo/src/db"
	"starnet-demo/src/models"
	"starnet-demo/src/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ControlAddConstellation(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := checkTheAccessPermission(c, db.CONTROL); err != nil {
			WithoutPermissionJSONResp(err.Error(), c)
			return
		}

		var req models.AddConstellationReq
		if err := c.ShouldBindJSON(&req); err != nil {
			ParamsTypeErrorJSONResp(err.Error(), c)
			return
		}

		err := isStringRequiredParamsEmpty(req.ConstellationId, req.ConstellationName,
			req.SatelliteLinkState)
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}

		// 1. 星座信息上主链
		// 2. 入库

		constellation := &db.Constellation{
			ConstellationId:    req.ConstellationId,
			ConstellationName:  req.ConstellationName,
			SatelliteLinkState: req.SatelliteLinkState,
			SatelliteTotalNum:  req.SatelliteTotalNum,
			SatelliteUpNum:     req.SatelliteUpNum,
			SatelliteDownNum:   req.SatelliteDownNum,
		}

		err = s.InsertOneObjertToDB(constellation)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		SuccessfulJSONResp("", c)
	}
}

func ControlGetConstellationList(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := checkTheAccessPermission(c, db.CONTROL); err != nil {
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
			ModelStruct: new(db.Constellation),
			Page:        int32(page),
			PageSize:    int32(pageSize),
			SortType:    sortType,
			SearchInput: searchInput,
			SearchIndex: make([]string, 0),
		}

		if len(searchInput) != 0 {
			params.SearchIndex = append(params.SearchIndex, "constellation_id")
			params.SearchIndex = append(params.SearchIndex, "constellation_name")
		}

		sqlRows, err := s.QueryObjectsWithPage(params)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		defer sqlRows.Close()

		resp := make([]*models.ConstellationInfo, 0)

		for sqlRows.Next() {
			var constellation db.Constellation
			err := s.ScanRows(sqlRows, &constellation)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.ConstellationInfo{
				ConstellationId:    constellation.ConstellationId,
				ConstellationName:  constellation.ConstellationName,
				SatelliteLinkState: constellation.SatelliteLinkState,
				SatelliteTotalNum:  constellation.SatelliteTotalNum,
				SatelliteUpNum:     constellation.SatelliteUpNum,
				SatelliteDownNum:   constellation.SatelliteDownNum,
				BaseRespInfo: models.BaseRespInfo{
					Id:        constellation.Id,
					LastTime:  constellation.LastTime,
					ChainTime: constellation.ChainTime,
				},
			})
		}

		SuccessfulJSONResp(resp, c)
	}
}

func TraceGetConstellation(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := checkTheAccessPermission(c, db.TRACE); err != nil {
			WithoutPermissionJSONResp(err.Error(), c)
			return
		}
		constellationId := c.Query("constellationId")

		err := isStringRequiredParamsEmpty(constellationId)
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}

		model := new(db.Constellation)

		sqlRows, err := s.QueryObjectsByCondition(model, "constellation_id", constellationId)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}
		defer sqlRows.Close()

		resp := make([]*models.ConstellationHistoryInfo, 0)

		for sqlRows.Next() {
			var constellation db.Constellation
			err := s.ScanRows(sqlRows, &constellation)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.ConstellationHistoryInfo{
				ConstellationInfo: models.ConstellationInfo{
					ConstellationId:    constellation.ConstellationId,
					ConstellationName:  constellation.ConstellationName,
					SatelliteLinkState: constellation.SatelliteLinkState,
					SatelliteTotalNum:  constellation.SatelliteTotalNum,
					SatelliteUpNum:     constellation.SatelliteUpNum,
					SatelliteDownNum:   constellation.SatelliteDownNum,
					BaseRespInfo: models.BaseRespInfo{
						Id:        constellation.Id,
						LastTime:  constellation.LastTime,
						ChainTime: constellation.ChainTime,
					},
				},

				HistoryRespInfo: models.HistoryRespInfo{
					ChainId:     constellation.ChainId,
					BlockHeight: constellation.BlockHeight,
					TxId:        constellation.TxId,
				},
			})
		}
		SuccessfulJSONResp(resp, c)

	}
}
