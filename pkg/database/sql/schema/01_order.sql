
CREATE TABLE "order" (
    order_id          SERIAL    PRIMARY KEY NOT NULL,
    market_id         TEXT      NOT NULL,
    order_price       TEXT      NOT NULL,
    order_amount      TEXT      NOT NULL,
    order_side        TEXT      NOT NULL,
    order_created     TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    external_order_id INTEGER
);

CREATE TABLE order_status (
    order_status_id        SERIAL    PRIMARY KEY NOT NULL,
    order_id               INTEGER   NOT NULL REFERENCES "order"(order_id),
    order_status_type      TEXT      NOT NULL,
    order_status_timestamp TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
