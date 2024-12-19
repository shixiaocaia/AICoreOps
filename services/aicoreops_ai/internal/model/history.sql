-- Active: 1734573960622@@127.0.0.1@3306@aicoreops
CREATE TABLE IF NOT EXISTS history (
    id INT AUTO_INCREMENT PRIMARY KEY,
    session_id VARCHAR(255),
    question TEXT,
    answer TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

# mock 一些对话数据
INSERT INTO history (session_id, question, answer, created_at) VALUES ('1', '你好', '你好，我是AI助手', '2024-01-01 00:00:00');
INSERT INTO history (session_id, question, answer, created_at) VALUES ('1', '你是谁', '我是AI助手，一个由深度求索（DeepSeek）公司开发的智能助手。', '2024-01-01 00:00:00'); 
INSERT INTO history (session_id, question, answer, created_at) VALUES ('1', '你有什么功能', '我可以帮助你回答问题、提供信息、执行任务等。', '2024-01-01 00:00:00');     