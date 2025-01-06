CREATE OR REPLACE FUNCTION set_updated_at()
    RETURNS TRIGGER AS
$$
BEGIN
    OLD.updated_at = NOW();
    RETURN OLD;
END;
$$ LANGUAGE 'plpgsql'
;


CREATE TRIGGER set_updated_at
    BEFORE UPDATE
    ON public.users 
    FOR EACH ROW
EXECUTE PROCEDURE set_updated_at()
;