package handlers

import (
	"starnet-demo/src/db"
	"starnet-demo/src/models"
	"starnet-demo/src/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ExecAddControl(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := checkTheAccessPermission(c, db.EXEC); err != nil {
			WithoutPermissionJSONResp(err.Error(), c)
			return
		}

		var req models.AddControlsReq
		if err := c.ShouldBindJSON(&req); err != nil {
			ParamsTypeErrorJSONResp(err.Error(), c)
			return
		}

		err := isStringRequiredParamsEmpty(req.SatelliteId, req.SatelliteName,
			req.SatelliteAttitude, req.SatellitePower)
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}

		// 1. 管控信息上星座链
		// 2. 入库

		control := &db.Control{
			SatelliteId:          req.SatelliteId,
			SatelliteName:        req.SatelliteName,
			SatelliteAttitude:    req.SatelliteAttitude,
			SatelliteTemperature: req.SatelliteTemperature,
			SatellitePower:       req.SatellitePower,
		}

		err = s.InsertOneObjertToDB(control)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		SuccessfulJSONResp("", c)
	}
}

func ExecGetControlList(s *services.Server) gin.HandlerFunc {
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
			ModelStruct: new(db.Control),
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

		resp := make([]*models.ControlsInfo, 0)

		for sqlRows.Next() {
			var control db.Control
			err := s.ScanRows(sqlRows, &control)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.ControlsInfo{
				SatelliteId:          control.SatelliteId,
				SatelliteName:        control.SatelliteName,
				SatelliteAttitude:    control.SatelliteAttitude,
				SatelliteTemperature: control.SatelliteTemperature,
				SatellitePower:       control.SatellitePower,
				BaseRespInfo: models.BaseRespInfo{
					Id:        control.Id,
					LastTime:  control.LastTime,
					ChainTime: control.ChainTime,
				},
			})
		}

		SuccessfulJSONRespWithPage(resp, len(resp), c)
	}
}

func TraceGetControl(s *services.Server) gin.HandlerFunc {
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

		model := new(db.Control)

		sqlRows, err := s.QueryObjectsByCondition(model, "satellite_id", satelliteId)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}
		defer sqlRows.Close()

		resp := make([]*models.ControlsHistoryInfo, 0)

		for sqlRows.Next() {
			var control db.Control
			err := s.ScanRows(sqlRows, &control)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.ControlsHistoryInfo{
				ControlsInfo: models.ControlsInfo{
					SatelliteId:          control.SatelliteId,
					SatelliteName:        control.SatelliteName,
					SatelliteAttitude:    control.SatelliteAttitude,
					SatelliteTemperature: control.SatelliteTemperature,
					SatellitePower:       control.SatellitePower,
					BaseRespInfo: models.BaseRespInfo{
						Id:        control.Id,
						LastTime:  control.LastTime,
						ChainTime: control.ChainTime,
					},
				},

				HistoryRespInfo: models.HistoryRespInfo{
					ChainId:     control.ChainId,
					BlockHeight: control.BlockHeight,
					TxId:        control.TxId,
				},
			})
		}

		SuccessfulJSONRespWithPage(resp, len(resp), c)
	}
}
