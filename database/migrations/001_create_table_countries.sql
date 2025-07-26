-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE countries (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    iso2 CHAR(2) NOT NULL,
    iso3 CHAR(3) NOT NULL,
    capital VARCHAR(100),
    region VARCHAR(100),
    language VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +migrate StatementEnd


-- +migrate Down
-- +migrate StatementBegin
DROP TABLE IF EXISTS countries;
-- +migrate StatementEnd
