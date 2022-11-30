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
		var addDebirsReq models.AddDebirsReq
		if err := c.ShouldBindBodyWith(&addDebirsReq, binding.JSON); err != nil {
			return
		}
		debirs := &db.Debris{
			DebrisId:   addDebirsReq.DebrisId,
			DebrisName: addDebirsReq.DebrisName,
			Angle:      addDebirsReq.Angle,
			Speed:      addDebirsReq.Speed,
			Height:     addDebirsReq.Height,
			Volunme:    addDebirsReq.Volume,
			Type:       db.LARGE,
		}

		err := s.InsertOneObjertToDB(debirs)
		if err != nil {
			return
		}

		SuccessfulJSONResp("", c)
	}
}

func GetDebrisList(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.QueryListReq
		if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
			return
		}

		params := &services.QueryObjectsParams{
			ModelStruct: new(db.Debris),
			Page:        req.Page,
			PageSize:    req.PageSize,
			SortType:    services.SORTTYPE_TIME_STR,
			SearchInput: req.SearchConditions,
		}

		params.SearchIndex = append(params.SearchIndex, "debris_id")
		params.SearchIndex = append(params.SearchIndex, "debris_name")

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
				ServerErrorJSONResp(c)
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
