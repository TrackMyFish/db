CREATE TRIGGER set_timestamp
BEFORE UPDATE ON tank_statistics
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
