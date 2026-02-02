-- down
-- Add integer id back to tags
ALTER TABLE blog.tags
    ADD COLUMN new_id serial;

-- Create mapping table
CREATE TEMP TABLE tag_id_map AS
SELECT
    id AS old_uuid,
    new_id
FROM
    blog.tags;

-- Update fleeting back to int[]
ALTER TABLE blog.fleeting
    ADD COLUMN new_tag_ids int[] DEFAULT '{}';

UPDATE
    blog.fleeting f
SET
    new_tag_ids = ARRAY (
        SELECT
            m.new_id
        FROM
            unnest(f.tag_ids) AS t (tid)
            JOIN tag_id_map m ON m.old_uuid = t.tid);

-- Drop and rename
ALTER TABLE blog.fleeting
    DROP COLUMN tag_ids;

ALTER TABLE blog.fleeting RENAME COLUMN new_tag_ids TO tag_ids;

ALTER TABLE blog.tags
    DROP CONSTRAINT tags_pkey;

ALTER TABLE blog.tags
    DROP COLUMN id;

ALTER TABLE blog.tags RENAME COLUMN new_id TO id;

ALTER TABLE blog.tags
    ADD PRIMARY KEY (id);

-- Remove user_id
ALTER TABLE blog.tags
    DROP COLUMN user_id;

