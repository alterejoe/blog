-- down
DROP INDEX IF EXISTS blog.idx_fleeting_deleted_at;

ALTER TABLE blog.fleeting
    DROP COLUMN deleted_at;

