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

func JWTAuthMiddleware(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("token")
		if len(token) == 0 {
			ParamsMissingJSONResp("the token not found", c)
			c.Abort()
			return
		}

		claims, err := s.ParseToken(token)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			c.Abort()
			return
		}

		c.Set("token", claims)
		c.Next()
	}
}

func Login(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {

		var req models.LoginReq

		if err := c.ShouldBindJSON(&req); err != nil {
			ParamsTypeErrorJSONResp(err.Error(), c)
			return
		}

		err := isStringRequiredParamsEmpty(req.UserName, req.UserPwd)
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}
		user := new(db.User)
		err = s.QueryObjectByCondition(user, "user_name", req.UserName)
		if err != nil {
			NotExistJSONResp(err.Error(), c)
			return
		}

		if req.UserPwd != user.UserPwd {
			PwdErrorJSONResp("", c)
			return
		}

		tokenExpiresAt := time.Now().Add(2 * time.Hour).Unix()
		token, err := s.GenToken(user.Id,
			user.UserName, db.UserRoleTypeName[user.UserRole], tokenExpiresAt)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		// 记录登录日志

		// 1. 上链
		// 2. 存到数据库

		loginLog := &db.LoginLog{
			UserName:  user.UserName,
			LoginIp:   c.ClientIP(),
			LoginTime: time.Now().Unix(),
			BlockChainField: db.BlockChainField{
				ChainId: s.GetMasterChainId(),
			},
		}

		client, err := s.GetSdkClient(user.UserName + s.GetMasterChainId())
		if err != nil {
			NotInChainJSONResp(err.Error(), c)
			return
		}

		kvs := contract.LoginLogConvert(loginLog)

		chainResp, err := client.InvokeContract(s.GetMasterContractName(),
			contract.MASTER_CONTRACT_FUNC_NAME_PUT_LOGINLOG, "", kvs, -1, true)
		if err != nil {
			PutChainFailJSONResp(err.Error(), c)
			return
		}
		if chainResp.Code != common.TxStatusCode_SUCCESS {
			PutChainFailJSONResp(chainResp.Message, c)
			return
		}

		loginLog.BlockChainField, err = GetBlockChainFiledFromResp(chainResp.ContractResult)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		err = s.InsertOneObjertToDB(loginLog)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		SuccessfulJSONResp(&models.LoginInfo{
			UserNickName: user.UserNickName,
			UserRole:     db.UserRoleTypeName[user.UserRole],
			Expires:      tokenExpiresAt,
			Token:        token,
		}, c)
	}
}

func Register(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.RegisterReq
		if err := c.ShouldBindJSON(&req); err != nil {
			ParamsTypeErrorJSONResp(err.Error(), c)
			return
		}

		err := isStringRequiredParamsEmpty(req.UserName, req.UserPwd,
			req.UserPhoneNum, req.UserNickName, req.UserEmail, req.UserRole)
		if err != nil {
			ParamsMissingJSONResp(err.Error(), c)
			return
		}

		err = checkThePhoneNum(req.UserPhoneNum)
		if err != nil {
			ParamsFormatErrorJSONResp(err.Error(), c)
			return
		}

		err = checkTheEmail(req.UserEmail)
		if err != nil {
			ParamsFormatErrorJSONResp(err.Error(), c)
			return
		}

		role, ok := db.UserRoleTypeValue[req.UserRole]
		if !ok {
			ParamsValueJSONResp("role type error", c)
			return
		}

		user := new(db.User)
		err = s.QueryObjectByCondition(user, "user_name", req.UserName)
		if err == nil {
			UniqueIndexJSONResp("用户已存在", c)
			return
		}

		err = s.InsertOneObjertToDB(&db.User{
			UserName:     req.UserName,
			UserRole:     role,
			UserPwd:      req.UserPwd,
			UserNickName: req.UserNickName,
			UserPhoneNum: req.UserPhoneNum,
			UserEmail:    req.UserEmail,
		})
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		SuccessfulJSONResp("", c)
	}
}

func GetUserInfo(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {

		token, ok1 := c.Get("token")
		claims, ok2 := token.(*services.MyClaims)

		if !ok1 || !ok2 {
			ServerErrorJSONResp("get the token from context failed", c)
			return
		}

		user := new(db.User)
		err := s.QueryObjectById(user, claims.Id)
		if err != nil {
			NotExistJSONResp(err.Error(), c)
			return
		}

		SuccessfulJSONResp(&models.UserInfo{
			Id:           user.Id,
			UserName:     user.UserName,
			UserRole:     db.UserRoleTypeName[user.UserRole],
			UserNickName: user.UserNickName,
			UserPhoneNum: user.UserPhoneNum,
			UserEmail:    user.UserEmail,
			Expires:      claims.StandardClaims.ExpiresAt,
		}, c)

	}
}

func ControlGetLoginLogList(s *services.Server) gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := checkTheAccessPermission(c, db.CONTROL, db.TRACE); err != nil {
			WithoutPermissionJSONResp(err.Error(), c)
			return
		}

		pageStr := c.Query("page")
		pageSizeStr := c.Query("pageSize")
		sortTypeStr := c.Query("sortType")

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
			ModelStruct: new(db.LoginLog),
			Page:        int32(page),
			PageSize:    int32(pageSize),
			SortType:    sortType,
			SearchIndex: make([]string, 0),
		}

		sqlRows, err := s.QueryObjectsWithPage(params)
		if err != nil {
			ServerErrorJSONResp(err.Error(), c)
			return
		}

		defer sqlRows.Close()

		resp := make([]*models.LoginLogInfo, 0)

		for sqlRows.Next() {
			var loginLog db.LoginLog
			err := s.ScanRows(sqlRows, &loginLog)
			if err != nil {
				ServerErrorJSONResp(err.Error(), c)
				return
			}

			resp = append(resp, &models.LoginLogInfo{
				UserName:  loginLog.UserName,
				LoginTime: loginLog.LoginTime,
				LoginIp:   loginLog.LoginIp,
				BaseRespInfo: models.BaseRespInfo{
					Id:        loginLog.Id,
					LastTime:  loginLog.LastTime,
					ChainTime: loginLog.ChainTime,
				},
			})
		}

		SuccessfulJSONRespWithPage(resp, len(resp), c)
	}
}
