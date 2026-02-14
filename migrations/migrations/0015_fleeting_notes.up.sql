-- up
ALTER TABLE blog.fleeting
    ADD COLUMN content text NOT NULL DEFAULT '';

ALTER TABLE blog.fleeting
    ALTER COLUMN updated_at TYPE timestamp
    USING updated_at::timestamp;

