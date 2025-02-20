CREATE TABLE "public"."loan"
(
    "id"                 uuid      NOT NULL DEFAULT uuid_generate_v4(),
    "user_id"            uuid      NOT NULL REFERENCES "user" (id),
    "disbursal_amount"   bigint    NOT NULL,
    "pending_amount"     bigint    NOT NULL,
    "currency"           text      NOT NULL DEFAULT 'INR',
    "term"               bigint    NOT NULL,
    "status"             text      NOT NULL DEFAULT 'PENDING',
    "disbursal_date"     timestamp NOT NULL,
    "created_at"         timestamp NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    "updated_at"         timestamp NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),

    CONSTRAINT "pk_loan" PRIMARY KEY ("id")
);

CREATE TRIGGER update_loan_modtime
    BEFORE UPDATE
    ON loan
    FOR EACH ROW
EXECUTE PROCEDURE update_modified_column();


CREATE INDEX idx_loan_user_id_status on loan (user_id, status);