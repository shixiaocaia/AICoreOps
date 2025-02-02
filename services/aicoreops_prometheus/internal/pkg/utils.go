package pkg

import "context"

// HandleList 处理搜索或获取所有记录
func HandleList[T any](ctx context.Context, search *string, searchFunc func(ctx context.Context, name string) ([]*T, error), listFunc func(ctx context.Context) ([]*T, error)) ([]*T, error) {
	if search != nil && *search != "" {
		return searchFunc(ctx, *search)
	}
	return listFunc(ctx)
}
