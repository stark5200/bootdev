CREATE TABLE people (
    id INTEGER,
    tag TEXT,
    name TEXT,
    age INTEGER,
    balance INTEGER,
    is_admin BOOLEAN
);

CREATE TABLE banks (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    location TEXT NOT NULL,
    routing_number INTEGER NOT NULL
);

CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    age INTEGER NOT NULL,
    country_code TEXT NOT NULL,
    username TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    is_admin BOOLEAN,
    bank_id INTEGER, 
    CONSTRAINT fk_banks
    FOREIGN KEY (bank_id)
    REFERENCES banks(id)
);

-- TEST SUITE, DON'T TOUCH BELOW THIS LINE --

SELECT sql AS users_table FROM sqlite_schema WHERE type='table' AND name='users';

