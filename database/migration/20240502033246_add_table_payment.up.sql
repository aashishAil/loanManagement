CREATE TABLE "public"."payment"
(
    "id"         uuid      NOT NULL DEFAULT uuid_generate_v4(),
    "loan_id"    uuid      NOT NULL REFERENCES "loan" ("id"),
    "amount"     bigint    NOT NULL,
    "currency"   text      NOT NULL DEFAULT 'INR',
    "created_at" timestamp NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),

    CONSTRAINT "pk_payment" PRIMARY KEY ("id")
);

CREATE INDEX "idx_payment_loan_id"
    ON "public"."payment" ("loan_id");