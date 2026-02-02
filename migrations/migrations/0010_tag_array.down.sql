-- down
ALTER TABLE blog.fleeting
    ALTER COLUMN tag_ids DROP DEFAULT;

ALTER TABLE blog.fleeting
    ALTER COLUMN tag_ids TYPE int
    USING COALESCE(tag_ids[1], 0);

ALTER TABLE blog.fleeting
    ALTER COLUMN tag_ids SET DEFAULT 0;

ALTER TABLE blog.fleeting
    ALTER COLUMN tag_ids SET NOT NULL;

ALTER TABLE blog.fleeting RENAME COLUMN tag_ids TO tag_id;

