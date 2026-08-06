-- +goose Up
CREATE TABLE sessions (
    id BIGSERIAL PRIMARY KEY,

    user_id BIGINT NOT NULL
        REFERENCES users(id)
        ON DELETE CASCADE,

    token_hash BYTEA NOT NULL,

    user_agent TEXT,
    ip_address INET,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    expires_at TIMESTAMPTZ NOT NULL,
    last_used_at TIMESTAMPTZ NOT NULL
);

CREATE INDEX idx_sessions_user_id
ON sessions(user_id);

CREATE INDEX idx_sessions_token_hash
ON sessions(token_hash);

CREATE INDEX idx_sessions_expires_at
ON sessions(expires_at);

-- +goose Down
DROP TABLE sessions;
