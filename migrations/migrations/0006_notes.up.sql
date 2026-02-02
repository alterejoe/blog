CREATE TABLE blog.fleeting_note_type (
    id serial PRIMARY KEY,
    name text NOT NULL
);

CREATE TABLE blog.fleeting (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    date text NOT NULL,
    note_type_id int NOT NULL REFERENCES blog.fleeting_note_type (id),
    created_at timestamp NOT NULL DEFAULT now(),
    user_id uuid REFERENCES blog.users (id)
);

CREATE INDEX idx_fleeting_date ON blog.fleeting (date);

