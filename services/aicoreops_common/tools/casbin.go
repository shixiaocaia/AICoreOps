package tools

import (
	"errors"
	"fmt"

	"github.com/casbin/casbin/v2"
)

var (
	ErrNilEnforcer = errors.New("enforcer不能为空")
	ErrInvalidUser = errors.New("无效的用户ID") 
	ErrInvalidPath = errors.New("path和method不能为空")
)

// CheckPermission 检查用户是否有权限访问指定的资源
func CheckPermission(enforcer *casbin.Enforcer, userID int64, path, method string) (bool, error) {
	// 参数校验
	if enforcer == nil {
		return false, ErrNilEnforcer
	}

	if userID <= 0 {
		return false, ErrInvalidUser
	}

	if path == "" || method == "" {
		return false, ErrInvalidPath
	}

	// 将userID转换为字符串
	user := fmt.Sprintf("%d", userID)

	// 检查用户直接权限
	allowed, err := enforcer.Enforce(user, path, method)
	if err != nil {
		return false, fmt.Errorf("检查用户权限失败: %w", err)
	}

	if allowed {
		return true, nil
	}

	// 获取用户的所有角色
	roles, err := enforcer.GetRolesForUser(user)
	if err != nil {
		return false, fmt.Errorf("获取用户角色失败: %w", err)
	}

	// 检查用户角色权限
	for _, role := range roles {
		allowed, err = enforcer.Enforce(role, path, method)
		if err != nil {
			return false, fmt.Errorf("检查角色[%s]权限失败: %w", role, err)
		}
		
		if allowed {
			return true, nil
		}
	}

	return false, nil
}
