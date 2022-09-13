package ginhelper

const (
	OK             = 200
	SpeedLimit     = -2
	ServerInterval = 500
)

var (
	CodeMsgMap = map[int]string{
		SpeedLimit:     "服务器正忙，请稍后再试",
		ServerInterval: "服务器错误，请联系客服",
	}
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
