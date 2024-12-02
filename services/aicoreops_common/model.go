/*
 * Copyright 2024 Bamboo
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * File: model.go
 */

package aicoreops_common

// BizCode 业务状态码类型
type BizCode int

const SuccessCode BizCode = 0
const BizCodeUnauthorized BizCode = 401
const BizCodeForbidden BizCode = 403

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
	r.Data = data
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
