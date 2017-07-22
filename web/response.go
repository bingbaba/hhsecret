package web

var (
	MSG_OK      string = "OK"
	CODE_OK     int    = 1000
	CODE_FAILED int    = 5000
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResponse(data interface{}) *Response {
	return &Response{CODE_OK, MSG_OK, data}
}

func NewResponseWithErr(err error, data interface{}) *Response {
	if err == nil {
		return NewResponse(data)
	} else {
		return &Response{CODE_FAILED, err.Error(), data}
	}
}
