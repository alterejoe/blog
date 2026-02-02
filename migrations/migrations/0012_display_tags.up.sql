CREATE TABLE blog.display_tags (
    tag_id uuid PRIMARY KEY REFERENCES blog.tags (id) ON DELETE CASCADE,
    bg_color text NOT NULL DEFAULT '',
    text_color text NOT NULL DEFAULT '',
    border_colors text[] NOT NULL DEFAULT '{}',
    bold boolean NOT NULL DEFAULT FALSE,
    italic boolean NOT NULL DEFAULT FALSE,
    underline boolean NOT NULL DEFAULT FALSE,
    font text NOT NULL DEFAULT ''
);

