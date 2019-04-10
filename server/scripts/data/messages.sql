DROP TABLE IF EXISTS messages;
CREATE TABLE messages
(
    id TEXT NOT NULL PRIMARY KEY,
    circle_id TEXT NOT NULL,
    sender_id TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at TEXT NOT NULL
);