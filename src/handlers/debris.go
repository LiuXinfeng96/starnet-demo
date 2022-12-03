package handlers

import (
	"starnet-demo/src/db"
	"starnet-demo/src/models"
	"starnet-demo/src/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ControlAddDebris(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := checkTheAccessPermission(c, db.CONTROL); err != nil {
			WithoutPermissionJSONResp(err.Error(), c)
			return
		}

		var req models.AddDebirsReq
		if err := c.ShouldBindJSON(&req); err != nil {
			ParamsTypeErrorJSONResp(err.Error(), c)
			return
		}

		err := isStringRequiredParamsEmpty(req.DebrisId,
			req.DebrisName, req.Type)
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}

		debrisType, ok := db.DebrisTypeValue[req.Type]
		if !ok {
			ParamsValueJSONResp("debirs type not as expected", c)
			return
		}

		// 1. 碎片信息上主链
		// 2. 入库

		debirs := &db.Debris{
			DebrisId:   req.DebrisId,
			DebrisName: req.DebrisName,
			Angle:      req.Angle,
			Speed:      req.Speed,
			Height:     req.Height,
			Volunme:    req.Volume,
			Type:       debrisType,
		}

		err = s.InsertOneObjertToDB(debirs)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		SuccessfulJSONResp("", c)
	}
}

func ControlGetDebrisList(s *services.Server) gin.HandlerFunc {
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
			ModelStruct: new(db.Debris),
			Page:        int32(page),
			PageSize:    int32(pageSize),
			SortType:    sortType,
			SearchInput: searchInput,
			SearchIndex: make([]string, 0),
		}

		if len(searchInput) != 0 {
			params.SearchIndex = append(params.SearchIndex, "debris_id")
			params.SearchIndex = append(params.SearchIndex, "debris_name")
		}

		sqlRows, err := s.QueryObjectsWithPage(params)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		defer sqlRows.Close()

		resp := make([]*models.DebrisInfo, 0)

		for sqlRows.Next() {
			var debris db.Debris
			err := s.ScanRows(sqlRows, &debris)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.DebrisInfo{
				DebrisId:   debris.DebrisId,
				DebrisName: debris.DebrisName,
				Angle:      debris.Angle,
				Speed:      debris.Speed,
				Height:     debris.Height,
				Volume:     debris.Volunme,
				Type:       db.DebrisTypeName[debris.Type],
				BaseRespInfo: models.BaseRespInfo{
					Id:        debris.Id,
					LastTime:  debris.LastTime,
					ChainTime: debris.ChainTime,
				},
			})
		}

		SuccessfulJSONResp(resp, c)
	}
}

func TraceGetDebris(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := checkTheAccessPermission(c, db.TRACE); err != nil {
			WithoutPermissionJSONResp(err.Error(), c)
			return
		}
		debirsId := c.Query("debrisId")

		err := isStringRequiredParamsEmpty(debirsId)
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}

		model := new(db.Debris)

		sqlRows, err := s.QueryObjectsByCondition(model, "debirs_id", debirsId)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}
		defer sqlRows.Close()

		resp := make([]*models.DebrisHistoryInfo, 0)

		for sqlRows.Next() {
			var debris db.Debris
			err := s.ScanRows(sqlRows, &debris)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.DebrisHistoryInfo{
				DebrisInfo: models.DebrisInfo{
					DebrisId:   debris.DebrisId,
					DebrisName: debris.DebrisName,
					Angle:      debris.Angle,
					Speed:      debris.Speed,
					Height:     debris.Height,
					Volume:     debris.Volunme,
					Type:       db.DebrisTypeName[debris.Type],
					BaseRespInfo: models.BaseRespInfo{
						Id:        debris.Id,
						LastTime:  debris.LastTime,
						ChainTime: debris.ChainTime,
					},
				},
				HistoryRespInfo: models.HistoryRespInfo{
					ChainId:     debris.ChainId,
					BlockHeight: debris.BlockHeight,
					TxId:        debris.TxId,
				},
			})
		}
		SuccessfulJSONResp(resp, c)

	}
}

func ExecAddDebris(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := checkTheAccessPermission(c, db.EXEC); err != nil {
			WithoutPermissionJSONResp(err.Error(), c)
			return
		}

		var req models.AddDebirsReq
		if err := c.ShouldBindJSON(&req); err != nil {
			ParamsTypeErrorJSONResp(err.Error(), c)
			return
		}

		err := isStringRequiredParamsEmpty(req.DebrisId,
			req.DebrisName, req.Type)
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}

		debrisType, ok := db.DebrisTypeValue[req.Type]
		if !ok {
			ParamsValueJSONResp("debirs type not as expected", c)
			return
		}

		// 2. 自主规避
		// - 碎片信息录入
		// - 构造星座链交易，发送至星座链，等待上链成功
		// - 将返回的指令信息和碎片信息存至数据库
		// - 指令执行开始，更新星座链执行状态
		// - 将指令执行结果发送至星座链，等待上链成功
		// - 将返回的指令执行结果存至数据库
		// - 构造避障操作日志交易发送至星座链，上链成功后存至数据库

		debirs := &db.Debris{
			DebrisId:   req.DebrisId,
			DebrisName: req.DebrisName,
			Angle:      req.Angle,
			Speed:      req.Speed,
			Height:     req.Height,
			Volunme:    req.Volume,
			Type:       debrisType,
			BlockChainField: db.BlockChainField{
				ChainId: "testchain",
			},
		}

		err = s.InsertOneObjertToDB(debirs)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		SuccessfulJSONResp("", c)
	}
}

func ExecGetDebrisList(s *services.Server) gin.HandlerFunc {
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
			ModelStruct: new(db.Debris),
			Page:        int32(page),
			PageSize:    int32(pageSize),
			SortType:    sortType,
			SearchInput: searchInput,
			SearchIndex: make([]string, 0),
		}

		if len(searchInput) != 0 {
			params.SearchIndex = append(params.SearchIndex, "debris_id")
			params.SearchIndex = append(params.SearchIndex, "debris_name")
		}

		sqlRows, err := s.QueryObjectsWithPageSC(params)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		defer sqlRows.Close()

		resp := make([]*models.DebrisInfo, 0)

		for sqlRows.Next() {
			var debris db.Debris
			err := s.ScanRows(sqlRows, &debris)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.DebrisInfo{
				DebrisId:   debris.DebrisId,
				DebrisName: debris.DebrisName,
				Angle:      debris.Angle,
				Speed:      debris.Speed,
				Height:     debris.Height,
				Volume:     debris.Volunme,
				Type:       db.DebrisTypeName[debris.Type],
				BaseRespInfo: models.BaseRespInfo{
					Id:        debris.Id,
					LastTime:  debris.LastTime,
					ChainTime: debris.ChainTime,
				},
			})
		}

		SuccessfulJSONResp(resp, c)
	}
}
