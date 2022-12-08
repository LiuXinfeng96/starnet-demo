package handlers

import (
	"starnet-demo/src/contract"
	"starnet-demo/src/db"
	"starnet-demo/src/models"
	"starnet-demo/src/services"
	"strconv"
	"time"

	"chainmaker.org/chainmaker/pb-go/v2/common"
	"github.com/gin-gonic/gin"
)

func ControlAddInstruction(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := checkTheAccessPermission(c, db.CONTROL); err != nil {
			WithoutPermissionJSONResp(err.Error(), c)
			return
		}

		var req models.AddInstructionReq
		if err := c.ShouldBindJSON(&req); err != nil {
			ParamsTypeErrorJSONResp(err.Error(), c)
			return
		}

		err := isStringRequiredParamsEmpty(req.InstructtionId,
			req.InstructionContent, req.DebrisId, req.DebrisName,
			req.SatelliteId, req.SatelliteName)
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}

		token, ok1 := c.Get("token")
		claims, ok2 := token.(*services.MyClaims)
		if !ok1 || !ok2 {
			ServerErrorJSONResp("get token failed", c)
			return
		}

		genInstructionTime := time.Now().Unix()
		instruction := &db.Instruction{
			InstructionId:      req.InstructtionId,
			InstructionSource:  claims.Name,
			Type:               db.OPERATION,
			ExecState:          db.NOTEXEC,
			InstructionContent: req.InstructionContent,
			DebrisId:           req.DebrisId,
			DebrisName:         req.DebrisName,
			SatelliteId:        req.SatelliteId,
			SatelliteName:      req.SatelliteName,
			GenInstructionTime: genInstructionTime,
			BlockChainField: db.BlockChainField{
				ChainId: s.GetMasterChainId(),
			},
		}

		// 未执行指令信息上主链
		masterClient, err := s.GetSdkClient(claims.Name + s.GetMasterChainId())
		if err != nil {
			NotInChainJSONResp(err.Error(), c)
			return
		}

		kvs := contract.InstructionConvert(instruction)

		mResp, err := masterClient.InvokeContract(s.GetMasterContractName(),
			contract.MASTER_CONTRACT_FUNC_NAME_PUT_INSTRUCTION, "", kvs, -1, true)
		if err != nil {
			PutChainFailJSONResp(err.Error(), c)
			return
		}
		if mResp.Code != common.TxStatusCode_SUCCESS {
			PutChainFailJSONResp(mResp.Message, c)
			return
		}

		instruction.BlockChainField, err = GetBlockChainFiledFromResp(mResp.ContractResult)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		// 未执行指令信息入库
		err = s.InsertOneObjertToDB(instruction)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		// 编辑操作上主链
		operation := &db.Operation{
			Operator:        claims.Name,
			OperatorIp:      c.ClientIP(),
			OperationTime:   genInstructionTime,
			SatelliteId:     instruction.SatelliteId,
			SatelliteName:   instruction.SatelliteName,
			OperationRecord: "编辑指令：" + instruction.InstructionId,
			BlockChainField: db.BlockChainField{
				ChainId: s.GetMasterChainId(),
			},
		}

		kvs = contract.OperationConvert(operation)

		omp, err := masterClient.InvokeContract(s.GetMasterContractName(),
			contract.MASTER_CONTRACT_FUNC_NAME_PUT_OPERATION, "", kvs, -1, true)
		if err != nil {
			PutChainFailJSONResp(err.Error(), c)
			return
		}
		if omp.Code != common.TxStatusCode_SUCCESS {
			PutChainFailJSONResp(omp.Message, c)
			return
		}

		operation.BlockChainField, err = GetBlockChainFiledFromResp(omp.ContractResult)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		// 编辑操作入库
		err = s.InsertOneObjertToDB(operation)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		// 执行指令上星座链
		execClient, err := s.GetSdkClient(claims.Name + s.GetExecChainId())
		if err != nil {
			NotInChainJSONResp(err.Error(), c)
			return
		}
		instruction.ExecState = db.INEXEC
		instruction.ExecInstructionTime = time.Now().Unix()

		kvs = contract.InstructionConvert(instruction)

		eResp, err := execClient.InvokeContract(s.GetExecContractName(),
			contract.EXEC_CONTRACT_FUNC_NAME_PUT_INSTRUCTION, "", kvs, -1, true)
		if err != nil {
			PutChainFailJSONResp(err.Error(), c)
			return
		}
		if eResp.Code != common.TxStatusCode_SUCCESS {
			PutChainFailJSONResp(eResp.Message, c)
			return
		}

		instruction.BlockChainField, err = GetBlockChainFiledFromResp(eResp.ContractResult)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		// 执行指令上主链
		go func() {
			mir, err := masterClient.InvokeContract(s.GetMasterContractName(),
				contract.MASTER_CONTRACT_FUNC_NAME_PUT_INSTRUCTION, "", kvs, -1, true)
			if err != nil {
				s.GetSuLogger().Warn(err)
				return
			}
			if mir.Code != common.TxStatusCode_SUCCESS {
				s.GetSuLogger().Warnf("invoke contract failed: tx status code: [%d], msg: [%s]\n",
					mir.Code, mir.Message)
				return
			}
		}()

		// 执行指令上入库
		err = s.InsertOneObjertToDB(instruction)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		// 执行结果上星座链
		instruction.ExecState = db.EXECSUCCESS
		kvs = contract.InstructionConvert(instruction)

		eir, err := execClient.InvokeContract(s.GetExecContractName(),
			contract.EXEC_CONTRACT_FUNC_NAME_PUT_INSTRUCTION, "", kvs, -1, true)
		if err != nil {
			PutChainFailJSONResp(err.Error(), c)
			return
		}
		if eir.Code != common.TxStatusCode_SUCCESS {
			PutChainFailJSONResp(eir.Message, c)
			return
		}

		instruction.BlockChainField, err = GetBlockChainFiledFromResp(eir.ContractResult)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		// 执行结果上主链
		go func() {
			mir, err := masterClient.InvokeContract(s.GetMasterContractName(),
				contract.MASTER_CONTRACT_FUNC_NAME_PUT_INSTRUCTION, "", kvs, -1, true)
			if err != nil {
				s.GetSuLogger().Warn(err)
				return
			}
			if mir.Code != common.TxStatusCode_SUCCESS {
				s.GetSuLogger().Warnf("invoke contract failed: tx status code: [%d], msg: [%s]\n",
					mir.Code, mir.Message)
				return
			}
		}()

		// 执行结果入库
		err = s.InsertOneObjertToDB(instruction)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		// 执行操作上星座链
		operation = &db.Operation{
			Operator:        "卫星执行系统",
			OperationTime:   genInstructionTime,
			SatelliteId:     instruction.SatelliteId,
			SatelliteName:   instruction.SatelliteName,
			OperationRecord: "执行指令：" + instruction.InstructionId,
			BlockChainField: db.BlockChainField{
				ChainId: s.GetMasterChainId(),
			},
		}

		kvs = contract.OperationConvert(operation)

		oep, err := execClient.InvokeContract(s.GetExecContractName(),
			contract.EXEC_CONTRACT_FUNC_NAME_PUT_OPERATION, "", kvs, -1, true)
		if err != nil {
			PutChainFailJSONResp(err.Error(), c)
			return
		}
		if oep.Code != common.TxStatusCode_SUCCESS {
			PutChainFailJSONResp(oep.Message, c)
			return
		}

		operation.BlockChainField, err = GetBlockChainFiledFromResp(oep.ContractResult)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		// 执行操作上主链
		go func() {
			mor, err := masterClient.InvokeContract(s.GetMasterContractName(),
				contract.MASTER_CONTRACT_FUNC_NAME_PUT_OPERATION, "", kvs, -1, true)
			if err != nil {
				s.GetSuLogger().Warn(err)
				return
			}
			if mor.Code != common.TxStatusCode_SUCCESS {
				s.GetSuLogger().Warnf("invoke contract failed: tx status code: [%d], msg: [%s]\n",
					mor.Code, mor.Message)
				return
			}
		}()

		// 执行操作入库
		err = s.InsertOneObjertToDB(operation)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		SuccessfulJSONResp("", c)
	}
}

func ControlGetInstructionList(s *services.Server) gin.HandlerFunc {
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
			ModelStruct: new(db.Instruction),
			Page:        int32(page),
			PageSize:    int32(pageSize),
			SortType:    sortType,
			SearchInput: searchInput,
			SearchIndex: make([]string, 0),
			QueryMap:    make(map[string]string),
		}

		if len(searchInput) != 0 {
			params.SearchIndex = append(params.SearchIndex, "instruction_id")
		}

		params.QueryMap["exec_state"] = strconv.Itoa(int(db.NOTEXEC))

		sqlRows, err := s.QueryObjectsWithPage(params)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		defer sqlRows.Close()

		resp := make([]*models.InstructionInfo, 0)

		for sqlRows.Next() {
			var instruction db.Instruction
			err := s.ScanRows(sqlRows, &instruction)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.InstructionInfo{
				InstructionId:       instruction.InstructionId,
				InstructionSource:   instruction.InstructionSource,
				InstructionContent:  instruction.InstructionContent,
				InstructionType:     db.InstructionTypeName[instruction.Type],
				ExecInstructionTime: instruction.ExecInstructionTime,
				GenInstructionTime:  instruction.GenInstructionTime,
				DebrisId:            instruction.DebrisId,
				DebrisName:          instruction.DebrisName,
				SatelliteId:         instruction.SatelliteId,
				SatelliteName:       instruction.SatelliteName,
				BaseRespInfo: models.BaseRespInfo{
					Id:        instruction.Id,
					LastTime:  instruction.LastTime,
					ChainTime: instruction.ChainTime,
				},
			})
		}

		SuccessfulJSONRespWithPage(resp, len(resp), c)
	}
}

func TraceGetInstruction(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := checkTheAccessPermission(c, db.TRACE); err != nil {
			WithoutPermissionJSONResp(err.Error(), c)
			return
		}
		instructionId := c.Query("instructionId")

		err := isStringRequiredParamsEmpty(instructionId)
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}

		model := new(db.Instruction)

		sqlRows, err := s.QueryObjectsByCondition(model, "instruction_id", instructionId)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}
		defer sqlRows.Close()

		resp := make([]*models.InstructionDetails, 0)

		for sqlRows.Next() {
			var instruction db.Instruction
			err := s.ScanRows(sqlRows, &instruction)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.InstructionDetails{
				ExecState: db.ExecStateName[instruction.ExecState],
				InstructionInfo: models.InstructionInfo{
					InstructionId:       instruction.InstructionId,
					InstructionSource:   instruction.InstructionSource,
					InstructionContent:  instruction.InstructionContent,
					InstructionType:     db.InstructionTypeName[instruction.Type],
					ExecInstructionTime: instruction.ExecInstructionTime,
					GenInstructionTime:  instruction.GenInstructionTime,
					DebrisId:            instruction.DebrisId,
					DebrisName:          instruction.DebrisName,
					SatelliteId:         instruction.SatelliteId,
					SatelliteName:       instruction.SatelliteName,
					BaseRespInfo: models.BaseRespInfo{
						Id:        instruction.Id,
						LastTime:  instruction.LastTime,
						ChainTime: instruction.ChainTime,
					},
				},
				HistoryRespInfo: models.HistoryRespInfo{
					ChainId:     instruction.ChainId,
					BlockHeight: instruction.BlockHeight,
					TxId:        instruction.TxId,
				},
			})
		}

		SuccessfulJSONRespWithPage(resp, len(resp), c)
	}
}

func ControlGetExecResultList(s *services.Server) gin.HandlerFunc {
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
			ModelStruct: new(db.Instruction),
			Page:        int32(page),
			PageSize:    int32(pageSize),
			SortType:    sortType,
			SearchInput: searchInput,
			SearchIndex: make([]string, 0),
		}

		if len(searchInput) != 0 {
			params.SearchIndex = append(params.SearchIndex, "instruction_id")
		}

		sqlRows, err := s.QueryObjectsWithPage(params)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		defer sqlRows.Close()

		resp := make([]*models.InstructionResultInfo, 0)

		for sqlRows.Next() {
			var instruction db.Instruction
			err := s.ScanRows(sqlRows, &instruction)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.InstructionResultInfo{
				InstructionInfo: models.InstructionInfo{
					InstructionId:       instruction.InstructionId,
					InstructionSource:   instruction.InstructionSource,
					InstructionContent:  instruction.InstructionContent,
					InstructionType:     db.InstructionTypeName[instruction.Type],
					ExecInstructionTime: instruction.ExecInstructionTime,
					GenInstructionTime:  instruction.GenInstructionTime,
					DebrisId:            instruction.DebrisId,
					DebrisName:          instruction.DebrisName,
					SatelliteId:         instruction.SatelliteId,
					SatelliteName:       instruction.SatelliteName,
					BaseRespInfo: models.BaseRespInfo{
						Id:        instruction.Id,
						LastTime:  instruction.LastTime,
						ChainTime: instruction.ChainTime,
					},
				},
				ExecState: db.ExecStateName[instruction.ExecState],
			})
		}

		SuccessfulJSONRespWithPage(resp, len(resp), c)
	}
}

func ControlGetExecResult(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := checkTheAccessPermission(c, db.CONTROL); err != nil {
			WithoutPermissionJSONResp(err.Error(), c)
			return
		}

		idStr := c.Query("id")

		err := isStringRequiredParamsEmpty(idStr)
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			ParamsTypeErrorJSONResp(err.Error(), c)
			return
		}

		instruction := new(db.Instruction)

		err = s.QueryObjectById(instruction, int32(id))
		if err != nil {
			SuccessfulJSONResp("", c)
			return
		}

		SuccessfulJSONResp(&models.InstructionResultInfo{
			ExecState: db.ExecStateName[instruction.ExecState],
			InstructionInfo: models.InstructionInfo{
				InstructionId:       instruction.InstructionId,
				InstructionSource:   instruction.InstructionSource,
				InstructionContent:  instruction.InstructionContent,
				InstructionType:     db.InstructionTypeName[instruction.Type],
				ExecInstructionTime: instruction.ExecInstructionTime,
				GenInstructionTime:  instruction.GenInstructionTime,
				DebrisId:            instruction.DebrisId,
				DebrisName:          instruction.DebrisName,
				SatelliteId:         instruction.SatelliteId,
				SatelliteName:       instruction.SatelliteName,
				BaseRespInfo: models.BaseRespInfo{
					Id:        instruction.Id,
					LastTime:  instruction.LastTime,
					ChainTime: instruction.ChainTime,
				},
			},
		}, c)
	}
}

func ExecGetInstructionList(s *services.Server) gin.HandlerFunc {
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
			ModelStruct: new(db.Instruction),
			Page:        int32(page),
			PageSize:    int32(pageSize),
			SortType:    sortType,
			SearchInput: searchInput,
			SearchIndex: make([]string, 0),
			QueryMap:    make(map[string]string),
		}

		if len(searchInput) != 0 {
			params.SearchIndex = append(params.SearchIndex, "instruction_id")
		}

		params.QueryMap["exec_state"] = strconv.Itoa(int(db.NOTEXEC))

		sqlRows, err := s.QueryObjectsWithPageSC(params)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		defer sqlRows.Close()

		resp := make([]*models.InstructionInfo, 0)

		for sqlRows.Next() {
			var instruction db.Instruction
			err := s.ScanRows(sqlRows, &instruction)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.InstructionInfo{
				InstructionId:       instruction.InstructionId,
				InstructionSource:   instruction.InstructionSource,
				InstructionContent:  instruction.InstructionContent,
				InstructionType:     db.InstructionTypeName[instruction.Type],
				ExecInstructionTime: instruction.ExecInstructionTime,
				GenInstructionTime:  instruction.GenInstructionTime,
				DebrisId:            instruction.DebrisId,
				DebrisName:          instruction.DebrisName,
				SatelliteId:         instruction.SatelliteId,
				SatelliteName:       instruction.SatelliteName,
				BaseRespInfo: models.BaseRespInfo{
					Id:        instruction.Id,
					LastTime:  instruction.LastTime,
					ChainTime: instruction.ChainTime,
				},
			})
		}

		SuccessfulJSONRespWithPage(resp, len(resp), c)
	}
}

func ExecGetInstruction(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := checkTheAccessPermission(c, db.EXEC); err != nil {
			WithoutPermissionJSONResp(err.Error(), c)
			return
		}

		idStr := c.Query("id")

		err := isStringRequiredParamsEmpty(idStr)
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			ParamsTypeErrorJSONResp(err.Error(), c)
			return
		}

		instruction := new(db.Instruction)

		err = s.QueryObjectById(instruction, int32(id))
		if err != nil {
			SuccessfulJSONResp("", c)
			return
		}

		SuccessfulJSONResp(&models.InstructionDetails{
			ExecState: db.ExecStateName[instruction.ExecState],
			InstructionInfo: models.InstructionInfo{
				InstructionId:       instruction.InstructionId,
				InstructionSource:   instruction.InstructionSource,
				InstructionContent:  instruction.InstructionContent,
				InstructionType:     db.InstructionTypeName[instruction.Type],
				ExecInstructionTime: instruction.ExecInstructionTime,
				GenInstructionTime:  instruction.GenInstructionTime,
				DebrisId:            instruction.DebrisId,
				DebrisName:          instruction.DebrisName,
				SatelliteId:         instruction.SatelliteId,
				SatelliteName:       instruction.SatelliteName,
				BaseRespInfo: models.BaseRespInfo{
					Id:        instruction.Id,
					LastTime:  instruction.LastTime,
					ChainTime: instruction.ChainTime,
				},
			},
			HistoryRespInfo: models.HistoryRespInfo{
				ChainId:     instruction.ChainId,
				BlockHeight: instruction.BlockHeight,
				TxId:        instruction.TxId,
			},
		}, c)
	}
}

func ExecGetExecResultList(s *services.Server) gin.HandlerFunc {
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
			ModelStruct: new(db.Instruction),
			Page:        int32(page),
			PageSize:    int32(pageSize),
			SortType:    sortType,
			SearchInput: searchInput,
			SearchIndex: make([]string, 0),
		}

		if len(searchInput) != 0 {
			params.SearchIndex = append(params.SearchIndex, "instruction_id")
		}

		sqlRows, err := s.QueryObjectsWithPageSC(params)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		defer sqlRows.Close()

		resp := make([]*models.InstructionResultInfo, 0)

		for sqlRows.Next() {
			var instruction db.Instruction
			err := s.ScanRows(sqlRows, &instruction)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.InstructionResultInfo{
				InstructionInfo: models.InstructionInfo{
					InstructionId:       instruction.InstructionId,
					InstructionSource:   instruction.InstructionSource,
					InstructionContent:  instruction.InstructionContent,
					InstructionType:     db.InstructionTypeName[instruction.Type],
					ExecInstructionTime: instruction.ExecInstructionTime,
					GenInstructionTime:  instruction.GenInstructionTime,
					DebrisId:            instruction.DebrisId,
					DebrisName:          instruction.DebrisName,
					SatelliteId:         instruction.SatelliteId,
					SatelliteName:       instruction.SatelliteName,
					BaseRespInfo: models.BaseRespInfo{
						Id:        instruction.Id,
						LastTime:  instruction.LastTime,
						ChainTime: instruction.ChainTime,
					},
				},
				ExecState: db.ExecStateName[instruction.ExecState],
			})
		}

		SuccessfulJSONRespWithPage(resp, len(resp), c)
	}
}

func ExecGetExecResult(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := checkTheAccessPermission(c, db.EXEC); err != nil {
			WithoutPermissionJSONResp(err.Error(), c)
			return
		}

		idStr := c.Query("id")

		err := isStringRequiredParamsEmpty(idStr)
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			ParamsTypeErrorJSONResp(err.Error(), c)
			return
		}

		instruction := new(db.Instruction)

		err = s.QueryObjectById(instruction, int32(id))
		if err != nil {
			SuccessfulJSONResp("", c)
			return
		}

		SuccessfulJSONResp(&models.InstructionResultInfo{
			ExecState: db.ExecStateName[instruction.ExecState],
			InstructionInfo: models.InstructionInfo{
				InstructionId:       instruction.InstructionId,
				InstructionSource:   instruction.InstructionSource,
				InstructionContent:  instruction.InstructionContent,
				InstructionType:     db.InstructionTypeName[instruction.Type],
				ExecInstructionTime: instruction.ExecInstructionTime,
				GenInstructionTime:  instruction.GenInstructionTime,
				DebrisId:            instruction.DebrisId,
				DebrisName:          instruction.DebrisName,
				SatelliteId:         instruction.SatelliteId,
				SatelliteName:       instruction.SatelliteName,
				BaseRespInfo: models.BaseRespInfo{
					Id:        instruction.Id,
					LastTime:  instruction.LastTime,
					ChainTime: instruction.ChainTime,
				},
			},
		}, c)
	}
}

func TraceGetInstructionList(s *services.Server) gin.HandlerFunc {
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
			ModelStruct: new(db.Instruction),
			Page:        int32(page),
			PageSize:    int32(pageSize),
			SortType:    sortType,
			SearchInput: searchInput,
			SearchIndex: make([]string, 0),
			GroupIndex:  "instruction_id",
		}

		if len(searchInput) != 0 {
			params.SearchIndex = append(params.SearchIndex, "instruction_id")
		}

		sqlRows, err := s.QueryLatestObjectsWithPage(params)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		defer sqlRows.Close()

		resp := make([]*models.InstructionInfo, 0)

		for sqlRows.Next() {
			var instruction db.Instruction
			err := s.ScanRows(sqlRows, &instruction)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.InstructionInfo{
				InstructionId:       instruction.InstructionId,
				InstructionSource:   instruction.InstructionSource,
				InstructionContent:  instruction.InstructionContent,
				InstructionType:     db.InstructionTypeName[instruction.Type],
				ExecInstructionTime: instruction.ExecInstructionTime,
				GenInstructionTime:  instruction.GenInstructionTime,
				DebrisId:            instruction.DebrisId,
				DebrisName:          instruction.DebrisName,
				SatelliteId:         instruction.SatelliteId,
				SatelliteName:       instruction.SatelliteName,
				BaseRespInfo: models.BaseRespInfo{
					Id:        instruction.Id,
					LastTime:  instruction.LastTime,
					ChainTime: instruction.ChainTime,
				},
			})
		}

		SuccessfulJSONRespWithPage(resp, len(resp), c)
	}
}
