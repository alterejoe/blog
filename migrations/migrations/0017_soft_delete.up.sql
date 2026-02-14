-- up
ALTER TABLE blog.fleeting
    ADD COLUMN deleted_at timestamptz;

CREATE INDEX idx_fleeting_deleted_at ON blog.fleeting (deleted_at)
WHERE
    deleted_at IS NULL;

