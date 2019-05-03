DROP TABLE IF EXISTS chats;
CREATE TABLE chats
(
    id TEXT NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    last_circle_name TEXT,
    last_message_at TEXT,
    created_at TEXT NOT NULL
);