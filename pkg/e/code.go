package e

const (
	Success                    = 200
	Error                      = 500
	InvalidParams              = 400
	ErrorAuthToken             = 600
	ErrorAuthCheckTokenTimeout = 601
)

var MsgFlags = map[int]string{
	Success:                    "ok",
	Error:                      "fall",
	InvalidParams:              "参数请求错误",
	ErrorAuthToken:             "认证错误",
	ErrorAuthCheckTokenTimeout: "认证过期",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[Error]
}
