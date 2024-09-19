CREATE TABLE IF NOT EXISTS budgets (
    budget_id SERIAL PRIMARY KEY,
    user_id INT,
    amount NUMERIC(20, 2),
    spent NUMERIC(20, 2),
    currency VARCHAR(10),
    category VARCHAR(50),
    date TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);
