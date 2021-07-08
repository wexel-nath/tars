
CREATE TABLE run (
    run_id       SERIAL    PRIMARY KEY NOT NULL,
    run_config   TEXT      NOT NULL,
    run_started  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    run_finished TIMESTAMP WITH TIME ZONE
);

CREATE TABLE position (
    position_id       SERIAL          PRIMARY KEY NOT NULL,
    run_id            INTEGER         NOT NULL REFERENCES run(run_id),
    market_id         TEXT            NOT NULL,
    position_type     TEXT            NOT NULL,
    position_status   TEXT            NOT NULL,
    position_created  TIMESTAMP       WITH TIME ZONE NOT NULL DEFAULT NOW(),
    amount            NUMERIC(32, 16) NOT NULL,
    open_price        NUMERIC(32, 16) NOT NULL,
    close_price       NUMERIC(32, 16),
    open_fee          NUMERIC(32, 16),
    close_fee         NUMERIC(32, 16),
    open_order_id     TEXT,
    close_order_id    TEXT
);
