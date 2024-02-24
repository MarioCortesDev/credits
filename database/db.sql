DROP TABLE IF EXISTS credits_assignments;

CREATE TABLE credits_assignments (
    id SERIAL PRIMARY KEY,
    investment INT NOT NULL,
    count_300 INT NOT NULL,
    count_500 INT NOT NULL,
    count_700 INT NOT NULL,
    successful BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);