package handlers

import (
	"starnet-demo/src/db"
	"starnet-demo/src/models"
	"starnet-demo/src/services"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func AddDebris(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := checkTheAccessPermission(c, db.CONTROL); err != nil {
			WithoutPermissionJSONResp(err.Error(), c)
			return
		}

		var addDebirsReq models.AddDebirsReq
		if err := c.ShouldBindBodyWith(&addDebirsReq, binding.JSON); err != nil {
			ParamsTypeErrorJSONResp(err.Error(), c)
			return
		}

		err := isStringRequiredParamsEmpty(addDebirsReq.DebrisId,
			addDebirsReq.DebrisName, addDebirsReq.Type)
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}

		debrisType, ok := db.DebrisTypeValue[addDebirsReq.Type]
		if !ok {
			ParamsValueJSONResp("debirs type not as expected", c)
			return
		}
		debirs := &db.Debris{
			DebrisId:   addDebirsReq.DebrisId,
			DebrisName: addDebirsReq.DebrisName,
			Angle:      addDebirsReq.Angle,
			Speed:      addDebirsReq.Speed,
			Height:     addDebirsReq.Height,
			Volunme:    addDebirsReq.Volume,
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

func GetDebrisList(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.QueryListReq
		if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
			ParamsTypeErrorJSONResp(err.Error(), c)
			return
		}

		err := isIntRequiredParamsEmpty(int(req.Page),
			int(req.PageSize))
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}

		sortType, ok := services.SortTypeValue[req.SortType]
		if !ok {
			sortType = services.SORTTYPE_TIME
		}

		params := &services.QueryObjectsParams{
			ModelStruct: new(db.Debris),
			Page:        req.Page,
			PageSize:    req.PageSize,
			SortType:    sortType,
			SearchInput: req.SearchConditions,
		}

		if len(req.SearchConditions) != 0 {
			params.SearchIndex = append(params.SearchIndex, "debris_id")
			params.SearchIndex = append(params.SearchIndex, "debris_name")
		}

		sqlRows, err := s.QueryObjectsWithPage(params)
		if err != nil {
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
