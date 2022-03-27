CREATE TABLE products
(
    id          BIGSERIAL PRIMARY KEY,
    name        text        NOT NULL,
    description text        NOT NULL default '',
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    category_id BIGINT      NOT NULL REFERENCES product_categories (id)
);

CREATE TRIGGER set_timestamp
    BEFORE UPDATE
    ON products
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();