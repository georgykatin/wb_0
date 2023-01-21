CREATE TABLE IF NOT EXISTS order
(
    order_uid varchar PRIMARY KEY NOT NULL ,
    track_number varchar,
    entry varchar,
    locale varchar,
    internal_signature varchar,
    customer_id varchar,
    delivery_service varchar,
    shardkey varchar,
    sm_id bigint,
    data_created timestamp,
    oof_shard varchar
);
CREATE TABLE IF NOT EXISTS delivery
(
    order_uid varchar PRIMARY KEY NOT NULL REFERENCES order(order_uid) ON DELETE CASCADE ,
    name varchar ,
    phone varchar,
    zip varchar,
    city varchar,
    adress varchar ,
    region varchar ,
    email varchar
);
CREATE TABLE IF NOT EXISTS payment
(
    order_uid varchar PRIMARY KEY NOT NULL REFERENCES order(order_uid) ON DELETE CASCADE ,
    transaction varchar ,
    request_id varchar,
    currency varchar,
    provider varchar,
    amount bigint,
    payment_dt bigint,
    bank varchar,
    delivery_cost bigint,
    goods_total bigint,
    custom_fee bigint
);
CREATE TABLE IF NOT EXISTS items
(
    order_uid varchar PRIMARY KEY NOT NULL REFERENCES order(order_uid) ON DELETE CASCADE ,
    track_number varchar,
    price bigint,
    rid varchar,
    name varchar,
    sale bigint,
    size varchar,
    total_price bigint,
    nm_id bigint,
    brand varchar,
    status bigint
);