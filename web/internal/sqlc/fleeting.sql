-- name: CreateFleetingNote :one
INSERT INTO blog.fleeting (tag_id, user_id)
    VALUES (@date, (
            SELECT
                id
            FROM
                blog.tags
            WHERE
                name = @tag_name), @user_id)
RETURNING
    *;

