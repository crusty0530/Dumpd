CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    username    TEXT NOT NULL UNIQUE,
    email   TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT username_is_lowercase CHECK (username = lower(username)),
    CONSTRAINT username_no_spaces CHECK (username NOT LIKE '% %'),
    CONSTRAINT username_no_at CHECK (username NOT LIKE '%@%'),
    CONSTRAINT username_valid_length CHECK (length(username) BETWEEN 3 AND 32),

    CONSTRAINT email_is_lowercase CHECK (email = lower(email)),
    CONSTRAINT email_contains_at CHECK (email LIKE '%@%')
);

CREATE TABLE sessions (
    token_hash BYTEA PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    last_used_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX sessions_user_id_idx ON sessions (user_id);