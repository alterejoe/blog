-- name: CreateFleetingNote :one
INSERT INTO blog.fleeting (content, tag_ids, user_id)
    VALUES (@content, ARRAY (
            SELECT
                id
            FROM
                blog.tags
            WHERE
                name = ANY (@tag_names::text[])),
            @user_id)
RETURNING
    *;

-- name: GetFleetingNote :one
SELECT
    *
FROM
    blog.fleeting
WHERE
    id = @id;

-- name: GetFleetingNotesAsc :many
SELECT
    *
FROM
    blog.fleeting AS b
WHERE
    user_id = @id
ORDER BY
    b.created_at ASC;

-- name: GetFleetingNotesDesc :many
SELECT
    *
FROM
    blog.fleeting
WHERE
    user_id = @user_id
    AND deleted_at IS NULL
ORDER BY
    created_at DESC;

-- name: UpdateFleetingNoteContent :exec
UPDATE
    blog.fleeting
SET
    content = @content,
    updated_at = now()
WHERE
    id = @id
    AND deleted_at IS NULL;

-- name: SoftDeleteFleetingNote :exec
UPDATE
    blog.fleeting
SET
    deleted_at = now(),
    updated_at = now()
WHERE
    id = @id
    AND deleted_at IS NULL;

-- name: RestoreFleetingNote :exec
UPDATE
    blog.fleeting
SET
    deleted_at = NULL,
    updated_at = now()
WHERE
    id = @id
    AND deleted_at IS NOT NULL;

