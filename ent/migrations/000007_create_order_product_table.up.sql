BEGIN;
CREATE TABLE order_products
(
    id          BIGINT          NOT NULL    GENERATED BY DEFAULT AS IDENTITY,
    order_id    BIGINT          NOT NULL,
    product_id  BIGINT          NOT NULL,
    amount      BIGINT          NOT NULL,
    price       BIGINT          NOT NULL,

    created_at  TIMESTAMP       NOT NULL    DEFAULT (now() at time zone 'utc'),
    updated_at  TIMESTAMP       NOT NULL    DEFAULT (now() at time zone 'utc'),

    PRIMARY KEY (id),
    CONSTRAINT order_product_order_fk FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE RESTRICT,
    CONSTRAINT order_product_product_fk FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE RESTRICT,
    CONSTRAINT order_product_order_id_product_id_unique UNIQUE (order_id, product_id)
);
COMMIT;
