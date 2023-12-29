CREATE TABLE "equipment" (
  "id" bigserial PRIMARY KEY,
  "model_number" varchar NOT NULL,
  "description" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "location" bigserial NOT NULL,
  "parent" bigserial,
  "powered_from" bigserial,
  "controlled_from" bigserial
);

CREATE TABLE "location" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar,
  "inside_of" bigserial
);

CREATE INDEX ON "equipment" ("model_number");

CREATE INDEX ON "equipment" ("description");

CREATE INDEX ON "location" ("name");

ALTER TABLE "equipment" ADD FOREIGN KEY ("location") REFERENCES "location" ("id");

ALTER TABLE "equipment" ADD FOREIGN KEY ("parent") REFERENCES "equipment" ("id");

ALTER TABLE "equipment" ADD FOREIGN KEY ("powered_from") REFERENCES "equipment" ("id");

ALTER TABLE "equipment" ADD FOREIGN KEY ("controlled_from") REFERENCES "equipment" ("id");

ALTER TABLE "location" ADD FOREIGN KEY ("inside_of") REFERENCES "location" ("id");
