package errcode

var ErrMap map[int]string

const (
	SUCCESS       = 0
	SERVER_ERROR  = 100001
	PARAM_INVALID = 100002
)

func init() {
	ErrMap = make(map[int]string)
	ErrMap[SUCCESS] = "success"
	ErrMap[SERVER_ERROR] = "server error"
	ErrMap[PARAM_INVALID] = "param invalid"
}
