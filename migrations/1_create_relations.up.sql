CREATE TABLE IF NOT EXISTS products (
    product_id SERIAL PRIMARY KEY,
    product_name VARCHAR(100),
    product_type VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS contracts (
    contract_id SERIAL PRIMARY KEY,
    product_id int,
    revenue decimal,
    date_signed timestamp,
    FOREIGN KEY(product_id) REFERENCES products(product_id)
);

CREATE TABLE IF NOT EXISTS revenue_recognitions (
    contract_id int,
    amount decimal,
    recognized_on timestamp,
    FOREIGN KEY(contract_id) REFERENCES contracts(contract_id)
);