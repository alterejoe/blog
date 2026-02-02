-- down
ALTER TABLE blog.tags RENAME TO fleeting_note_type;

ALTER TABLE blog.fleeting RENAME COLUMN tag_id TO note_type_id;

