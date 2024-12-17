CREATE TABLE IF NOT EXISTS products (
    id bigserial PRIMARY KEY,
    sku varchar(255) NOT NULL UNIQUE,
    product_name varchar(255) NOT NULL,
    category varchar(255) NOT NULL,
    price integer NOT NULL
);

INSERT INTO products (sku, product_name, category, price) VALUES 
('000001', 'BV Lean leather ankle boots', 'boots', 89000),
('000002', 'BV Lean leather ankle boots', 'boots', 99000),
('000003', 'Ashlington leather ankle boots', 'boots', 71000),
('000004', 'Naima embellished suede sandals', 'sandals', 79500),
('000005', 'Nathane leather sneakers', 'sneakers', 59000);

