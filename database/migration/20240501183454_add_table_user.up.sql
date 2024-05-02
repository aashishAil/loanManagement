CREATE TABLE "public"."user"
(
    "id"                 uuid        NOT NULL DEFAULT uuid_generate_v4(),
    "name"               text        NOT NULL,
    "email"              text        NOT NULL UNIQUE,
    "encrypted_password" text        NOT NULL,
    "type"               text        NOT NULL DEFAULT 'customer',
    "created_at"         timestamp NOT NULL DEFAULT current_timestamp,

    CONSTRAINT "pk_user" PRIMARY KEY ("id")
);

CREATE INDEX "idx_user_email" ON "public"."user" ("email");