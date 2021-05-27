package errorx

var (
	ParamsError         = NewErrCode(100, "参数错误")
	DBDataNotFound      = NewErrCode(101, "数据不存在")
	DBUpdateNotAffected = NewErrCode(102, "数据写入或者修改失败`")
)

type ErrCode struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewErrCode(code int, msg string) error {
	return &ErrCode{
		Code: code,
		Msg:  msg,
	}
}

func (e *ErrCode) Error() string {
	return e.Msg
}
