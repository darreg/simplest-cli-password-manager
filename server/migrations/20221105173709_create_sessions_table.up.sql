CREATE TABLE IF NOT EXISTS sessions (
     id UUID NOT NULL PRIMARY KEY,
     user_id UUID NOT NULL CONSTRAINT entries_users_fk REFERENCES users(id),
     login_time TIMESTAMP NOT NULL,
     last_seen_time TIMESTAMP NOT NULL
);