BEGIN;
CREATE TYPE order_status AS ENUM (
    CREATED,
    ACCEPTED,
    CANCALLED,
);

CREATE TABLE orders
(
    id          BIGINT          NOT NULL    GENERATED BY DEFAULT AS IDENTITY,
    user_id     BIGINT          NOT NULL,
    status      order_status    NOT NULL    DEFAULT(CREATED),

    created_at  TIMESTAMP       NOT NULL    DEFAULT (now() at time zone 'utc'),
    updated_at  TIMESTAMP       NOT NULL    DEFAULT (now() at time zone 'utc'),

    PRIMARY KEY (id),
    CONSTRAINT order_user_fk FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE RESTRICT
);
COMMIT;
