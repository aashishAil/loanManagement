CREATE TABLE "public"."schedule_repayment"
(
    "id"               uuid        NOT NULL DEFAULT uuid_generate_v4(),
    "loan_id"          uuid        NOT NULL REFERENCES "loan" ("id"),
    "scheduled_amount" bigint      NOT NULL,
    "pending_amount"   bigint      NOT NULL,
    "currency"         text        NOT NULL DEFAULT 'INR',
    "status"           text        NOT NULL DEFAULT 'PENDING',
    "scheduled_date"   date        NOT NULL,
    "created_at"       timestamptz NOT NULL DEFAULT current_timestamp,
    "updated_at"       timestamptz NOT NULL DEFAULT current_timestamp,

    CONSTRAINT "pk_schedule_repayment" PRIMARY KEY ("id")
);

CREATE TRIGGER update_schedule_repayment_modtime
    BEFORE UPDATE
    ON schedule_repayment
    FOR EACH ROW
EXECUTE PROCEDURE update_modified_column();

CREATE INDEX "idx_schedule_payment_loan_id"
    ON "public"."schedule_repayment" ("loan_id");


