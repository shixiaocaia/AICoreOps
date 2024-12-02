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
 * File: types.go
 */

package types

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type GetUserRequest struct {
	Id int `json:"id"`
}

type GetUserListRequest struct {
	PageNumber int `json:"page_number"`
	PageSize   int `json:"page_size"`
}

type DeleteUserRequest struct {
	Id int `json:"id"`
}

type UpdateUserRequest struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type LogoutRequest struct {
	JWTToken     string `json:"jwt_token"`
	RefreshToken string `json:"refresh_token"`
}

type GeneralResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type GeneralWithDataResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CreateApiRequest struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Version     string `json:"version"`
	Category    string `json:"category"`  // 分类
	IsPublic    bool   `json:"is_public"` // 是否公开
}

type GetApiRequest struct {
	Id int `json:"id"`
}

type UpdateApiRequest struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Version     string `json:"version"`
	Category    string `json:"category"`
	IsPublic    bool   `json:"is_public"`
}

type DeleteApiRequest struct {
	Id int `json:"id"`
}

type ListApisRequest struct {
	PageNumber int `json:"page_number"`
	PageSize   int `json:"page_size"`
}

// 菜单相关
type CreateMenuRequest struct {
	Name      string `json:"name"`
	Path      string `json:"path"`
	ParentId  int    `json:"parent_id"`  // 父级ID
	Component string `json:"component"`  // 组件
	Icon      string `json:"icon"`       // 图标
	SortOrder int    `json:"sort_order"` // 排序
	RouteName string `json:"route_name"` // 路由名称
	Hidden    bool   `json:"hidden"`     // 是否隐藏
}

type GetMenuRequest struct {
	Id int `json:"id"`
}

type UpdateMenuRequest struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	ParentId  int    `json:"parent_id"`
	Component string `json:"component"`
	Icon      string `json:"icon"`
	SortOrder int    `json:"sort_order"`
	RouteName string `json:"route_name"`
	Hidden    bool   `json:"hidden"`
}

type DeleteMenuRequest struct {
	Id int `json:"id"`
}

type ListMenusRequest struct {
	PageNumber int  `json:"page_number"`
	PageSize   int  `json:"page_size"`
	IsTree     bool `json:"is_tree"` // 是否树形结构
}

// 角色相关
type CreateRoleRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	RoleType    string `json:"role_type"`  // 角色类型
	IsDefault   bool   `json:"is_default"` // 是否默认角色
}

type GetRoleRequest struct {
	Id int `json:"id"`
}

type UpdateRoleRequest struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	RoleType    string `json:"role_type"`
	IsDefault   bool   `json:"is_default"`
}

type DeleteRoleRequest struct {
	Id int `json:"id"`
}

type ListRolesRequest struct {
	PageNumber int `json:"page_number"`
	PageSize   int `json:"page_size"`
}

type AssignPermissionsRequest struct {
	RoleId  int   `json:"role_id"`
	MenuIds []int `json:"menu_ids"`
	ApiIds  []int `json:"api_ids"`
}

type AssignRoleToUserRequest struct {
	UserId  int   `json:"user_id"`
	RoleIds []int `json:"role_ids"`
}

type RemoveRoleFromUserRequest struct {
	UserId int `json:"user_id"`
	RoleId int `json:"role_id"`
}

// 移除指定用户的所有权限
type RemoveUserPermissionsRequest struct {
	UserId int `json:"user_id"`
}
