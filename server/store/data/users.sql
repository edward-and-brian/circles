CREATE TABLE users
(
    id              TEXT    NOT NULL    PRIMARY KEY,
    name            TEXT    NOT NULL,
    phone_number    TEXT    NOT NULL    UNIQUE,
    display_name    TEXT    NOT NULL    UNIQUE,
    created_at      TEXT    DEFAULT     CURRENT_TIMESTAMP
);