DROP TABLE IF EXISTS circles;
CREATE TABLE circles
(
    id TEXT NOT NULL PRIMARY KEY,
    chat_id TEXT NOT NULL,
    name TEXT NOT NULL,
    last_message_content TEXT,
    last_message_at TEXT,
    created_at TEXT NOT NULL
);