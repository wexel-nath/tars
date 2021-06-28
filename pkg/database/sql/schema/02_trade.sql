
CREATE TABLE trade (
    trade_id             SERIAL    PRIMARY KEY NOT NULL,
    order_id             INTEGER   NOT NULL REFERENCES "order"(order_id),
    trade_price          TEXT      NOT NULL,
    trade_side           TEXT      NOT NULL,
    trade_fee            TEXT      NOT NULL,
    trade_liquidity_type TEXT      NOT NULL,
    trade_timestamp      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    external_order_id    INTEGER,
    external_trade_id    INTEGER
);
