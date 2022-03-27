CREATE TABLE product_categories
(
    id        BIGSERIAL PRIMARY KEY,
    name      TEXT NOT NULL,
    parent_id BIGINT,
    FOREIGN KEY (parent_id) REFERENCES product_categories (id)
);