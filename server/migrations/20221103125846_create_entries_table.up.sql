CREATE TABLE IF NOT EXISTS entries
(
    id UUID NOT NULL PRIMARY KEY,
    user_id UUID NOT NULL CONSTRAINT entries_users_fk REFERENCES users(id),
    type_id UUID NOT NULL CONSTRAINT entries_types_fk REFERENCES types(id),
    name VARCHAR(255) NOT NULL,
    metadata TEXT NOT NULL,
    data BYTEA NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
