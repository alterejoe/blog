-- down
ALTER TABLE blog.fleeting
    ALTER COLUMN updated_at TYPE text
    USING updated_at::text;

ALTER TABLE blog.fleeting
    DROP COLUMN content;

