CREATE TABLE user_settings
(
    id          TEXT    NOT NULL    PRIMARY KEY,
    user_id     TEXT    NOT NULL,
    created_at  TEXT    DEFAULT     CURRENT_TIMESTAMP
);