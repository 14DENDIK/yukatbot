CREATE TYPE step AS ENUM ('main');

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    telegram_id BIGINT UNIQUE NOT NULL,
    first_name VARCHAR(255) NULL,
    last_name VARCHAR(255) NULL,
    username VARCHAR(255) NULL,
    language_code char(3) NOT NULL,
    current_step step NOT NULL DEFAULT 'main'
);

CREATE TABLE commands (
    id SERIAL PRIMARY KEY,
    name VARCHAR(20) NOT NULL UNIQUE,
    text TEXT NULL
);