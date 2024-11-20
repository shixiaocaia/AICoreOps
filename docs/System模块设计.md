本设计文档涵盖了微服务系统中的菜单管理、用户管理、角色管理（使用 Casbin）、以及 API 接口管理模块的数据库设计，所有时间字段均以Unix Timestamp（时间戳）形式存储。

## 一、数据库设计

### 数据库选型

使用关系型数据库如 MySQL，因为它在企业级应用中具有良好的性能和扩展性。

### 表结构设计

#### 菜单管理 (`menus` 表)

| 字段名         | 数据类型         | 约束                          | 说明                   |
|-------------|--------------|-----------------------------|----------------------|
| id          | BIGINT       | PRIMARY KEY, AUTO_INCREMENT | 菜单ID                 |
| name        | VARCHAR(100) | NOT NULL                    | 菜单名称                 |
| parent_id   | BIGINT       | FOREIGN KEY (id) NULL       | 父菜单ID，顶级菜单为NULL      |
| path        | VARCHAR(255) | NOT NULL, UNIQUE            | 路由路径                 |
| component   | VARCHAR(255) |                             | 组件路径                 |
| icon        | VARCHAR(50)  |                             | 菜单图标                 |
| sort_order  | INT          | DEFAULT 0                   | 排序顺序                 |
| route_name  | VARCHAR(100) | UNIQUE                      | 前端路由名称               |
| hidden      | TINYINT(1)   | DEFAULT 0                   | 是否隐藏菜单（0:否, 1:是）     |
| create_time | BIGINT       | DEFAULT 0                   | 创建时间（Unix Timestamp） |
| update_time | BIGINT       | DEFAULT 0                   | 更新时间（Unix Timestamp） |
| is_deleted  | TINYINT(1)   | DEFAULT 0                   | 软删除标志（0:否, 1:是）      |

```sql
CREATE TABLE `menus`
(
    `id`          BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '菜单ID',
    `name`        VARCHAR(100) NOT NULL COMMENT '菜单名称',
    `parent_id`   BIGINT NULL COMMENT '父菜单ID，顶级菜单为NULL',
    `path`        VARCHAR(255) NOT NULL COMMENT '路由路径',
    `component`   VARCHAR(255) COMMENT '组件路径',
    `icon`        VARCHAR(50) COMMENT '菜单图标',
    `sort_order`  INT    DEFAULT 0 COMMENT '排序顺序',
    `route_name`  VARCHAR(100) UNIQUE COMMENT '前端路由名称',
    `hidden`      TINYINT(1) DEFAULT 0 COMMENT '是否隐藏菜单（0:否, 1:是）',
    `create_time` BIGINT DEFAULT 0 COMMENT '创建时间',
    `update_time` BIGINT DEFAULT 0 COMMENT '更新时间',
    `is_deleted`  TINYINT(1) DEFAULT 0 COMMENT '软删除标志（0:否, 1:是）',
    FOREIGN KEY (`parent_id`) REFERENCES `menus` (`id`) ON DELETE SET NULL,
    INDEX (`parent_id`),
    INDEX (`sort_order`),
    UNIQUE KEY `unique_path` (`path`),
    UNIQUE KEY `unique_route_name` (`route_name`)
) ENGINE=InnoDB COMMENT='菜单管理表';
```

#### 用户管理 (`users` 表)

| 字段名             | 数据类型         | 约束                                      | 说明                       |
|-----------------|--------------|-----------------------------------------|--------------------------|
| id              | BIGINT       | PRIMARY KEY, AUTO_INCREMENT             | 用户ID                     |
| username        | VARCHAR(50)  | UNIQUE, NOT NULL                        | 用户名                      |
| password        | VARCHAR(255) | NOT NULL                                | 密码（加密存储）                 |
| password_salt   | VARCHAR(255) |                                         | 密码盐                      |
| email           | VARCHAR(100) | UNIQUE                                  | 邮箱                       |
| phone           | VARCHAR(20)  | UNIQUE                                  | 电话                       |
| nickname        | VARCHAR(50)  |                                         | 昵称                       |
| avatar          | VARCHAR(255) |                                         | 用户头像URL                  |
| status          | TINYINT      | DEFAULT 1, CHECK (status IN (1, 2, 3))  | 用户状态（1:活跃, 2:不活跃, 3:被封禁） |
| last_login_time | BIGINT       | DEFAULT 0                               | 最后登录时间（Unix Timestamp）   |
| create_time     | BIGINT       | DEFAULT 0                               | 创建时间（Unix Timestamp）     |
| update_time     | BIGINT       | DEFAULT 0                               | 更新时间（Unix Timestamp）     |
| is_deleted      | TINYINT      | DEFAULT 0, CHECK (is_deleted IN (0, 1)) | 软删除标志（0:否, 1:是）          |

```sql
CREATE TABLE `users`
(
    `id`              BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '用户ID',
    `username`        VARCHAR(50) UNIQUE NOT NULL COMMENT '用户名',
    `password`        VARCHAR(255)       NOT NULL COMMENT '加密后的密码',
    `password_salt`   VARCHAR(255) COMMENT '密码盐',
    `email`           VARCHAR(100) UNIQUE COMMENT '用户邮箱',
    `phone`           VARCHAR(20) UNIQUE COMMENT '用户电话',
    `nickname`        VARCHAR(50) COMMENT '用户昵称',
    `avatar`          VARCHAR(255) COMMENT '用户头像URL',
    `status`          TINYINT DEFAULT 1 CHECK (`status` IN (1, 2, 3)) COMMENT '用户状态（1:活跃, 2:不活跃, 3:被封禁）',
    `last_login_time` BIGINT  DEFAULT 0 COMMENT '最后登录时间',
    `create_time`     BIGINT  DEFAULT 0 COMMENT '创建时间',
    `update_time`     BIGINT  DEFAULT 0 COMMENT '更新时间',
    `is_deleted`      TINYINT DEFAULT 0 CHECK (`is_deleted` IN (0, 1)) COMMENT '软删除标志（0:否, 1:是）',
    INDEX (`username`),
    INDEX (`email`),
    INDEX (`phone`)
) ENGINE=InnoDB COMMENT='用户管理表';
```

#### 角色管理 (`roles` 表)

| 字段名         | 数据类型         | 约束                                                        | 说明                         |
|-------------|--------------|-----------------------------------------------------------|----------------------------|
| id          | BIGINT       | PRIMARY KEY, AUTO_INCREMENT                               | 角色ID                       |
| name        | VARCHAR(50)  | UNIQUE, NOT NULL                                          | 角色名称                       |
| description | VARCHAR(255) |                                                           | 角色描述                       |
| role_type   | VARCHAR(50)  | NOT NULL, CHECK (role_type IN (‘ADMIN’, ‘USER’, ‘GUEST’)) | 角色类型（如 ADMIN, USER, GUEST） |
| is_default  | TINYINT      | DEFAULT 0, CHECK (is_default IN (0, 1))                   | 是否为默认角色（0:否, 1:是）          |
| create_time | BIGINT       | DEFAULT 0                                                 | 创建时间（Unix Timestamp）       |
| update_time | BIGINT       | DEFAULT 0                                                 | 更新时间（Unix Timestamp）       |
| is_deleted  | TINYINT      | DEFAULT 0, CHECK (is_deleted IN (0, 1))                   | 软删除标志（0:否, 1:是）            |

```sql
CREATE TABLE `roles`
(
    `id`          BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '角色ID',
    `name`        VARCHAR(50) UNIQUE NOT NULL COMMENT '角色名称',
    `description` VARCHAR(255) COMMENT '角色描述',
    `role_type`   VARCHAR(50)        NOT NULL CHECK (`role_type` IN ('ADMIN', 'USER', 'GUEST')) COMMENT '角色类型（如 ADMIN, USER, GUEST）',
    `is_default`  TINYINT DEFAULT 0 CHECK (`is_default` IN (0, 1)) COMMENT '是否为默认角色（0:否, 1:是）',
    `create_time` BIGINT  DEFAULT 0 COMMENT '创建时间',
    `update_time` BIGINT  DEFAULT 0 COMMENT '更新时间',
    `is_deleted`  TINYINT DEFAULT 0 CHECK (`is_deleted` IN (0, 1)) COMMENT '软删除标志（0:否, 1:是）',
    INDEX (`name`)
) ENGINE=InnoDB COMMENT='角色管理表';
```

#### 用户角色关联 (`user_roles` 表)

| 字段名         | 数据类型   | 约束                     | 说明                   |
|-------------|--------|------------------------|----------------------|
| user_id     | BIGINT | FOREIGN KEY (users.id) | 用户ID                 |
| role_id     | BIGINT | FOREIGN KEY (roles.id) | 角色ID                 |
| create_time | BIGINT | DEFAULT 0              | 创建时间（Unix Timestamp） |

```sql
CREATE TABLE `user_roles`
(
    `user_id`     BIGINT NOT NULL COMMENT '用户ID',
    `role_id`     BIGINT NOT NULL COMMENT '角色ID',
    `create_time` BIGINT DEFAULT 0 COMMENT '创建时间',
    PRIMARY KEY (`user_id`, `role_id`),
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
    FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE,
    INDEX (`role_id`),
    INDEX (`user_id`)
) ENGINE=InnoDB COMMENT='用户角色关联表';
```

#### Casbin 权限管理 (`casbin_rules` 表)

Casbin 通常使用一个表来存储策略规则(此表为casbin自动生成)，默认表结构如下：

| 字段名   | 数据类型         | 约束          | 说明            |
|-------|--------------|-------------|---------------|
| id    | BIGINT       | PRIMARY KEY | 规则ID          |
| ptype | VARCHAR(100) | NOT NULL    | 策略类型（如 p, g）  |
| v0    | VARCHAR(100) |             | 主体（如角色名）      |
| v1    | VARCHAR(100) |             | 对象（如资源）       |
| v2    | VARCHAR(100) |             | 操作（如请求方法）     |
| v3    | VARCHAR(100) |             | 额外字段（可选，根据需求） |
| v4    | VARCHAR(100) |             | 额外字段（可选，根据需求） |

```sql
CREATE TABLE `casbin_rules`
(
    `id`    BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '规则ID',
    `ptype` VARCHAR(100) NOT NULL COMMENT '策略类型（如 p, g）',
    `v0`    VARCHAR(100) COMMENT '主体（如角色名）',
    `v1`    VARCHAR(100) COMMENT '对象（如资源）',
    `v2`    VARCHAR(100) COMMENT '操作（如请求方法）',
    `v3`    VARCHAR(100) COMMENT '额外字段（可选，根据需求）',
    `v4`    VARCHAR(100) COMMENT '额外字段（可选，根据需求）',
    INDEX (`ptype`),
    INDEX (`v0`),
    INDEX (`v1`),
    INDEX (`v2`),
    INDEX (`v3`),
    INDEX (`v4`)
) ENGINE=InnoDB COMMENT='Casbin 权限管理表';
```

#### API 接口管理 (`apis` 表)

| 字段名         | 数据类型         | 约束                                                                         | 说明                                          |
|-------------|--------------|----------------------------------------------------------------------------|---------------------------------------------|
| id          | BIGINT       | PRIMARY KEY, AUTO_INCREMENT                                                | API ID                                      |
| name        | VARCHAR(100) | NOT NULL                                                                   | API名称                                       |
| path        | VARCHAR(255) | NOT NULL                                                                   | API路径                                       |
| method      | VARCHAR(10)  | NOT NULL                                                                   | 请求方法（GET, POST 等）                           |
| description | VARCHAR(255) |                                                                            | API描述                                       |
| version     | VARCHAR(20)  | DEFAULT ‘v1’                                                               | API版本                                       |
| category    | VARCHAR(100) | NOT NULL, CHECK (category IN (‘USER’, ‘ADMIN’, ‘PAYMENT’, ‘NOTIFICATION’)) | API分类（如 USER, ADMIN, PAYMENT, NOTIFICATION） |
| is_public   | TINYINT      | DEFAULT 0, CHECK (is_public IN (0, 1))                                     | 是否为公开 API（0:否, 1:是）                         |
| create_time | BIGINT       | DEFAULT 0                                                                  | 创建时间（Unix Timestamp）                        |
| update_time | BIGINT       | DEFAULT 0                                                                  | 更新时间（Unix Timestamp）                        |
| is_deleted  | TINYINT      | DEFAULT 0, CHECK (is_deleted IN (0, 1))                                    | 软删除标志（0:否, 1:是）                             |

```sql
CREATE TABLE `apis`
(
    `id`          BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT 'API ID',
    `name`        VARCHAR(100) NOT NULL COMMENT 'API名称',
    `path`        VARCHAR(255) NOT NULL COMMENT 'API路径',
    `method`      VARCHAR(10)  NOT NULL COMMENT '请求方法（GET, POST 等）',
    `description` VARCHAR(255) COMMENT 'API描述',
    `version`     VARCHAR(20) DEFAULT 'v1' COMMENT 'API版本',
    `category`    VARCHAR(100) NOT NULL CHECK (`category` IN ('USER', 'ADMIN', 'PAYMENT', 'NOTIFICATION')) COMMENT 'API分类（如 USER, ADMIN, PAYMENT, NOTIFICATION）',
    `is_public`   TINYINT     DEFAULT 0 CHECK (`is_public` IN (0, 1)) COMMENT '是否为公开API（0:否, 1:是）',
    `create_time` BIGINT      DEFAULT 0 COMMENT '创建时间',
    `update_time` BIGINT      DEFAULT 0 COMMENT '更新时间',
    `is_deleted`  TINYINT     DEFAULT 0 CHECK (`is_deleted` IN (0, 1)) COMMENT '软删除标志（0:否, 1:是）',
    INDEX (`path`),
    INDEX (`method`),
    INDEX (`version`),
    INDEX (`category`),
    UNIQUE KEY `unique_path_method` (`path`, `method`)
) ENGINE=InnoDB COMMENT='API 接口管理表';
```

### 表关系说明

本设计文档中的各个表存在以下关系：

+ 用户 与 角色：为多对多关系，通过 `user_roles` 表进行关联。每个用户可以拥有多个角色，每个角色也可以分配给多个用户。
+ 角色 与 权限：通过 Casbin 进行权限管理，角色与具体的权限策略关联。Casbin 负责管理角色对资源的访问权限（如操作 API
  接口等），实现灵活的权限控制。
+ 菜单：支持多级菜单结构，通过 `parent_id` 字段实现自关联。顶级菜单的 `parent_id` 为 `NULL`，而子菜单的 `parent_id`
  指向其父菜单的 `id`，以实现层级结构。
+ API 接口：不同角色可以被授权访问特定的 API 接口。通过角色和 API 之间的权限关系，系统能够控制不同角色的 API 访问权限。

## 二、Proto文件定义

本文件定义了微服务系统中的 `菜单管理`、`用户管理`、`角色管理` 以及 `API 接口管理` 模块的 proto 结构。

### 菜单管理模块

```protobuf
syntax = "proto3";

package menu;

import "google/protobuf/timestamp.proto";

option go_package = "github.com/GoSimplicity/AICoreOps/proto/menu;menu";

// 菜单服务定义
service MenuService {
  // 创建菜单
  rpc CreateMenu (CreateMenuRequest) returns (CreateMenuResponse);
  // 获取菜单详情
  rpc GetMenu (GetMenuRequest) returns (GetMenuResponse);
  // 更新菜单
  rpc UpdateMenu (UpdateMenuRequest) returns (UpdateMenuResponse);
  // 删除菜单
  rpc DeleteMenu (DeleteMenuRequest) returns (DeleteMenuResponse);
  // 列出菜单
  rpc ListMenus (ListMenusRequest) returns (ListMenusResponse);
}

// 菜单对象
message Menu {
  int64 id = 1; // 菜单ID
  string name = 2; // 名称
  int64 parent_id = 3; // 父菜单ID，根菜单为0
  string path = 4; // 路径
  string component = 5; // 组件
  string icon = 6; // 图标
  int32 sort_order = 7; // 排序顺序
  string route_name = 8; // 路由名称
  bool hidden = 9; // 是否隐藏
  google.protobuf.Timestamp create_time = 10; // 创建时间
  google.protobuf.Timestamp update_time = 11; // 更新时间
  bool is_deleted = 12; // 是否已删除
}

// 创建菜单请求
message CreateMenuRequest {
  string name = 1; // 名称
  int64 parent_id = 2; // 父菜单ID，根菜单为0
  string path = 3; // 路径
  string component = 4; // 组件
  string icon = 5; // 图标
  int32 sort_order = 6; // 排序顺序
  string route_name = 7; // 路由名称
  bool hidden = 8; // 是否隐藏
}

// 创建菜单响应
message CreateMenuResponse {
  bool success = 1; // 是否成功
  string message = 2; // 消息
  Menu menu = 3; // 创建的菜单对象
}

// 获取菜单请求
message GetMenuRequest {
  int64 id = 1; // 菜单ID
}

// 获取菜单响应
message GetMenuResponse {
  Menu menu = 1; // 菜单对象
}

// 更新菜单请求
message UpdateMenuRequest {
  Menu menu = 1; // 需要更新的菜单对象
}

// 更新菜单响应
message UpdateMenuResponse {
  bool success = 1; // 是否成功
  string message = 2; // 消息
  Menu menu = 3; // 更新后的菜单对象
}

// 删除菜单请求
message DeleteMenuRequest {
  int64 id = 1; // 菜单ID
}

// 删除菜单响应
message DeleteMenuResponse {
  bool success = 1; // 是否成功
  string message = 2; // 消息
}

// 列出菜单请求
message ListMenusRequest {
  int32 page_number = 1; // 页码
  int32 page_size = 2; // 每页数量
  string filter = 3; // 过滤条件（例如：名称、父菜单ID）
  string sort_by = 4; // 排序字段
  bool descending = 5; // 是否降序
}

// 列出菜单响应
message ListMenusResponse {
  repeated Menu menus = 1; // 菜单列表
  int32 total = 2; // 总数量
  int32 page_number = 3; // 当前页码
  int32 page_size = 4; // 每页数量
}
```

用户管理模块

```protobuf
syntax = "proto3";

package user;

import "google/protobuf/timestamp.proto";

option go_package = "github.com/GoSimplicity/AICoreOps/proto/user;user";

// 用户服务定义
service UserService {
  // 创建用户
  rpc CreateUser (CreateUserRequest) returns (CreateUserResponse);
  // 获取用户详情
  rpc GetUser (GetUserRequest) returns (GetUserResponse);
  // 更新用户
  rpc UpdateUser (UpdateUserRequest) returns (UpdateUserResponse);
  // 删除用户
  rpc DeleteUser (DeleteUserRequest) returns (DeleteUserResponse);
  // 列出用户
  rpc ListUsers (ListUsersRequest) returns (ListUsersResponse);
}

// 用户对象
message User {
  int64 id = 1; // 用户ID
  string username = 2; // 用户名
  string email = 3; // 邮箱
  string phone = 4; // 电话
  string nickname = 5; // 昵称
  string avatar = 6; // 头像URL
  repeated string roles = 12; // 用户角色列表
  UserStatus status = 7; // 用户状态
  google.protobuf.Timestamp last_login_time = 8; // 最后登录时间
  google.protobuf.Timestamp create_time = 9; // 创建时间
  google.protobuf.Timestamp update_time = 10; // 更新时间
  bool is_deleted = 11; // 是否已删除
}

// 用户状态枚举
enum UserStatus {
  STATUS_UNSPECIFIED = 0; // 未指定
  ACTIVE = 1; // 活跃
  INACTIVE = 2; // 不活跃
  BANNED = 3; // 被封禁
  // 根据需求添加更多状态
}

// 创建用户请求
message CreateUserRequest {
  string username = 1; // 用户名
  string password = 2; // 密码
  string email = 3; // 邮箱
  string phone = 4; // 电话
  string nickname = 5; // 昵称
  string avatar = 6; // 头像URL
  UserStatus status = 7; // 用户状态
  bool is_default = 8; // 是否为默认用户
}

// 创建用户响应
message CreateUserResponse {
  bool success = 1; // 是否成功
  string message = 2; // 消息
  User user = 3; // 创建的用户对象
}

// 获取用户请求
message GetUserRequest {
  int64 id = 1; // 用户ID
}

// 获取用户响应
message GetUserResponse {
  User user = 1; // 用户对象
}

// 更新用户请求
message UpdateUserRequest {
  int64 id = 1; // 用户ID
  string username = 2; // 用户名
  string email = 3; // 邮箱
  string phone = 4; // 电话
  string nickname = 5; // 昵称
  string avatar = 6; // 头像URL
  UserStatus status = 7; // 用户状态
  bool is_deleted = 8; // 是否已删除
}

// 更新用户响应
message UpdateUserResponse {
  bool success = 1; // 是否成功
  string message = 2; // 消息
  User user = 3; // 更新后的用户对象
}

// 删除用户请求
message DeleteUserRequest {
  int64 id = 1; // 用户ID
}

// 删除用户响应
message DeleteUserResponse {
  bool success = 1; // 是否成功
  string message = 2; // 消息
}

// 列出用户请求
message ListUsersRequest {
  int32 page_number = 1; // 页码
  int32 page_size = 2; // 每页数量
  string filter = 3; // 过滤条件（例如：用户名、邮箱）
  string sort_by = 4; // 排序字段
  bool descending = 5; // 是否降序
}

// 列出用户响应
message ListUsersResponse {
  repeated User users = 1; // 用户列表
  int32 total = 2; // 总数量
  int32 page_number = 3; // 当前页码
  int32 page_size = 4; // 每页数量
}
```

角色管理模块

```protobuf
syntax = "proto3";

package role;

import "google/protobuf/timestamp.proto";

option go_package = "github.com/GoSimplicity/AICoreOps/proto/role;role";

// 角色服务定义
service RoleService {
  // 创建角色
  rpc CreateRole (CreateRoleRequest) returns (CreateRoleResponse);
  // 获取角色详情
  rpc GetRole (GetRoleRequest) returns (GetRoleResponse);
  // 更新角色
  rpc UpdateRole (UpdateRoleRequest) returns (UpdateRoleResponse);
  // 删除角色
  rpc DeleteRole (DeleteRoleRequest) returns (DeleteRoleResponse);
  // 列出角色
  rpc ListRoles (ListRolesRequest) returns (ListRolesResponse);
}

// 角色对象
message Role {
  int64 id = 1; // 角色ID
  string name = 2; // 角色名称
  string description = 3; // 角色描述
  RoleType role_type = 4; // 角色类型
  bool is_default = 5; // 是否为默认角色
  google.protobuf.Timestamp create_time = 6; // 创建时间
  google.protobuf.Timestamp update_time = 7; // 更新时间
  bool is_deleted = 8; // 是否已删除
}

// 角色类型枚举
enum RoleType {
  ROLE_TYPE_UNSPECIFIED = 0; // 未指定
  ADMIN = 1; // 管理员
  USER = 2; // 普通用户
  GUEST = 3; // 访客
  // 根据需求添加更多角色类型
}

// 创建角色请求
message CreateRoleRequest {
  string name = 1; // 角色名称
  string description = 2; // 角色描述
  RoleType role_type = 3; // 角色类型
  bool is_default = 4; // 是否为默认角色
}

// 创建角色响应
message CreateRoleResponse {
  bool success = 1; // 是否成功
  string message = 2; // 消息
  Role role = 3; // 创建的角色对象
}

// 获取角色请求
message GetRoleRequest {
  int64 id = 1; // 角色ID
}

// 获取角色响应
message GetRoleResponse {
  Role role = 1; // 角色对象
}

// 更新角色请求
message UpdateRoleRequest {
  Role role = 1; // 需要更新的角色对象
}

// 更新角色响应
message UpdateRoleResponse {
  bool success = 1; // 是否成功
  string message = 2; // 消息
  Role role = 3; // 更新后的角色对象
}

// 删除角色请求
message DeleteRoleRequest {
  int64 id = 1; // 角色ID
}

// 删除角色响应
message DeleteRoleResponse {
  bool success = 1; // 是否成功
  string message = 2; // 消息
}

// 列出角色请求
message ListRolesRequest {
  int32 page_number = 1; // 页码
  int32 page_size = 2; // 每页数量
  string filter = 3; // 过滤条件（例如：名称、角色类型）
  string sort_by = 4; // 排序字段
  bool descending = 5; // 是否降序
}

// 列出角色响应
message ListRolesResponse {
  repeated Role roles = 1; // 角色列表
  int32 total = 2; // 总数量
  int32 page_number = 3; // 当前页码
  int32 page_size = 4; // 每页数量
}
```

API 接口管理模块

```protobuf
syntax = "proto3";

package api;

import "google/protobuf/timestamp.proto";

option go_package = "github.com/GoSimplicity/AICoreOps/proto/api;api";

service ApiService {
  // 创建 API
  rpc CreateApi (CreateApiRequest) returns (CreateApiResponse);
  // 获取 API 详情
  rpc GetApi (GetApiRequest) returns (GetApiResponse);
  // 更新 API
  rpc UpdateApi (UpdateApiRequest) returns (UpdateApiResponse);
  // 删除 API
  rpc DeleteApi (DeleteApiRequest) returns (DeleteApiResponse);
  // 列出 APIs
  rpc ListApis (ListApisRequest) returns (ListApisResponse);
}

// API 对象
message Api {
  int64 id = 1; // API ID
  string name = 2; // 名称
  string path = 3; // 路径
  HttpMethod method = 4; // HTTP 方法
  string description = 5; // 描述
  string version = 6; // 版本
  ApiCategory category = 7; // 分类
  bool is_public = 8; // 是否公开
  google.protobuf.Timestamp create_time = 9; // 创建时间
  google.protobuf.Timestamp update_time = 10; // 更新时间
  bool is_deleted = 11; // 是否已删除
}

// HTTP 方法枚举
enum HttpMethod {
  UNKNOWN_METHOD = 0;
  GET = 1;
  POST = 2;
  PUT = 3;
  DELETE = 4;
  PATCH = 5;
  OPTIONS = 6;
  HEAD = 7;
}

// API 分类枚举
enum ApiCategory {
  UNKNOWN_CATEGORY = 0;
  USER = 1;
  ADMIN = 2;
  PAYMENT = 3;
  NOTIFICATION = 4;
  // 根据需求添加更多分类
}

// 创建 API 请求
message CreateApiRequest {
  string name = 1; // 名称
  string path = 2; // 路径
  HttpMethod method = 3; // HTTP 方法
  string description = 4; // 描述
  string version = 5; // 版本
  ApiCategory category = 6; // 分类
  bool is_public = 7; // 是否公开
}

// 创建 API 响应
message CreateApiResponse {
  bool success = 1; // 是否成功
  string message = 2; // 消息
  Api api = 3; // 创建的 API 对象
}

// 获取 API 请求
message GetApiRequest {
  int64 id = 1; // API ID
}

// 获取 API 响应
message GetApiResponse {
  Api api = 1; // API 对象
}

// 更新 API 请求
message UpdateApiRequest {
  Api api = 1; // 需要更新的 API 对象
}

// 更新 API 响应
message UpdateApiResponse {
  bool success = 1; // 是否成功
  string message = 2; // 消息
  Api api = 3; // 更新后的 API 对象
}

// 删除 API 请求
message DeleteApiRequest {
  int64 id = 1; // API ID
}

// 删除 API 响应
message DeleteApiResponse {
  bool success = 1; // 是否成功
  string message = 2; // 消息
}

// 列出 APIs 请求
message ListApisRequest {
  int32 page_number = 1; // 页码
  int32 page_size = 2; // 每页数量
  string filter = 3; // 过滤条件（例如：名称、分类）
  string sort_by = 4; // 排序字段
  bool descending = 5; // 是否降序
}

// 列出 APIs 响应
message ListApisResponse {
  repeated Api apis = 1; // API 列表
  int32 total = 2; // 总数量
  int32 page_number = 3; // 当前页码
  int32 page_size = 4; // 每页数量
}
```

### Proto文件说明

本 Proto 文件定义了四个主要服务（MenuService、UserService、RoleService、ApiService），每个服务包含基本的 CRUD
操作接口。每个实体的消息结构遵循数据库设计中的字段定义，确保数据一致性。

