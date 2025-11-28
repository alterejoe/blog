-- name: NewUserSession :exec
INSERT INTO notes.user_sessions (user_id, last_token)
    VALUES (sqlc.arg (user_id), sqlc.arg (last_token))
ON CONFLICT (user_id)
    DO UPDATE SET
        last_token = sqlc.arg (last_token);

-- name: UpsertUserBySub :one
INSERT INTO notes.users (auth0_sub)
    VALUES ($1)
ON CONFLICT (auth0_sub)
    DO UPDATE SET
        auth0_sub = EXCLUDED.auth0_sub
    RETURNING
        id;

