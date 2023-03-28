package base

type ErrCode int32

type ECode struct {
	Msg  string
	Code ErrCode
}

func NewECode(Code ErrCode, msg string) ECode {
	return ECode{
		Code: Code,
		Msg:  msg,
	}
}

func (e ECode) Error() string {
	return e.Msg
}

func (e ECode) SetMsg(msg string) {
	e.Msg = msg
}
