package aicoreops_common

// BizCode 业务状态码类型
type BizCode int

const SuccessCode BizCode = 0

type Result struct {
	Code    BizCode `json:"code"`           // 业务状态码
	Message string  `json:"message"`        // 响应信息
	Data    any     `json:"data,omitempty"` // 响应数据
}

func NewResultResponse() *Result {
	return &Result{}
}

// SetFailResponse 设置失败响应
func (r *Result) SetFailResponse(code BizCode, msg string) {
	r.Code = code
	r.Message = msg
}

// SetSuccessResponse 设置成功响应
func (r *Result) SetSuccessResponse(data any) {
	r.Code = SuccessCode
	r.Message = "success"
	if data != nil {
		r.Data = data
	}
}

// HandleResponse 处理响应结果
func (r *Result) HandleResponse(data any, err error) *Result {
	if err != nil {
		r.SetFailResponse(500, err.Error())
		return r
	}
	r.SetSuccessResponse(data)

	return r
}
