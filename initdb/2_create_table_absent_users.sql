CREATE TABLE absent_users
(
    id            UUID PRIMARY KEY,
    user_id       UUID NOT NULL,
    ip_address    VARCHAR(255) NOT NULL,
    latitude      FLOAT NOT NULL,
    longitude     FLOAT NOT NULL,
    status        VARCHAR(255) NOT NULL,
    created_at    TIMESTAMPTZ  NOT NULL DEFAULT current_timestamp,
    updated_at    TIMESTAMPTZ  NOT NULL DEFAULT current_timestamp,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id)
);
