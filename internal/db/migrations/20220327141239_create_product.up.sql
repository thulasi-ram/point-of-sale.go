CREATE TABLE products
(
    id          BIGSERIAL PRIMARY KEY,
    name        text   NOT NULL,
    description text   NOT NULL default '',
    category_id BIGINT NOT NULL REFERENCES product_categories (id)
);