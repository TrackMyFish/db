CREATE TRIGGER set_timestamp
BEFORE UPDATE ON fish
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
