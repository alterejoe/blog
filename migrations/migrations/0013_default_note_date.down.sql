-- down
DROP TRIGGER IF EXISTS set_updated_at ON blog.fleeting;

ALTER TABLE blog.fleeting
    ALTER COLUMN created_at DROP DEFAULT,
    ALTER COLUMN created_at DROP NOT NULL;

ALTER TABLE blog.fleeting RENAME COLUMN updated_at TO date;

ALTER TABLE blog.fleeting
    ALTER COLUMN date DROP DEFAULT,
    ALTER COLUMN date DROP NOT NULL;

DROP FUNCTION IF EXISTS blog.update_updated_at ();

