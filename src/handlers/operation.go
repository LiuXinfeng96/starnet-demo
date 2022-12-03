package handlers

import (
	"starnet-demo/src/db"
	"starnet-demo/src/models"
	"starnet-demo/src/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ControlGetOperationList(s *services.Server) gin.HandlerFunc {
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
			ModelStruct: new(db.Operation),
			Page:        int32(page),
			PageSize:    int32(pageSize),
			SortType:    sortType,
			SearchInput: searchInput,
			SearchIndex: make([]string, 0),
		}

		if len(searchInput) != 0 {
			params.SearchIndex = append(params.SearchIndex, "operator")
			params.SearchIndex = append(params.SearchIndex, "satellite_id")
			params.SearchIndex = append(params.SearchIndex, "satellite_name")
		}

		sqlRows, err := s.QueryObjectsWithPage(params)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		defer sqlRows.Close()

		resp := make([]*models.OperationInfo, 0)

		for sqlRows.Next() {
			var operation db.Operation
			err := s.ScanRows(sqlRows, &operation)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.OperationInfo{
				Operator:        operation.Operator,
				OperationTime:   operation.OperationTime,
				OperationIp:     operation.OperatorIp,
				OperationRecord: operation.OperationRecord,
				SatelliteId:     operation.SatelliteId,
				SatelliteName:   operation.SatelliteName,
				BaseRespInfo: models.BaseRespInfo{
					Id:        operation.Id,
					LastTime:  operation.LastTime,
					ChainTime: operation.ChainTime,
				},
			})
		}

		SuccessfulJSONResp(resp, c)
	}
}

func TraceGetOperation(s *services.Server) gin.HandlerFunc {
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

		model := new(db.Operation)

		sqlRows, err := s.QueryObjectsByCondition(model, "satellite_id", satelliteId)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}
		defer sqlRows.Close()

		resp := make([]*models.OperationHistoryInfo, 0)

		for sqlRows.Next() {
			var operation db.Operation
			err := s.ScanRows(sqlRows, &operation)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.OperationHistoryInfo{
				OperationInfo: models.OperationInfo{
					Operator:        operation.Operator,
					OperationTime:   operation.OperationTime,
					OperationIp:     operation.OperatorIp,
					OperationRecord: operation.OperationRecord,
					SatelliteId:     operation.SatelliteId,
					SatelliteName:   operation.SatelliteName,
					BaseRespInfo: models.BaseRespInfo{
						Id:        operation.Id,
						LastTime:  operation.LastTime,
						ChainTime: operation.ChainTime,
					},
				},

				HistoryRespInfo: models.HistoryRespInfo{
					ChainId:     operation.ChainId,
					BlockHeight: operation.BlockHeight,
					TxId:        operation.TxId,
				},
			})
		}

		SuccessfulJSONResp(resp, c)
	}
}

func ExecGetOperationList(s *services.Server) gin.HandlerFunc {
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
			ModelStruct: new(db.Operation),
			Page:        int32(page),
			PageSize:    int32(pageSize),
			SortType:    sortType,
			SearchInput: searchInput,
			SearchIndex: make([]string, 0),
		}

		if len(searchInput) != 0 {
			params.SearchIndex = append(params.SearchIndex, "operator")
			params.SearchIndex = append(params.SearchIndex, "satellite_id")
			params.SearchIndex = append(params.SearchIndex, "satellite_name")
		}

		sqlRows, err := s.QueryObjectsWithPageSC(params)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		defer sqlRows.Close()

		resp := make([]*models.OperationInfo, 0)

		for sqlRows.Next() {
			var operation db.Operation
			err := s.ScanRows(sqlRows, &operation)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.OperationInfo{
				Operator:        operation.Operator,
				OperationTime:   operation.OperationTime,
				OperationIp:     operation.OperatorIp,
				OperationRecord: operation.OperationRecord,
				SatelliteId:     operation.SatelliteId,
				SatelliteName:   operation.SatelliteName,
				BaseRespInfo: models.BaseRespInfo{
					Id:        operation.Id,
					LastTime:  operation.LastTime,
					ChainTime: operation.ChainTime,
				},
			})
		}

		SuccessfulJSONResp(resp, c)
	}
}
