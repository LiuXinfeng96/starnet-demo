package handlers

import (
	"starnet-demo/src/contract"
	"starnet-demo/src/db"
	"starnet-demo/src/models"
	"starnet-demo/src/services"
	"strconv"
	"time"

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

		err := isStringRequiredParamsEmpty(req.InstructionId,
			req.InstructionContent, req.DebrisId, req.DebrisName,
			req.SatelliteId, req.SatelliteName)
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}

		err = checkTheKeyRule(req.InstructionId)
		if err != nil {
			ParamsFormatErrorJSONResp(err.Error(), c)
			return
		}

		token, ok1 := c.Get("token")
		claims, ok2 := token.(*services.MyClaims)
		if !ok1 || !ok2 {
			ServerErrorJSONResp("get token failed", c)
			return
		}

		masterClient, err := s.GetSdkClient(s.GetMasterChainUserName() + s.GetMasterChainId())
		if err != nil {
			NotInChainJSONResp(err.Error(), c)
			return
		}

		execClient, err := s.GetSdkClient(s.GetExecChainUserName() + s.GetExecChainId())
		if err != nil {
			NotInChainJSONResp(err.Error(), c)
			return
		}

		genInstructionTime := time.Now().Unix()
		noExecI := &db.Instruction{
			InstructionId:      req.InstructionId,
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
		// 未执行指令信息入库
		err = s.InsertOneObjertToDB(noExecI)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		noikvs := contract.InstructionConvert(noExecI)
		// 未执行指令信息上主链
		go s.SendTxToBlockChain(s.GetMasterContractName(),
			contract.MASTER_CONTRACT_FUNC_NAME_PUT_INSTRUCTION, masterClient,
			noikvs, noExecI, &noExecI.BlockChainField)

		operation := &db.Operation{
			Operator:        claims.Name,
			OperatorIp:      c.ClientIP(),
			OperationTime:   genInstructionTime,
			SatelliteId:     req.SatelliteId,
			SatelliteName:   req.SatelliteName,
			OperationRecord: "编辑指令：" + req.InstructionId,
			BlockChainField: db.BlockChainField{
				ChainId: s.GetMasterChainId(),
			},
		}

		// 编辑操作入库
		err = s.InsertOneObjertToDB(operation)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		okvs := contract.OperationConvert(operation)

		go s.SendTxToBlockChain(s.GetMasterContractName(),
			contract.MASTER_CONTRACT_FUNC_NAME_PUT_OPERATION, masterClient,
			okvs, operation, &operation.BlockChainField)

		execInstructionTime := time.Now().Unix()
		inExecI := &db.Instruction{
			InstructionId:       req.InstructionId,
			InstructionSource:   claims.Name,
			Type:                db.OPERATION,
			ExecState:           db.INEXEC,
			InstructionContent:  req.InstructionContent,
			DebrisId:            req.DebrisId,
			DebrisName:          req.DebrisName,
			SatelliteId:         req.SatelliteId,
			SatelliteName:       req.SatelliteName,
			GenInstructionTime:  genInstructionTime,
			ExecInstructionTime: execInstructionTime,
			BlockChainField: db.BlockChainField{
				ChainId: s.GetMasterChainId(),
			},
		}
		// 执行指令入库
		err = s.InsertOneObjertToDB(inExecI)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		ineikvs := contract.InstructionConvert(inExecI)
		go s.SendTxToBlockChain(s.GetExecContractName(),
			contract.EXEC_CONTRACT_FUNC_NAME_PUT_INSTRUCTION, execClient,
			ineikvs, inExecI, &inExecI.BlockChainField)

		go s.SendTxToBlockChain(s.GetMasterContractName(),
			contract.MASTER_CONTRACT_FUNC_NAME_PUT_INSTRUCTION, masterClient,
			ineikvs, nil, nil)

		//假设执行中-------------------------------
		time.Sleep(time.Millisecond * 500)
		//-----------------------------------------

		endExecI := &db.Instruction{
			InstructionId:       req.InstructionId,
			InstructionSource:   claims.Name,
			Type:                db.OPERATION,
			ExecState:           db.EXECSUCCESS,
			InstructionContent:  req.InstructionContent,
			DebrisId:            req.DebrisId,
			DebrisName:          req.DebrisName,
			SatelliteId:         req.SatelliteId,
			SatelliteName:       req.SatelliteName,
			GenInstructionTime:  genInstructionTime,
			ExecInstructionTime: execInstructionTime,
			BlockChainField: db.BlockChainField{
				ChainId: s.GetMasterChainId(),
			},
		}

		// 执行结果入库
		err = s.InsertOneObjertToDB(endExecI)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		rikvs := contract.InstructionConvert(endExecI)

		// 执行结果上星座链
		go s.SendTxToBlockChain(s.GetExecContractName(),
			contract.EXEC_CONTRACT_FUNC_NAME_PUT_INSTRUCTION, execClient,
			rikvs, endExecI, &endExecI.BlockChainField)

		// 执行结果上主链
		go s.SendTxToBlockChain(s.GetMasterContractName(),
			contract.MASTER_CONTRACT_FUNC_NAME_PUT_INSTRUCTION, masterClient,
			rikvs, nil, nil)

		operation = &db.Operation{
			Operator:        claims.Name,
			OperationTime:   execInstructionTime,
			OperatorIp:      c.ClientIP(),
			SatelliteId:     req.SatelliteId,
			SatelliteName:   req.SatelliteName,
			OperationRecord: "执行指令：" + req.InstructionId,
			BlockChainField: db.BlockChainField{
				ChainId: s.GetMasterChainId(),
			},
		}

		// 执行操作入库
		err = s.InsertOneObjertToDB(operation)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		opkvs := contract.OperationConvert(operation)
		// 执行操作上星座链
		go s.SendTxToBlockChain(s.GetExecContractName(),
			contract.EXEC_CONTRACT_FUNC_NAME_PUT_OPERATION, execClient,
			opkvs, operation, &operation.BlockChainField)

		// 执行操作上主链
		go s.SendTxToBlockChain(s.GetMasterContractName(),
			contract.MASTER_CONTRACT_FUNC_NAME_PUT_OPERATION, masterClient,
			opkvs, nil, nil)

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

		sqlRows, total, err := s.QueryObjectsWithPage(params)
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

		SuccessfulJSONRespWithPage(resp, total, c)
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

		SuccessfulJSONResp(resp, c)
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

		sqlRows, total, err := s.QueryObjectsWithPage(params)
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

		SuccessfulJSONRespWithPage(resp, total, c)
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

		sqlRows, total, err := s.QueryObjectsWithPageSC(params)
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

		SuccessfulJSONRespWithPage(resp, total, c)
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

		sqlRows, total, err := s.QueryObjectsWithPage(params)
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

		SuccessfulJSONRespWithPage(resp, total, c)
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

		sqlRows, total, err := s.QueryLatestObjectsWithPage(params)
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

		SuccessfulJSONRespWithPage(resp, total, c)
	}
}
