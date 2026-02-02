CREATE TABLE blog.note_component (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    user_id uuid NOT NULL REFERENCES blog.users (id) ON DELETE CASCADE,
    parent_id uuid REFERENCES blog.note_component (id) ON DELETE CASCADE,
    name text NOT NULL,
    pattern text NOT NULL,
    sort_order int NOT NULL DEFAULT 0,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now(),
    UNIQUE (user_id, pattern),
    UNIQUE (user_id, parent_id, name)
);

CREATE TRIGGER set_updated_at
    BEFORE UPDATE ON blog.note_component
    FOR EACH ROW
    EXECUTE FUNCTION blog.update_updated_at ();

CREATE INDEX idx_note_component_parent ON blog.note_component (parent_id);

CREATE INDEX idx_note_component_user ON blog.note_component (user_id);

