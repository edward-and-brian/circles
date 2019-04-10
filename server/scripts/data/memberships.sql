DROP TABLE IF EXISTS memberships;
CREATE TABLE memberships
(
    id TEXT NOT NULL PRIMARY KEY,
    user_id TEXT NOT NULL,
    chat_id TEXT NOT NULL,
    created_at TEXT NOT NULL
);