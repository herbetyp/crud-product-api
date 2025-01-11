CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	username VARCHAR(100) NOT NULL,
	email VARCHAR(100) NOT NULL,
	password VARCHAR(255) NOT NULL,
	uid UUID DEFAULT gen_random_uuid (),
	created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
	UNIQUE (email)
);

CREATE TRIGGER set_updated_at
    BEFORE UPDATE
    ON public.users 
    FOR EACH ROW
EXECUTE PROCEDURE set_updated_at()
;