CREATE TABLE "accounts" (
  "id" UUID NOT NULL UNIQUE,
  "name" VARCHAR(255) NOT NULL,
  "email" VARCHAR(255) NOT NULL UNIQUE,
  "password" VARCHAR(255) NOT NULL,
  "avatar_url" VARCHAR(255) NOT NULL,
  "created_at" TIMESTAMP NOT NULL,

  PRIMARY KEY ("id")
);