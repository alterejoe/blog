-- rename blog.fleeting_note_type to blog.tags
ALTER TABLE blog.fleeting_note_type RENAME TO tags;

-- change note_type_id to tag_id
ALTER TABLE blog.fleeting RENAME COLUMN note_type_id TO tag_id;

