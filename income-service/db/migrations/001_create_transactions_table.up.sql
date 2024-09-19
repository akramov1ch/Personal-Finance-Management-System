CREATE TABLE IF NOT EXISTS transactions (
    transaction_id SERIAL PRIMARY KEY,
    user_id INT,
    type VARCHAR(50),
    amount NUMERIC(20, 2),
    currency VARCHAR(10),
    category VARCHAR(50),
    date TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);
