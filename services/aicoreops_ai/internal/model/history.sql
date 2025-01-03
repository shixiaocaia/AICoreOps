-- Active: 1734965197401@@127.0.0.1@3306@aicoreops
CREATE DATABASE AICoreOps
    CHARACTER SET utf8mb4
    COLLATE utf8mb4_general_ci;

USE AICoreOps;
CREATE TABLE IF NOT EXISTS history (
    id bigint AUTO_INCREMENT,
    session_id varchar(255) NOT NULL,
    question text NOT NULL,
    answer text NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'history table';

CREATE TABLE IF NOT EXISTS history_session (
    id bigint AUTO_INCREMENT,
    user_id bigint NOT NULL,
    session_id varchar(255) NOT NULL,
    title varchar(255) NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'history session table';

# goctl model mysql ddl --src history.sql --dir .