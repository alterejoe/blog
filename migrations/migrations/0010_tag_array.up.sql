-- up
ALTER TABLE blog.fleeting
    DROP CONSTRAINT IF EXISTS fleeting_note_type_id_fkey;

ALTER TABLE blog.fleeting
    ALTER COLUMN tag_id DROP DEFAULT;

ALTER TABLE blog.fleeting RENAME COLUMN tag_id TO tag_ids;

ALTER TABLE blog.fleeting
    ALTER COLUMN tag_ids TYPE int[]
    USING ARRAY[tag_ids];

ALTER TABLE blog.fleeting
    ALTER COLUMN tag_ids SET DEFAULT '{}';

ALTER TABLE blog.fleeting
    ALTER COLUMN tag_ids DROP NOT NULL;

