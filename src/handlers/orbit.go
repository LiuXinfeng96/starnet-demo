package handlers

import (
	"starnet-demo/src/db"
	"starnet-demo/src/models"
	"starnet-demo/src/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ControlAddOrbit(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := checkTheAccessPermission(c, db.CONTROL); err != nil {
			WithoutPermissionJSONResp(err.Error(), c)
			return
		}

		var req models.AddOrbitReq
		if err := c.ShouldBindJSON(&req); err != nil {
			ParamsTypeErrorJSONResp(err.Error(), c)
			return
		}

		err := isStringRequiredParamsEmpty(req.OrbitId, req.OrbitType)
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}

		// 1. 轨道信息上主链
		// 2. 入库

		orbit := &db.Orbit{
			OrbitId:                req.OrbitId,
			OrbitType:              req.OrbitType,
			OrbitSemiMajorAxis:     req.OrbitSemiMajorAxis,
			OrbitEccentricity:      req.OrbitEccentricity,
			OrbitAngle:             req.OrbitAngle,
			AscendingNodeLongitude: req.AscendingNodeLongitude,
			Perigee:                req.Perigee,
		}

		err = s.InsertOneObjertToDB(orbit)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		SuccessfulJSONResp("", c)
	}
}

func ControlGetOrbitList(s *services.Server) gin.HandlerFunc {
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
			ModelStruct: new(db.Orbit),
			Page:        int32(page),
			PageSize:    int32(pageSize),
			SortType:    sortType,
			SearchInput: searchInput,
			SearchIndex: make([]string, 0),
		}

		if len(searchInput) != 0 {
			params.SearchIndex = append(params.SearchIndex, "orbit_id")
		}

		sqlRows, err := s.QueryObjectsWithPage(params)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		defer sqlRows.Close()

		resp := make([]*models.OrbitInfo, 0)

		for sqlRows.Next() {
			var orbit db.Orbit
			err := s.ScanRows(sqlRows, &orbit)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.OrbitInfo{
				OrbitId:                orbit.OrbitId,
				OrbitType:              orbit.OrbitType,
				OrbitEccentricity:      orbit.OrbitEccentricity,
				OrbitSemiMajorAxis:     orbit.OrbitSemiMajorAxis,
				OrbitAngle:             orbit.OrbitAngle,
				AscendingNodeLongitude: orbit.AscendingNodeLongitude,
				Perigee:                orbit.Perigee,
				BaseRespInfo: models.BaseRespInfo{
					Id:        orbit.Id,
					LastTime:  orbit.LastTime,
					ChainTime: orbit.ChainTime,
				},
			})
		}

		SuccessfulJSONResp(resp, c)
	}
}
