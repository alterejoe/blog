CREATE TABLE notes.user_sessions (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    user_id uuid NOT NULL REFERENCES notes.users (id),
    last_token text REFERENCES notes.sessions (token) ON DELETE CASCADE,
    updated_at timestamptz NOT NULL DEFAULT now(),
    UNIQUE (user_id)
);

CREATE TRIGGER user_sessions_updated_at
    BEFORE UPDATE ON notes.user_sessions
    FOR EACH ROW
    EXECUTE FUNCTION set_updated_at ();

CREATE INDEX sessions_user_idx ON notes.user_sessions (user_id);

CREATE INDEX sessions_last_seen_idx ON notes.user_sessions (updated_at);

