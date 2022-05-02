BEGIN;

CREATE TABLE sale_orders
(
    id                  BIGSERIAL PRIMARY KEY,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    customer_id         BIGINT      NOT NULL REFERENCES customers (id),
    additional_discount MONEY       not null
);

CREATE TABLE sale_order_items
(
    id            BIGSERIAL PRIMARY KEY,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    sale_order_id BIGINT      NOT NULL REFERENCES sale_orders (id),
    product_id    BIGINT      NOT NULL REFERENCES products (id),
    quantity      NUMERIC     not null,
    amount        MONEY       not null,
    discount      MONEY       not null
);

CREATE TRIGGER set_timestamp
    BEFORE UPDATE
    ON sale_orders
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp
    BEFORE UPDATE
    ON sale_order_items
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

COMMIT;