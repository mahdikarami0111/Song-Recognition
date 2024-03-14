CREATE TABLE "requests" (
  "id" bigserial PRIMARY KEY,
  "email" varchar NOT NULL,
  "status" varchar NOT NULL,
  "songid" varchar NOT NULL DEFAULT ''
);

CREATE INDEX ON "requests" ("email");

CREATE INDEX ON "requests" ("status");

CREATE INDEX ON "requests" ("songid");
