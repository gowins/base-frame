package constant

type ErrorCode int

func (c ErrorCode) Value() int { return int(c) }

const (
	CONFIG_PARSE_ERR_CODE ErrorCode = iota + 101000

	RESOURCE_NIL_ERR_CODE = iota + 101150
	RESOURCE_DESTRUCT_ERR_CODE

	SDK_UNIVERSAL_ERR_CODE ErrorCode = iota + 101100
	SDK_PARAM_ERR_CODE
	SDK_ENCODE_ERR_CODE
	SDK_DECODE_ERR_CODE
	SDK_SEND_ERR_CODE
	SDK_UNKNOWN_ERR_CODE

	DB_QUERY_ERR_CODE = iota + 101200
	DB_INSERT_ERR_CODE
	DB_UPDATE_ERR_CODE
	DB_DELETE_ERR_CODE

	MQ_PRODUCER_ERR_CODE = iota + 101300
	MQ_PRODUCER_SENDRESULT_ERR_CODE

	MQ_CONSUMER_ERR_CODE = iota + 101400

	MSGAPI_UNIVERSAL_ERR_CODE ErrorCode = iota + 101500
	MSGAPI_ENCODE_ERR_CODE
	MSGAPI_UNKNOW_ERR_CODE

	MSGWORKER_UNIVERSAL_ERR_CODE ErrorCode = iota + 101600
	MSGWORKER_ENCODE_ERR_CODE
	MSGWORKER_DECODE_ERR_CODE

	EVENT_UNWARP_ERR_CODE = iota + 101700
	EVENT_UNMARSHAL_ERR_CODE
	EVENT_ENCODE_ERR_CODE
	EVENT_DECODE_ERR_CODE
)

type ApiErrorCode int

func (c ApiErrorCode) Value() int { return int(c) }

// 对外 api
const (
	API_REQUEST_PARAM_BIND_ERR_CODE ApiErrorCode = iota + 101800
	API_REQUEST_PARAM_CHECK_ERR_CODE
	API_QUERY_TASK_ERR_CODE
	API_QUERY_SUBTASK_ERR_CODE
	API_QUERY_TEMPLATE_ERR_CODE
)

type ErrorContent string

func (c ErrorContent) Value() string { return string(c) }

const (
	CONFIG_ACCOUNT_NOT_FOUND         ErrorContent = "账号配置信息不存在"
	DINGTALK_WEBHOOK_EMPTY           ErrorContent = "钉钉机器人 webhook 配置不能为空"
	DINGTALK_SENDRESULT_UNABLE_PARSE ErrorContent = "钉钉机器人 webhook 配置错误"

	API_PARAM_UMARSHAL_ERR       ErrorContent = "模版参数解析错误"
	API_PARAM_BINDJSON_FAILD     ErrorContent = "request params bind json failed"
	API_PARAM_NOT_FOUND          ErrorContent = "param not found"
	API_PARAM_EXECTIME_NOT_ALLOW ErrorContent = "param[exec_time] must be within 24h"
	API_TASK_QUERY_FAILD         ErrorContent = "task 查询失败，任务已禁用或不存在"
	API_SUBTASK_QUERY_FAILD      ErrorContent = "subtask 查询失败，请在管理后台检查配置"
	API_TEMPLATE_QUERY_FAILD     ErrorContent = "template 查询失败，请在管理后台确认是否可用"

	EVENT_MARSHAL_ERROR ErrorContent = "clound event marshal error"

	DB_QUERY_ERROR ErrorContent = "db query error"

	MQ_PRODUCER_ERROR         ErrorContent = "mq producer error"
	MQ_PRODUCER_SEND_ERROR    ErrorContent = "mq producer send msg failed"
	MQ_CONSUMER_ERROR         ErrorContent = "mq consumer failed"
	MQ_CONSUMER_RECEIVE_ERROR ErrorContent = "mq consumer send msg failed"
)
