CREATE TABLE deliveries (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    phone VARCHAR,
    zip VARCHAR,
    city VARCHAR,
    address VARCHAR,
    region VARCHAR,
    email VARCHAR
);

CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    transaction VARCHAR,
    request_id VARCHAR,
    currency VARCHAR,
    provider VARCHAR,
    amount INTEGER,
    payment_dt INTEGER,
    bank VARCHAR,
    delivery_cost INTEGER,
    goods_total INTEGER,
    custom_fee INTEGER
);

CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    chrt_id INTEGER,
    track_number VARCHAR,
    price INTEGER,
    rid VARCHAR,
    name VARCHAR,
    sale INTEGER,
    size VARCHAR,
    total_price INTEGER,
    nm_id INTEGER,
    brand VARCHAR,
    status INTEGER
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    order_uid VARCHAR UNIQUE NOT NULL,
    track_number VARCHAR,
    entry VARCHAR,
    delivery_id INTEGER,
    payment_id INTEGER,
    items_id INTEGER,
    locale VARCHAR,
    internal_signature VARCHAR,
    customer_id VARCHAR,
    delivery_service VARCHAR,
    shardkey VARCHAR,
    sm_id INTEGER,
    date_created timestamp,
    oof_shard VARCHAR,

    FOREIGN KEY (delivery_id) REFERENCES deliveries(id),
    FOREIGN KEY (payment_id) REFERENCES payments(id),
    FOREIGN KEY (items_id) REFERENCES items(id)
);