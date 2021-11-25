CREATE TRIGGER set_timestamp
BEFORE UPDATE ON tanks
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
