CREATE TABLE IF NOT EXISTS "tanks" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "make" VARCHAR(40) DEFAULT '',
  "model" VARCHAR(40) DEFAULT '',
  "name" VARCHAR(40) DEFAULT '',
  "location" VARCHAR(40) DEFAULT '',
  "capacity_measurement" VARCHAR(20) DEFAULT '',
  "capacity" FLOAT DEFAULT NULL,
  "description" VARCHAR(255) DEFAULT '',
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
