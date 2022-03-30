BEGIN;

CREATE TABLE purchase_orders
(
    id                  BIGSERIAL PRIMARY KEY,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    supplier_id         BIGINT      NOT NULL REFERENCES suppliers (id),
    additional_discount MONEY       not null
);

CREATE TABLE purchase_order_items
(
    id                BIGSERIAL PRIMARY KEY,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    purchase_order_id BIGINT      NOT NULL REFERENCES purchase_orders (id),
    product_id        BIGINT      NOT NULL REFERENCES products (id),
    quantity          int         not null,
    amount            MONEY       not null,
    discount          MONEY       not null
);

CREATE TRIGGER set_timestamp
    BEFORE UPDATE
    ON purchase_orders
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp
    BEFORE UPDATE
    ON purchase_order_items
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

COMMIT;