package handlers

import (
	"encoding/json"
	"starnet-demo/src/contract"
	"starnet-demo/src/db"
	"starnet-demo/src/models"
	"starnet-demo/src/services"
	"strconv"
	"time"

	"chainmaker.org/chainmaker/pb-go/v2/common"
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
			req.DebrisName, req.Type, req.DebrisSource)
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}

		err = checkTheKeyRule(req.DebrisId)
		if err != nil {
			ParamsFormatErrorJSONResp(err.Error(), c)
			return
		}

		debrisType, ok := db.DebrisTypeValue[req.Type]
		if !ok {
			ParamsValueJSONResp("debirs type not as expected", c)
			return
		}

		client, err := s.GetSdkClient(s.GetMasterChainUserName() + s.GetMasterChainId())
		if err != nil {
			NotInChainJSONResp(err.Error(), c)
			return
		}

		debirs := &db.Debris{
			DebrisId:     req.DebrisId,
			DebrisName:   req.DebrisName,
			DebrisSource: req.DebrisSource,
			Angle:        req.Angle,
			Speed:        req.Speed,
			Height:       req.Height,
			Volunme:      req.Volume,
			Type:         debrisType,
			BlockChainField: db.BlockChainField{
				ChainId: s.GetMasterChainId(),
			},
		}

		err = s.InsertOneObjertToDB(debirs)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		kvs := contract.DebrisConvert(debirs)

		go s.SendTxToBlockChain(s.GetMasterContractName(),
			contract.MASTER_CONTRACT_FUNC_NAME_PUT_DEBRIS, client,
			kvs, debirs, &debirs.BlockChainField)

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

		sqlRows, total, err := s.QueryObjectsWithPage(params)
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
				DebrisId:     debris.DebrisId,
				DebrisName:   debris.DebrisName,
				DebrisSource: debris.DebrisSource,
				Angle:        debris.Angle,
				Speed:        debris.Speed,
				Height:       debris.Height,
				Volume:       debris.Volunme,
				Type:         db.DebrisTypeName[debris.Type],
				BaseRespInfo: models.BaseRespInfo{
					Id:        debris.Id,
					LastTime:  debris.LastTime,
					ChainTime: debris.ChainTime,
				},
			})
		}

		SuccessfulJSONRespWithPage(resp, total, c)
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

		sqlRows, err := s.QueryObjectsByCondition(model, "debris_id", debirsId)
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
					DebrisId:     debris.DebrisId,
					DebrisName:   debris.DebrisName,
					DebrisSource: debris.DebrisSource,
					Angle:        debris.Angle,
					Speed:        debris.Speed,
					Height:       debris.Height,
					Volume:       debris.Volunme,
					Type:         db.DebrisTypeName[debris.Type],
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

		err = checkTheKeyRule(req.DebrisId)
		if err != nil {
			ParamsFormatErrorJSONResp(err.Error(), c)
			return
		}

		debrisType, ok := db.DebrisTypeValue[req.Type]
		if !ok {
			ParamsValueJSONResp("debirs type not as expected", c)
			return
		}

		token, ok1 := c.Get("token")
		claims, ok2 := token.(*services.MyClaims)
		if !ok1 || !ok2 {
			ServerErrorJSONResp("get the token from context failed", c)
			return
		}

		execClient, err := s.GetSdkClient(s.GetExecChainUserName() + s.GetExecChainId())
		if err != nil {
			NotInChainJSONResp(err.Error(), c)
			return
		}

		masterClient, err := s.GetSdkClient(s.GetMasterChainUserName() + s.GetMasterChainId())
		if err != nil {
			NotInChainJSONResp(err.Error(), c)
			return
		}
		//------------------------------------------------------------------------------------------------
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
				ChainId: s.GetExecChainId(),
			},
		}

		err = s.InsertOneObjertToDB(debirs)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		debirsKvs := contract.DebrisConvert(debirs)

		// 碎片信息和指令信息上星座链
		go func() {
			resp, err := execClient.InvokeContract(s.GetExecContractName(),
				contract.EXEC_CONTRACT_FUNC_NAME_GEN_INSTRUCTION, "", debirsKvs, -1, true)
			if err != nil {
				s.GetSuLogger().Warnf("invoke gen instruction func failed, err: [%s]\n",
					err.Error())
				return
			}
			if resp.Code != common.TxStatusCode_SUCCESS {
				s.GetSuLogger().Warnf("invoke gen instruction resp no success, code: [%d], msg: [%s]",
					resp.Code, resp.Message)
				return
			}

			var irs *models.InstructionContractResp
			err = json.Unmarshal(resp.ContractResult.Result, &irs)
			if err != nil {
				s.GetSuLogger().Warnf("invoke gen instruction unmarshal resp failed, err: [%s]\n",
					err.Error())
				return
			}
			chainTime := time.Now().Unix()
			debirs.BlockHeight = irs.BlockHeight
			debirs.TxId = irs.TxId
			debirs.ChainTime = chainTime

			err = s.UpdateObject(debirs)
			if err != nil {
				s.GetSuLogger().Warnf("invoke gen instruction update object failed, err: [%s]\n",
					err.Error())
				return
			}

			instructions := contract.InstructionContractRespConvert(irs,
				claims.Name, s.GetExecChainId())

			// 指令信息入库
			err = s.InsertObjectsToDB(instructions)
			if err != nil {
				s.GetSuLogger().Warnf("invoke gen instruction insert sto db failed, err: [%s]\n",
					err.Error())
				return
			}

			// 碎片信息和指令信息上主链
			go func() {
				go s.SendTxToBlockChain(s.GetMasterContractName(),
					contract.MASTER_CONTRACT_FUNC_NAME_PUT_DEBRIS, masterClient,
					debirsKvs, nil, nil)

				for _, v := range instructions {
					ikvs := contract.InstructionConvert(v)
					go s.SendTxToBlockChain(s.GetMasterContractName(),
						contract.MASTER_CONTRACT_FUNC_NAME_PUT_INSTRUCTION, masterClient,
						ikvs, nil, nil)
				}
			}()

			execInstructionTime := time.Now().Unix()
			for _, v := range instructions {
				inExecI := &db.Instruction{
					InstructionId:       v.InstructionId,
					InstructionSource:   v.InstructionSource,
					Type:                v.Type,
					ExecState:           db.INEXEC,
					InstructionContent:  v.InstructionContent,
					DebrisId:            v.DebrisId,
					DebrisName:          v.DebrisName,
					SatelliteId:         v.SatelliteId,
					SatelliteName:       v.SatelliteName,
					GenInstructionTime:  v.GenInstructionTime,
					ExecInstructionTime: execInstructionTime,
					Treaten:             v.Treaten,
					BlockChainField: db.BlockChainField{
						ChainId: s.GetExecChainId(),
					},
				}

				// 开始执行指令信息入库
				err = s.InsertOneObjertToDB(inExecI)
				if err != nil {
					ServerErrorJSONResp(err.Error(), c)
					return
				}

				// 开始执行指令信息上星座链

				eikvs := contract.InstructionConvert(inExecI)
				s.SendTxToBlockChain(s.GetExecContractName(),
					contract.EXEC_CONTRACT_FUNC_NAME_PUT_INSTRUCTION, execClient,
					eikvs, inExecI, &inExecI.BlockChainField)

				// 开始执行指令信息上主链
				go s.SendTxToBlockChain(s.GetMasterContractName(),
					contract.MASTER_CONTRACT_FUNC_NAME_PUT_INSTRUCTION, masterClient,
					eikvs, nil, nil)

				// 假设执行中 -------------------------------------
				time.Sleep(time.Millisecond * 500)
				//--------------------------------------------------

				endExecI := &db.Instruction{
					InstructionId:       v.InstructionId,
					InstructionSource:   v.InstructionSource,
					Type:                v.Type,
					ExecState:           db.EXECSUCCESS,
					InstructionContent:  v.InstructionContent,
					DebrisId:            v.DebrisId,
					DebrisName:          v.DebrisName,
					SatelliteId:         v.SatelliteId,
					SatelliteName:       v.SatelliteName,
					GenInstructionTime:  v.GenInstructionTime,
					ExecInstructionTime: execInstructionTime,
					Treaten:             v.Treaten,
					BlockChainField: db.BlockChainField{
						ChainId: s.GetExecChainId(),
					},
				}
				// 执行结果入库
				err = s.InsertOneObjertToDB(endExecI)
				if err != nil {
					ServerErrorJSONResp(err.Error(), c)
					return
				}

				// 执行结果上星座链
				rikvs := contract.InstructionConvert(endExecI)
				s.SendTxToBlockChain(s.GetExecContractName(),
					contract.EXEC_CONTRACT_FUNC_NAME_PUT_INSTRUCTION, execClient,
					rikvs, endExecI, &endExecI.BlockChainField)

				// 执行结果上主链
				go s.SendTxToBlockChain(s.GetMasterContractName(),
					contract.MASTER_CONTRACT_FUNC_NAME_PUT_INSTRUCTION, masterClient,
					rikvs, nil, nil)

				// 执行指令操作上星座链
				operation := &db.Operation{
					Operator:        claims.Name,
					OperatorIp:      c.ClientIP(),
					OperationTime:   execInstructionTime,
					OperationRecord: "执行指令：" + v.InstructionId,
					SatelliteId:     v.SatelliteId,
					SatelliteName:   v.SatelliteName,
					BlockChainField: db.BlockChainField{
						ChainId: s.GetExecChainId(),
					},
				}

				if err := s.InsertOneObjertToDB(operation); err != nil {
					ServerErrorJSONResp(err.Error(), c)
					return
				}
				okvs := contract.OperationConvert(operation)

				go s.SendTxToBlockChain(s.GetExecContractName(),
					contract.EXEC_CONTRACT_FUNC_NAME_PUT_OPERATION, execClient,
					okvs, operation, &operation.BlockChainField)

				go s.SendTxToBlockChain(s.GetMasterContractName(),
					contract.MASTER_CONTRACT_FUNC_NAME_PUT_OPERATION, masterClient,
					okvs, nil, nil)
			}
		}()

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

		sqlRows, total, err := s.QueryObjectsWithPageSC(params)
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

		SuccessfulJSONRespWithPage(resp, total, c)
	}
}

func TraceGetDebrisList(s *services.Server) gin.HandlerFunc {
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
			ModelStruct: new(db.Debris),
			Page:        int32(page),
			PageSize:    int32(pageSize),
			SortType:    sortType,
			SearchInput: searchInput,
			SearchIndex: make([]string, 0),
			GroupIndex:  "debris_id",
		}

		if len(searchInput) != 0 {
			params.SearchIndex = append(params.SearchIndex, "debris_id")
			params.SearchIndex = append(params.SearchIndex, "debris_name")
		}

		sqlRows, total, err := s.QueryLatestObjectsWithPage(params)
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

		SuccessfulJSONRespWithPage(resp, total, c)
	}
}
