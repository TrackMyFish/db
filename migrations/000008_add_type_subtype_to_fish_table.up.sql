ALTER TABLE fish
ADD COLUMN IF NOT EXISTS "type" VARCHAR(255) DEFAULT '',
ADD COLUMN IF NOT EXISTS "subtype" VARCHAR(255) DEFAULT '';
