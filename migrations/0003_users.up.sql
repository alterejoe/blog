CREATE TABLE "notes".users (
    id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid (),
    auth0_sub text UNIQUE, -- unique creates index
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE TRIGGER users_updated_at
    BEFORE UPDATE ON "notes".users
    FOR EACH ROW
    EXECUTE FUNCTION set_updated_at ();

