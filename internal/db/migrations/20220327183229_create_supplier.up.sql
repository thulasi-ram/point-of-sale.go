

CREATE TABLE suppliers
(
    id         BIGSERIAL PRIMARY KEY,
    name       text        NOT NULL,
    phone      text        NOT NULL default '',
    address    text        NOT NULL default '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_timestamp
    BEFORE UPDATE
    ON suppliers
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();