-- up
-- Reusable trigger function (use for any table)
CREATE OR REPLACE FUNCTION blog.update_updated_at ()
    RETURNS TRIGGER
    AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$
LANGUAGE plpgsql;

-- Rename date to updated_at
ALTER TABLE blog.fleeting RENAME COLUMN date TO updated_at;

-- Set defaults
ALTER TABLE blog.fleeting
    ALTER COLUMN updated_at SET DEFAULT NOW(),
    ALTER COLUMN updated_at SET NOT NULL;

ALTER TABLE blog.fleeting
    ALTER COLUMN created_at SET DEFAULT NOW(),
    ALTER COLUMN created_at SET NOT NULL;

-- Attach trigger to fleeting
CREATE TRIGGER set_updated_at
    BEFORE UPDATE ON blog.fleeting
    FOR EACH ROW
    EXECUTE FUNCTION blog.update_updated_at ();

