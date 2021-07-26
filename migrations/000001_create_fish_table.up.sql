CREATE TABLE IF NOT EXISTS "fish" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "genus" VARCHAR(255) DEFAULT '',
  "species" VARCHAR(255) DEFAULT '',
  "common_name" VARCHAR(255) DEFAULT '',
  "name" VARCHAR(255) DEFAULT '',
  "color" VARCHAR(255) DEFAULT '',
  "gender" VARCHAR(255) DEFAULT '',
  "purchase_date" VARCHAR(255) DEFAULT '',
  "ecosystem_name" VARCHAR(255) DEFAULT '',
  "ecosystem_type" VARCHAR(255) DEFAULT '',
  "ecosystem_location" VARCHAR(255) DEFAULT '',
  "salinity" VARCHAR(255) DEFAULT '',
  "climate" VARCHAR(255) DEFAULT '',
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
