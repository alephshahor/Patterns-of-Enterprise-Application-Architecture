CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    product_name VARCHAR(100),
    product_type VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS contracts (
    ID SERIAL PRIMARY KEY,
    CONSTRAINT fk_product_id
        FOREIGN KEY(product_id)
            REFERENCES products(id)
);

CREATE TABLE IF NOT EXISTS revenue_recognitions (
    CONSTRAINT fk_contract_id
        FOREIGN KEY(contract_id)
            REFERENCES contracts(id),
    amount decimal,
    recognized_on timestamp,
);