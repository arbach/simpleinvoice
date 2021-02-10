CREATE DATABASE simpleinvoice;

\c simpleinvoice;
CREATE TABLE IF NOT EXISTS invoices(
    id serial primary key,
    status varchar,
    description varchar,
    amount decimal,
    payment_address varchar,
    paid_amount decimal,
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp DEFAULT now() NOT NULL
);