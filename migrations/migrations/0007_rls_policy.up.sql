-- migration: create_rls_helpers
CREATE OR REPLACE FUNCTION blog.current_user_id ()
    RETURNS uuid
    LANGUAGE sql
    STABLE
    AS $$
    SELECT
        nullif (current_setting('app.current_user_id', TRUE), '')::uuid
$$;

CREATE OR REPLACE FUNCTION blog.is_owner (user_id uuid)
    RETURNS boolean
    LANGUAGE sql
    STABLE
    AS $$
    SELECT
        user_id = blog.current_user_id ()
$$;

