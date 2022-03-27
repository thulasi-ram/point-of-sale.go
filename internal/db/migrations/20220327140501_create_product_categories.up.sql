CREATE TABLE product_categories
(
    id         BIGSERIAL PRIMARY KEY,
    name       TEXT        NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    parent_id  BIGINT REFERENCES product_categories (id)
);

CREATE TRIGGER set_timestamp
    BEFORE UPDATE
    ON product_categories
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();