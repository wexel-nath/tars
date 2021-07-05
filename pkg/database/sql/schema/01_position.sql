
CREATE TABLE position (
    position_id       SERIAL          PRIMARY KEY NOT NULL,
    market_id         TEXT            NOT NULL,
    position_price    NUMERIC(32, 16) NOT NULL,
    position_amount   NUMERIC(32, 16) NOT NULL,
    position_type     TEXT            NOT NULL,
    position_status   TEXT            NOT NULL,
    position_created  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    position_updated  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    external_order_id TEXT
);
