CREATE TABLE IF NOT EXISTS notifapp.users (
    id BIGSERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,
    full_name VARCHAR(100) NOT NULL CHECK (
        char_length(full_name) BETWEEN 3 AND 100
    ),
    password_hash TEXT NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE CHECK (
        char_length(email) BETWEEN 5 and 100
        AND email LIKE '%@%.%'
    ),
    phone_number VARCHAR(15) UNIQUE CHECK (
        char_length(phone_number) BETWEEN 10 AND 15
        AND phone_number ~ '^\+[0-9]+$'
    ),
    telegram VARCHAR(33) UNIQUE CHECK (
        char_length(telegram) BETWEEN 6 AND 33
        AND telegram ~ '^\@[A-Za-z0-9_]+$'
    )
);