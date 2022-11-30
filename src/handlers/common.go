package handlers

import (
	"net/http"
	"starnet-demo/src/models"

	"github.com/gin-gonic/gin"
)

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

func SuccessfulJSONRespWithPage(data interface{}, total int32, c *gin.Context) {
	resp := models.StandardRespWithPage{
		StandardResp: models.StandardResp{
			Code: models.RESP_CODE_SUCCESS,
			Data: data,
		},
		Total: total,
	}
	c.JSON(http.StatusOK, resp)
}

func ServerErrorJSONResp(c *gin.Context) {
	resp := models.StandardResp{
		Code: models.RESP_CODE_SERVER_ERROR,
		Msg:  models.RESP_MSG_SERVER_ERROR,
	}
	c.JSON(http.StatusOK, resp)
}

func ParamsTypeErrorJSONResp(c *gin.Context) {
	resp := models.StandardResp{
		Code: models.RESP_CODE_PARAMS_TYPE_ERROR,
		Msg:  models.RESP_MSG_PARAMS_TYPE_ERROR,
	}
	c.JSON(http.StatusOK, resp)
}

func ParamsMissingJSONResp(c *gin.Context) {
	resp := models.StandardResp{
		Code: models.RESP_CODE_PARAMS_MISSING,
		Msg:  models.RESP_MSG_PARAMS_MISSING,
	}
	c.JSON(http.StatusOK, resp)
}

func ParamsFormatErrorJSONResp(c *gin.Context) {
	resp := models.StandardResp{
		Code: models.RESP_CODE_PARAMS_FORMAT_ERROR,
		Msg:  models.RESP_MSG_PARAMS_FORMAT_ERROR,
	}
	c.JSON(http.StatusOK, resp)
}
