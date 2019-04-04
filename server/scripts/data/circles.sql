DROP TABLE IF EXISTS circles;
CREATE TABLE circles
(
    id          TEXT    NOT NULL    PRIMARY KEY,
    chat_id     TEXT    NOT NULL,
    name        TEXT    NOT NULL,
    created_at  TEXT    DEFAULT     CURRENT_TIMESTAMP
);