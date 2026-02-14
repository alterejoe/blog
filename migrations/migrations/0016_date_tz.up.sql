BEGIN;
ALTER TABLE blog.fleeting
    ALTER COLUMN created_at TYPE timestamptz
    USING created_at AT TIME ZONE 'UTC',
    ALTER COLUMN updated_at TYPE timestamptz
    USING updated_at AT TIME ZONE 'UTC';
ALTER TABLE blog.note_component
    ALTER COLUMN created_at TYPE timestamptz
    USING created_at AT TIME ZONE 'UTC',
    ALTER COLUMN updated_at TYPE timestamptz
    USING updated_at AT TIME ZONE 'UT';
COMMIT;

