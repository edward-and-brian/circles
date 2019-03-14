CREATE TABLE messages
(
    id          TEXT    NOT NULL    PRIMARY KEY,
    chat_id     TEXT    NOT NULL,
    circle_id   TEXT    NOT NULL,
    sender_id   TEXT    NOT NULL,
    content     TEXT    NOT NULL,
    created_at  TEXT    DEFAULT     CURRENT_TIMESTAMP
);