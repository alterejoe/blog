-- down
BEGIN;
ALTER TABLE blog.fleeting
    ALTER COLUMN created_at TYPE timestamp,
    ALTER COLUMN updated_at TYPE timestamp;
ALTER TABLE blog.note_component
    ALTER COLUMN created_at TYPE timestamp,
    ALTER COLUMN updated_at TYPE timestamp;
COMMIT;

