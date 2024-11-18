本设计文档涵盖了微服务系统中的菜单管理、用户管理、角色管理（使用 Casbin）、以及 API 接口管理模块的数据库设计，所有时间字段均以 Unix Timestamp（时间戳）形式存储。

## 一、数据库设计
### 数据库选型
使用关系型数据库如 MySQL，因为它在企业级应用中具有良好的性能和扩展性。

### 表结构设计
#### 菜单管理 (`menus` 表)
| 字段名 | 数据类型 | 约束 | 说明 |
| --- | --- | --- | --- |
| id | BIGINT | PRIMARY KEY, AUTO_INCREMENT | 菜单ID |
| name | VARCHAR(100) | NOT NULL | 菜单名称 |
| parent_id | BIGINT | FOREIGN KEY (`id`) NULL | 父菜单ID，顶级菜单为NULL |
| path | VARCHAR(255) | NOT NULL | 路由路径 |
| component | VARCHAR(255) |  | 组件路径 |
| icon | VARCHAR(50) |  | 菜单图标 |
| sort_order | INT | DEFAULT 0 | 排序顺序 |
| route_name | VARCHAR(100) |  | 前端路由名称 |
| hidden | BOOLEAN | DEFAULT FALSE | 是否隐藏菜单 |
| create_time | BIGINT | DEFAULT 0 | 创建时间（Unix Timestamp） |
| update_time | BIGINT | DEFAULT 0 | 更新时间（Unix Timestamp） |
| is_deleted | BOOLEAN | DEFAULT FALSE | 软删除标志 |


```sql
CREATE TABLE `menus` (
  `id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '菜单ID',
  `name` VARCHAR(100) NOT NULL COMMENT '菜单名称',
  `parent_id` BIGINT NULL COMMENT '父菜单ID，顶级菜单为NULL',
  `path` VARCHAR(255) NOT NULL COMMENT '路由路径',
  `component` VARCHAR(255) COMMENT '组件路径',
  `icon` VARCHAR(50) COMMENT '菜单图标',
  `sort_order` INT DEFAULT 0 COMMENT '排序顺序',
  `route_name` VARCHAR(100) COMMENT '前端路由名称',
  `hidden` BOOLEAN DEFAULT FALSE COMMENT '是否隐藏菜单',
  `create_time` BIGINT DEFAULT 0 COMMENT '创建时间',
  `update_time` BIGINT DEFAULT 0 COMMENT '更新时间',
  `is_deleted` BOOLEAN DEFAULT FALSE COMMENT '软删除标志',
  FOREIGN KEY (`parent_id`) REFERENCES `menus`(`id`) ON DELETE SET NULL,
  INDEX (`parent_id`),
  INDEX (`sort_order`)
) ENGINE=InnoDB COMMENT='菜单管理表';
```

#### 用户管理 (`users` 表)
| 字段名 | 数据类型 | 约束 | 说明 |
| --- | --- | --- | --- |
| id | BIGINT | PRIMARY KEY, AUTO_INCREMENT | 用户ID |
| username | VARCHAR(50) | UNIQUE, NOT NULL | 用户名 |
| password | VARCHAR(255) | NOT NULL | 密码（加密存储） |
| password_salt | VARCHAR(255) |  | 密码盐 |
| email | VARCHAR(100) | UNIQUE | 邮箱 |
| phone | VARCHAR(20) | UNIQUE | 电话 |
| nickname | VARCHAR(50) |  | 昵称 |
| avatar | VARCHAR(255) |  | 用户头像URL |
| status | TINYINT | DEFAULT 1 | 用户状态（1:正常, 0:禁用） |
| last_login_time | BIGINT | DEFAULT 0 | 最后登录时间（Unix Timestamp） |
| create_time | BIGINT | DEFAULT 0 | 创建时间（Unix Timestamp） |
| update_time | BIGINT | DEFAULT 0 | 更新时间（Unix Timestamp） |
| is_deleted | BOOLEAN | DEFAULT FALSE | 软删除标志 |


```sql
CREATE TABLE `users` (
  `id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '用户ID',
  `username` VARCHAR(50) UNIQUE NOT NULL COMMENT '用户名',
  `password` VARCHAR(255) NOT NULL COMMENT '加密后的密码',
  `password_salt` VARCHAR(255) COMMENT '密码盐',
  `email` VARCHAR(100) UNIQUE COMMENT '用户邮箱',
  `phone` VARCHAR(20) UNIQUE COMMENT '用户电话',
  `nickname` VARCHAR(50) COMMENT '用户昵称',
  `avatar` VARCHAR(255) COMMENT '用户头像URL',
  `status` TINYINT DEFAULT 1 COMMENT '用户状态（1:正常, 0:禁用）',
  `last_login_time` BIGINT DEFAULT 0 COMMENT '最后登录时间',
  `create_time` BIGINT DEFAULT 0 COMMENT '创建时间',
  `update_time` BIGINT DEFAULT 0 COMMENT '更新时间',
  `is_deleted` BOOLEAN DEFAULT FALSE COMMENT '软删除标志',
  INDEX (`username`),
  INDEX (`email`),
  INDEX (`phone`)
) ENGINE=InnoDB COMMENT='用户管理表';
```

#### 角色管理 (`roles` 表)
| 字段名 | 数据类型 | 约束 | 说明 |
| --- | --- | --- | --- |
| id | BIGINT | PRIMARY KEY, AUTO_INCREMENT | 角色ID |
| name | VARCHAR(50) | UNIQUE, NOT NULL | 角色名称 |
| description | VARCHAR(255) |  | 角色描述 |
| role_type | VARCHAR(50) |  | 角色类型 |
| is_default | BOOLEAN | DEFAULT FALSE | 是否为默认角色 |
| create_time | BIGINT | DEFAULT 0 | 创建时间（Unix Timestamp） |


```sql
CREATE TABLE `roles` (
  `id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '角色ID',
  `name` VARCHAR(50) UNIQUE NOT NULL COMMENT '角色名称',
  `description` VARCHAR(255) COMMENT '角色描述',
  `role_type` VARCHAR(50) COMMENT '角色类型',
  `is_default` BOOLEAN DEFAULT FALSE COMMENT '是否为默认角色',
  `create_time` BIGINT DEFAULT 0 COMMENT '创建时间',
  `update_time` BIGINT DEFAULT 0 COMMENT '更新时间',
  `is_deleted` BOOLEAN DEFAULT FALSE COMMENT '软删除标志',
  INDEX (`name`)
) ENGINE=InnoDB COMMENT='角色管理表';
```

#### 用户角色关联 (`user_roles` 表)
| 字段名 | 数据类型 | 约束 | 说明 |
| --- | --- | --- | --- |
| user_id | BIGINT | FOREIGN KEY (`users.id`), PRIMARY KEY | 用户ID |
| role_id | BIGINT | FOREIGN KEY (`roles.id`), PRIMARY KEY | 角色ID |
| create_time | BIGINT | DEFAULT 0 | 创建时间（Unix Timestamp） |


```sql
CREATE TABLE `user_roles` (
  `user_id` BIGINT NOT NULL COMMENT '用户ID',
  `role_id` BIGINT NOT NULL COMMENT '角色ID',
  `create_time` BIGINT DEFAULT 0 COMMENT '创建时间',
  PRIMARY KEY (`user_id`, `role_id`),
  FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`role_id`) REFERENCES `roles`(`id`) ON DELETE CASCADE,
  INDEX (`role_id`)
) ENGINE=InnoDB COMMENT='用户角色关联表';
```

#### Casbin 权限管理 (`casbin_rules` 表)
Casbin 通常使用一个表来存储策略规则(此表为casbin自动生成)，默认表结构如下：

| 字段名 | 数据类型 | 约束 | 说明 |
| --- | --- | --- | --- |
| id | BIGINT | PRIMARY KEY | 规则ID |
| ptype | VARCHAR(100) | NOT NULL | 策略类型（如 `p`, `g`） |
| v0 | VARCHAR(100) |  | 主体（如角色名） |
| v1 | VARCHAR(100) |  | 对象（如资源） |
| v2 | VARCHAR(100) |  | 操作（如请求方法） |


#### API 接口管理 (`apis` 表)
| 字段名 | 数据类型 | 约束 | 说明 |
| --- | --- | --- | --- |
| id | BIGINT | PRIMARY KEY, AUTO_INCREMENT | API ID |
| name | VARCHAR(100) | NOT NULL | API名称 |
| path | VARCHAR(255) | NOT NULL | API路径 |
| method | VARCHAR(10) | NOT NULL | 请求方法（GET, POST 等） |
| description | VARCHAR(255) |  | API描述 |
| version | VARCHAR(20) | DEFAULT 'v1' | API版本 |
| category | VARCHAR(100) |  | API分类 |
| is_public | BOOLEAN | DEFAULT FALSE | 是否为公开 API |
| create_time | BIGINT | DEFAULT 0 | 创建时间（Unix Timestamp） |
| update_time | BIGINT | DEFAULT 0 | 更新时间（Unix Timestamp） |
| is_deleted | BOOLEAN | DEFAULT FALSE | 软删除标志 |


```sql
CREATE TABLE `apis` (
  `id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT 'API ID',
  `name` VARCHAR(100) NOT NULL COMMENT 'API名称',
  `path` VARCHAR(255) NOT NULL COMMENT 'API路径',
  `method` VARCHAR(10) NOT NULL COMMENT '请求方法（GET, POST 等）',
  `description` VARCHAR(255) COMMENT 'API描述',
  `version` VARCHAR(20) DEFAULT 'v1' COMMENT 'API版本',
  `category` VARCHAR(100) COMMENT 'API分类',
  `is_public` BOOLEAN DEFAULT FALSE COMMENT '是否为公开API',
  `create_time` BIGINT DEFAULT 0 COMMENT '创建时间',
  `update_time` BIGINT DEFAULT 0 COMMENT '更新时间',
  `is_deleted` BOOLEAN DEFAULT FALSE COMMENT '软删除标志',
  INDEX (`path`),
  INDEX (`method`),
  INDEX (`version`)
) ENGINE=InnoDB COMMENT='API 接口管理表';
```

### 表关系说明
本设计文档中的各个表存在以下关系：

+ 用户 与 角色：为多对多关系，通过 `user_roles` 表进行关联。每个用户可以拥有多个角色，每个角色也可以分配给多个用户。
+ 角色 与 权限：通过 Casbin 进行权限管理，角色与具体的权限策略关联。Casbin 负责管理角色对资源的访问权限（如操作 API 接口等），实现灵活的权限控制。
+ 菜单：支持多级菜单结构，通过 `parent_id` 字段实现自关联。顶级菜单的 `parent_id` 为 `NULL`，而子菜单的 `parent_id` 指向其父菜单的 `id`，以实现层级结构。
+ API 接口：不同角色可以被授权访问特定的 API 接口。通过角色和 API 之间的权限关系，系统能够控制不同角色的 API 访问权限。

## 二、Proto文件定义
本文件定义了微服务系统中的 `菜单管理`、`用户管理`、`角色管理` 以及 `API 接口管理` 模块的 proto 结构。

### 菜单管理模块
```protobuf
syntax = "proto3";

package menu;

service MenuService {
    rpc CreateMenu (Menu) returns (MenuResponse);
    rpc GetMenu (MenuId) returns (Menu);
    rpc UpdateMenu (Menu) returns (MenuResponse);
    rpc DeleteMenu (MenuId) returns (MenuResponse);
    rpc ListMenus (MenuListRequest) returns (MenuListResponse);
}

message Menu {
    int64 id = 1;
    string name = 2;
    int64 parent_id = 3;
    string path = 4;
    string component = 5;
    string icon = 6;
    int32 sort_order = 7;
    string route_name = 8;
    bool hidden = 9;
    int64 create_time = 10;
    int64 update_time = 11;
    bool is_deleted = 12;
}

message MenuId {
    int64 id = 1;
}

message MenuResponse {
    bool success = 1;
    string message = 2;
}

message MenuListRequest {}

message MenuListResponse {
    repeated Menu menus = 1;
}
```

用户管理模块

```protobuf
syntax = "proto3";

package user;

service UserService {
    rpc CreateUser (User) returns (UserResponse);
    rpc GetUser (UserId) returns (User);
    rpc UpdateUser (User) returns (UserResponse);
    rpc DeleteUser (UserId) returns (UserResponse);
    rpc ListUsers (UserListRequest) returns (UserListResponse);
}

message User {
    int64 id = 1;
    string username = 2;
    string password = 3;
    string password_salt = 4;
    string email = 5;
    string phone = 6;
    string nickname = 7;
    string avatar = 8;
    int32 status = 9;
    int64 last_login_time = 10;
    int64 create_time = 11;
    int64 update_time = 12;
    bool is_deleted = 13;
}

message UserId {
    int64 id = 1;
}

message UserResponse {
    bool success = 1;
    string message = 2;
}

message UserListRequest {}

message UserListResponse {
    repeated User users = 1;
}
```

角色管理模块

```protobuf
syntax = "proto3";

package role;

service RoleService {
    rpc CreateRole (Role) returns (RoleResponse);
    rpc GetRole (RoleId) returns (Role);
    rpc UpdateRole (Role) returns (RoleResponse);
    rpc DeleteRole (RoleId) returns (RoleResponse);
    rpc ListRoles (RoleListRequest) returns (RoleListResponse);
}

message Role {
    int64 id = 1;
    string name = 2;
    string description = 3;
    string role_type = 4;
    bool is_default = 5;
    int64 create_time = 6;
    int64 update_time = 7;
    bool is_deleted = 8;
}

message RoleId {
    int64 id = 1;
}

message RoleResponse {
    bool success = 1;
    string message = 2;
}

message RoleListRequest {}

message RoleListResponse {
    repeated Role roles = 1;
}
```

API 接口管理模块

```protobuf
syntax = "proto3";

package api;

service ApiService {
    rpc CreateApi (Api) returns (ApiResponse);
    rpc GetApi (ApiId) returns (Api);
    rpc UpdateApi (Api) returns (ApiResponse);
    rpc DeleteApi (ApiId) returns (ApiResponse);
    rpc ListApis (ApiListRequest) returns (ApiListResponse);
}

message Api {
    int64 id = 1;
    string name = 2;
    string path = 3;
    string method = 4;
    string description = 5;
    string version = 6;
    string category = 7;
    bool is_public = 8;
    int64 create_time = 9;
    int64 update_time = 10;
    bool is_deleted = 11;
}

message ApiId {
    int64 id = 1;
}

message ApiResponse {
    bool success = 1;
    string message = 2;
}

message ApiListRequest {}

message ApiListResponse {
    repeated Api apis = 1;
}
```

### Proto文件说明
本 Proto 文件定义了四个主要服务（MenuService、UserService、RoleService、ApiService），每个服务包含基本的 CRUD 操作接口。每个实体的消息结构遵循数据库设计中的字段定义，确保数据一致性。

