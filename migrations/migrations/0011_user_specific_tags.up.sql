-- up
-- Add uuid column to tags
ALTER TABLE blog.tags
    ADD COLUMN new_id uuid DEFAULT gen_random_uuid ();

-- Add user_id to tags (nullable first)
ALTER TABLE blog.tags
    ADD COLUMN user_id uuid;

-- Set user_id for existing rows
UPDATE
    blog.tags
SET
    user_id = (
        SELECT
            id
        FROM
            blog.users
        LIMIT 1)
WHERE
    user_id IS NULL;

-- Now make it NOT NULL
ALTER TABLE blog.tags
    ALTER COLUMN user_id SET NOT NULL;

-- Create mapping table for old int -> new uuid
CREATE TEMP TABLE tag_id_map AS
SELECT
    id AS old_id,
    new_id
FROM
    blog.tags;

-- Update fleeting to use new UUIDs
ALTER TABLE blog.fleeting
    ADD COLUMN new_tag_ids uuid[] DEFAULT '{}';

UPDATE
    blog.fleeting f
SET
    new_tag_ids = ARRAY (
        SELECT
            m.new_id
        FROM
            unnest(f.tag_ids) AS t (tid)
            JOIN tag_id_map m ON m.old_id = t.tid);

-- Drop old columns and rename new ones
ALTER TABLE blog.fleeting
    DROP COLUMN tag_ids;

ALTER TABLE blog.fleeting RENAME COLUMN new_tag_ids TO tag_ids;

-- Swap id columns in tags
ALTER TABLE blog.tags
    DROP COLUMN id;

ALTER TABLE blog.tags RENAME COLUMN new_id TO id;

ALTER TABLE blog.tags
    ADD PRIMARY KEY (id);

