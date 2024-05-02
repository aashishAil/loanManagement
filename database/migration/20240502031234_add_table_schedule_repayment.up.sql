CREATE TABLE "public"."schedule_repayment"
(
    "id"               uuid        NOT NULL DEFAULT uuid_generate_v4(),
    "loan_id"          uuid        NOT NULL REFERENCES "loan" ("id"),
    "scheduled_amount" bigint      NOT NULL,
    "pending_amount"   bigint      NOT NULL,
    "currency"         text        NOT NULL DEFAULT 'INR',
    "status"           text        NOT NULL DEFAULT 'PENDING',
    "scheduled_date"   timestamp        NOT NULL,
    "created_at"       timestamp NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    "updated_at"       timestamp NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),

    CONSTRAINT "pk_schedule_repayment" PRIMARY KEY ("id")
);

CREATE TRIGGER update_schedule_repayment_modtime
    BEFORE UPDATE
    ON schedule_repayment
    FOR EACH ROW
EXECUTE PROCEDURE update_modified_column();

CREATE INDEX "idx_schedule_repayment_loan_id"
    ON "public"."schedule_repayment" ("loan_id");


