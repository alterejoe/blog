-- down
DROP TRIGGER IF EXISTS set_updated_at ON blog.note_component;

DROP INDEX IF EXISTS blog.idx_note_component_parent;

DROP INDEX IF EXISTS blog.idx_note_component_user;

DROP TABLE IF EXISTS blog.note_component;

