CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name text NOT NULL,
    unit text NOT NULL
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    recipient text NOT NULL,
    shipping_date date NOT NULL,
    comment text,
    is_submitted boolean NOT NULL
);

CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    product_id integer NOT NULL REFERENCES products(id),
    rq numeric NOT NULL,
    aq numeric,
    bn text,
    comment text,
    is_submitted boolean NOT NULL,
    order_id integer NOT NULL REFERENCES orders(id)
);
