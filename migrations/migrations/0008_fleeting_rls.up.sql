ALTER TABLE blog.fleeting ENABLE ROW LEVEL SECURITY;

CREATE POLICY fleeting_owner ON blog.fleeting
    USING (blog.is_owner (user_id))
    WITH CHECK (blog.is_owner (user_id));

