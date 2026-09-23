DROP TABLE IF EXISTS order_items, orders, products, customers, big_orders CASCADE;

CREATE TABLE customers (
  id         serial PRIMARY KEY,
  name       text NOT NULL,
  email      text NOT NULL UNIQUE,
  prefecture text NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE products (
  id       serial PRIMARY KEY,
  name     text NOT NULL,
  price    integer NOT NULL CHECK (price >= 0),
  stock    integer NOT NULL DEFAULT 0,
  category text NOT NULL
);

CREATE TABLE orders (
  id          serial PRIMARY KEY,
  customer_id integer NOT NULL REFERENCES customers(id),
  status      text NOT NULL DEFAULT 'pending',
  ordered_at  timestamptz NOT NULL
);

CREATE TABLE order_items (
  order_id   integer NOT NULL REFERENCES orders(id),
  product_id integer NOT NULL REFERENCES products(id),
  quantity   integer NOT NULL CHECK (quantity > 0),
  PRIMARY KEY (order_id, product_id)
);

INSERT INTO customers (name, email, prefecture, created_at) VALUES
  ('alice',  'alice@example.com',  '東京都',   '2026-08-01 10:00+09'),
  ('bob',    'bob@example.com',    '大阪府',   '2026-08-02 11:00+09'),
  ('carol',  'carol@example.com',  '東京都',   '2026-08-05 09:30+09'),
  ('dave',   'dave@example.com',   '福岡県',   '2026-08-10 18:00+09'),
  ('erin',   'erin@example.com',   '大阪府',   '2026-08-15 12:00+09');

INSERT INTO products (name, price, stock, category) VALUES
  ('Goで学ぶバックエンド入門', 3200, 12, 'book'),
  ('SQLアンチパターン攻略',    2800,  5, 'book'),
  ('メカニカルキーボード',     15800, 3, 'gadget'),
  ('USB-Cハブ',               4500, 20, 'gadget'),
  ('ドリップコーヒー30袋',     1980, 50, 'food'),
  ('プロテインバー12本',       1450,  0, 'food'),
  ('モニターライト',           6980,  8, 'gadget'),
  ('データベース設計の教科書', 3600,  7, 'book');

INSERT INTO orders (customer_id, status, ordered_at) VALUES
  (1, 'shipped',  '2026-09-01 10:15+09'),
  (2, 'shipped',  '2026-09-02 14:00+09'),
  (1, 'pending',  '2026-09-05 09:00+09'),
  (3, 'shipped',  '2026-09-06 20:30+09'),
  (2, 'canceled', '2026-09-08 08:45+09'),
  (4, 'pending',  '2026-09-10 16:20+09'),
  (1, 'shipped',  '2026-09-12 11:00+09');

INSERT INTO order_items (order_id, product_id, quantity) VALUES
  (1, 1, 1), (1, 5, 2),
  (2, 3, 1),
  (3, 8, 1), (3, 2, 1),
  (4, 4, 2), (4, 5, 1),
  (5, 5, 3),
  (6, 1, 1), (6, 7, 1),
  (7, 5, 4);
