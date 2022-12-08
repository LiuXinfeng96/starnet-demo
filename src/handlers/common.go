package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"starnet-demo/src/db"
	"starnet-demo/src/models"
	"starnet-demo/src/services"

	"chainmaker.org/chainmaker/pb-go/v2/common"
	"github.com/gin-gonic/gin"
)

var PHONE_NUM_REGEXP = regexp.MustCompile(`1[3456789]\d{9}`)

var EMAIL_REFEXP = regexp.MustCompile(`^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		method := c.Request.Method
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,Content-Type")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			c.Header("Access-Control-Max-Age", "172800")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func SuccessfulJSONResp(data interface{}, c *gin.Context) {
	resp := models.StandardResp{
		Code: models.RESP_CODE_SUCCESS,
		Data: data,
	}
	c.JSON(http.StatusOK, resp)
}

func SuccessfulJSONRespWithPage(data interface{}, total int, c *gin.Context) {
	resp := models.StandardRespWithPage{
		StandardResp: models.StandardResp{
			Code: models.RESP_CODE_SUCCESS,
			Data: data,
		},
		Total: int32(total),
	}
	c.JSON(http.StatusOK, resp)
}

func ServerErrorJSONResp(err string, c *gin.Context) {
	resp := models.StandardResp{
		Code: models.RESP_CODE_SERVER_ERROR,
		Msg:  models.RESP_MSG_SERVER_ERROR,
		Data: err,
	}
	c.JSON(http.StatusOK, resp)
}

func ParamsTypeErrorJSONResp(err string, c *gin.Context) {
	resp := models.StandardResp{
		Code: models.RESP_CODE_PARAMS_TYPE_ERROR,
		Msg:  models.RESP_MSG_PARAMS_TYPE_ERROR,
		Data: err,
	}
	c.JSON(http.StatusOK, resp)
}

func ParamsMissingJSONResp(err string, c *gin.Context) {
	resp := models.StandardResp{
		Code: models.RESP_CODE_PARAMS_MISSING,
		Msg:  models.RESP_MSG_PARAMS_MISSING,
		Data: err,
	}
	c.JSON(http.StatusOK, resp)
}

func ParamsFormatErrorJSONResp(err string, c *gin.Context) {
	resp := models.StandardResp{
		Code: models.RESP_CODE_PARAMS_FORMAT_ERROR,
		Msg:  models.RESP_MSG_PARAMS_FORMAT_ERROR,
		Data: err,
	}
	c.JSON(http.StatusOK, resp)
}

func ParamsValueJSONResp(err string, c *gin.Context) {
	resp := models.StandardResp{
		Code: models.RESP_CODE_PARAMS_VALUE_ERROR,
		Msg:  models.RESP_MSG_PARAMS_VALUE_ERROR,
		Data: err,
	}
	c.JSON(http.StatusOK, resp)
}

func WithoutPermissionJSONResp(err string, c *gin.Context) {
	resp := models.StandardResp{
		Code: models.RESP_CODE_WITHOUT_PERMISSION,
		Msg:  models.RESP_MSG_WITHOUT_PERMISSION,
		Data: err,
	}
	c.JSON(http.StatusOK, resp)
}

func UniqueIndexJSONResp(err string, c *gin.Context) {
	resp := models.StandardResp{
		Code: models.RESP_CODE_UNIQUE_INDEX,
		Msg:  models.RESP_MSG_UNIQUE_INDEX,
		Data: err,
	}
	c.JSON(http.StatusOK, resp)
}

func NotExistJSONResp(err string, c *gin.Context) {
	resp := models.StandardResp{
		Code: models.RESP_CODE_NOT_EXIST,
		Msg:  models.RESP_MSG_NOT_EXIST,
		Data: err,
	}
	c.JSON(http.StatusOK, resp)
}

func PwdErrorJSONResp(err string, c *gin.Context) {
	resp := models.StandardResp{
		Code: models.RESP_CODE_PWD_ERROR,
		Msg:  models.RESP_MSG_PWD_ERROR,
		Data: err,
	}
	c.JSON(http.StatusOK, resp)
}

func NotInChainJSONResp(err string, c *gin.Context) {
	resp := models.StandardResp{
		Code: models.RESP_CODE_NOT_IN_CHAIN,
		Msg:  models.RESP_MSG_NOT_IN_CHAIN,
		Data: err,
	}
	c.JSON(http.StatusOK, resp)
}

func PutChainFailJSONResp(err string, c *gin.Context) {
	resp := models.StandardResp{
		Code: models.RESP_CODE_CHAIN_PUT_FAIL,
		Msg:  models.RESP_MSG_CHAIN_PUT_FAIL,
		Data: err,
	}
	c.JSON(http.StatusOK, resp)
}

func isStringRequiredParamsEmpty(params ...string) error {
	for _, p := range params {
		if len(p) == 0 {
			err := errors.New("the required param is empty")
			return err
		}
	}
	return nil
}

func checkThePhoneNum(num string) error {
	if !PHONE_NUM_REGEXP.MatchString(num) {
		return errors.New("does not conform to the format of the mobile phone number")
	}
	return nil
}

func checkTheEmail(email string) error {
	if !EMAIL_REFEXP.MatchString(email) {
		return errors.New("does not conform to the email format")
	}
	return nil
}

func checkTheAccessPermission(c *gin.Context, dbRole ...db.UserRoleType) error {

	token, ok := c.Get("token")
	if !ok {
		return errors.New("get token from context failed")
	}

	tokenClaims, ok := token.(*services.MyClaims)
	if !ok {
		return errors.New("token type error")
	}

	roleType, ok := db.UserRoleTypeValue[tokenClaims.Role]

	if !ok {
		return errors.New("the user role not found")
	}

	for _, v := range dbRole {
		if roleType == v {
			return nil
		}
	}

	return errors.New("access without permission")
}

func GetBlockChainFiledFromResp(cr *common.ContractResult) (db.BlockChainField, error) {
	var resp models.ContractResp
	err := json.Unmarshal(cr.Result, &resp)
	if err != nil {
		return db.BlockChainField{}, err
	}
	var bcFiled db.BlockChainField
	bcFiled.TxId = resp.TxId
	bcFiled.BlockHeight = resp.BlockHeight
	bcFiled.ChainTime = resp.ChainTime
	return bcFiled, nil
}
