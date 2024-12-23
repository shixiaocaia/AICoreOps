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
 * File: ai_logic.go
 */

package logic

import (
	"context"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_api/internal/svc"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AiLogic {
	return &AiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AskQuestion 提问
func (l *AiLogic) AskQuestion(req *types.AskQuestionRequest) (*types.AskQuestionResponse, error) {
	panic("not implemented")
}
