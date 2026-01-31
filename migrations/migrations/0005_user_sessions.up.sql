CREATE TABLE blog.user_sessions (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    user_id uuid NOT NULL REFERENCES blog.users (id),
    last_token text REFERENCES blog.sessions (token) ON DELETE CASCADE,
    updated_at timestamptz NOT NULL DEFAULT now(),
    UNIQUE (user_id)
);

CREATE TRIGGER user_sessions_updated_at
    BEFORE UPDATE ON blog.user_sessions
    FOR EACH ROW
    EXECUTE FUNCTION set_updated_at ();

CREATE INDEX sessions_user_idx ON blog.user_sessions (user_id);

CREATE INDEX sessions_last_seen_idx ON blog.user_sessions (updated_at);

