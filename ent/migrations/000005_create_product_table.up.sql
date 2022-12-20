BEGIN;
CREATE TABLE products
(
    id              BIGINT          NOT NULL    GENERATED BY DEFAULT AS IDENTITY,
    category_id     BIGINT          NOT NULL,
    name            VARCHAR(100)    NOT NULL,
    description     TEXT,
    amount          BIGINT          NOT NULL    DEFAULT 0,
    price           BIGINT          NOT NULL,

    created_at      TIMESTAMP       NOT NULL    DEFAULT (now() at time zone 'utc'),
    updated_at      TIMESTAMP       NOT NULL    DEFAULT (now() at time zone 'utc'),

    PRIMARY KEY (id),
    CONSTRAINT product_category_fk FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE RESTRICT
);
COMMIT;
