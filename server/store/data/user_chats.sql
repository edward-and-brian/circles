DROP TABLE IF EXISTS user_chats;
CREATE TABLE user_chats
(
    id              TEXT    NOT NULL    PRIMARY KEY,
    user_id         TEXT    NOT NULL,
    chat_id         TEXT    NOT NULL,
    created_at      TEXT    DEFAULT     CURRENT_TIMESTAMP
);