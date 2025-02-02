-- Active: 1734965197401@@127.0.0.1@3306@aicoreops
CREATE DATABASE AICoreOps
    CHARACTER SET utf8mb4
    COLLATE utf8mb4_general_ci;

USE AICoreOps;

DROP TABLE IF EXISTS history;
CREATE TABLE IF NOT EXISTS history (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `session_id` varchar(255) NOT NULL COMMENT '会话ID',
    `question` text NOT NULL COMMENT '问题',
    `answer` text NOT NULL COMMENT '答案',
    `created_at` bigint DEFAULT NULL COMMENT '创建时间',
    `updated_at` bigint DEFAULT NULL COMMENT '更新时间',
    `deleted_at` bigint DEFAULT '0' COMMENT '删除时间',
    PRIMARY KEY (id),
    KEY `idx_session_id` (`session_id`)
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS history_session;
CREATE TABLE IF NOT EXISTS history_session (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `user_id` bigint NOT NULL COMMENT '用户ID',
    `session_id` varchar(255) NOT NULL COMMENT '会话ID',
    `title` varchar(255) NOT NULL COMMENT '标题',
    `created_at` bigint DEFAULT NULL COMMENT '创建时间',
    `updated_at` bigint DEFAULT NULL COMMENT '更新时间',
    `deleted_at` bigint DEFAULT '0' COMMENT '删除时间',
    PRIMARY KEY (id),
    KEY `idx_user_id` (`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;