CREATE TABLE "public"."loan"
(
    "id"                 uuid        NOT NULL DEFAULT uuid_generate_v4(),
    "user_id"            uuid        NOT NULL REFERENCES "user" (id),
    "disbursal_amount"   bigint      NOT NULL,
    "pending_amount"     bigint      NOT NULL,
    "weekly_installment" bigint      NOT NULL,
    "currency"           text        NOT NULL DEFAULT 'INR',
    "term"               bigint      NOT NULL,
    "status"             text        NOT NULL DEFAULT 'PENDING',
    "disbursal_date"     timestamptz NOT NULL,
    "created_at"         timestamptz NOT NULL DEFAULT current_timestamp,
    "updated_at"         timestamptz NOT NULL DEFAULT current_timestamp,

    CONSTRAINT "pk_loan" PRIMARY KEY ("id")
);

CREATE TRIGGER update_loan_modtime
    BEFORE UPDATE
    ON loan
    FOR EACH ROW
EXECUTE PROCEDURE update_modified_column();


CREATE INDEX idx_loan_user_id_status on loan(user_id, status);