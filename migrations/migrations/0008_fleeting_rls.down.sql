-- down
DROP POLICY fleeting_owner ON blog.fleeting;

ALTER TABLE blog.fleeting DISABLE ROW LEVEL SECURITY;

