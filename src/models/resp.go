package models

type StandardResp struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type StandardRespWithPage struct {
	StandardResp
	Total int32 `json:"total"`
}

const (
	RESP_CODE_SUCCESS = 200

	RESP_CODE_SERVER_ERROR = 500

	RESP_CODE_PARAMS_TYPE_ERROR = 101

	RESP_CODE_PARAMS_MISSING = 102

	RESP_CODE_PARAMS_FORMAT_ERROR = 103
)

const (
	RESP_MSG_SERVER_ERROR = "服务出错"

	RESP_MSG_PARAMS_TYPE_ERROR = "错误的参数类型"

	RESP_MSG_PARAMS_MISSING = "缺少必要参数"

	RESP_MSG_PARAMS_FORMAT_ERROR = "参数不符合规范"
)
