BEGIN;
CREATE TABLE product_images
(
    id          BIGINT          NOT NULL    GENERATED BY DEFAULT AS IDENTITY,
    product_id  BIGINT          NOT NULL,
    image_id    BIGINT          NOT NULL,
    
    created_at  TIMESTAMP       NOT NULL    DEFAULT (now() at time zone 'utc'),
    updated_at  TIMESTAMP       NOT NULL    DEFAULT (now() at time zone 'utc'),

    PRIMARY KEY (id),
    CONSTRAINT product_image_product_fk FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE,
    CONSTRAINT product_image_image_fk FOREIGN KEY (image_id) REFERENCES images (id) ON DELETE RESTRICT
);
COMMIT;
