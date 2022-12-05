package handlers

import (
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

		err := isStringRequiredParamsEmpty(req.InstructtionId,
			req.InstructionContent, req.DebrisId, req.DebrisName,
			req.SatelliteId, req.SatelliteName)
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}

		// 1. 上主链(未执行状态)
		// 2. 入库（未执行状态）
		// 3. 上星座链（执行中状态）
		// 4. 返回入库（执行中状态）
		// 5. 执行结果上星座链
		// 6. 执行结果上主链
		// 7. 执行结果入库（执行成功/失败状态）
		// 8. 避障操作上主链
		// 9. 避障操作入库

		token, ok1 := c.Get("token")
		claims, ok2 := token.(*services.MyClaims)
		if !ok1 || !ok2 {
			ServerErrorJSONResp("get token failed", c)
			return
		}

		genInstructionTime := time.Now().Unix()
		// 未执行入库
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
		}

		err = s.InsertOneObjertToDB(instruction)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		operation := &db.Operation{
			Operator:        instruction.InstructionSource,
			OperatorIp:      c.ClientIP(),
			OperationTime:   genInstructionTime,
			SatelliteId:     instruction.SatelliteId,
			SatelliteName:   instruction.SatelliteName,
			OperationRecord: "编辑指令：" + instruction.InstructionId,
		}

		err = s.InsertOneObjertToDB(operation)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		time.Sleep(time.Millisecond * 500)

		// 执行中入库

		execInstructionTime := time.Now().Unix()
		instruction = &db.Instruction{
			InstructionId:       req.InstructtionId,
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
		}

		err = s.InsertOneObjertToDB(instruction)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		operation = &db.Operation{
			Operator:        instruction.SatelliteName,
			OperatorIp:      "",
			OperationTime:   execInstructionTime,
			SatelliteId:     instruction.SatelliteId,
			SatelliteName:   instruction.SatelliteName,
			OperationRecord: "执行指令：" + instruction.InstructionId,
		}

		err = s.InsertOneObjertToDB(operation)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		time.Sleep(time.Millisecond * 1000)
		// 执行完成入库
		execState := db.EXECSUCCESS
		instruction = &db.Instruction{
			InstructionId:       req.InstructtionId,
			InstructionSource:   claims.Name,
			Type:                db.OPERATION,
			ExecState:           execState,
			InstructionContent:  req.InstructionContent,
			DebrisId:            req.DebrisId,
			DebrisName:          req.DebrisName,
			SatelliteId:         req.SatelliteId,
			SatelliteName:       req.SatelliteName,
			GenInstructionTime:  genInstructionTime,
			ExecInstructionTime: execInstructionTime,
		}

		err = s.InsertOneObjertToDB(instruction)
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
