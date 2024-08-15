CREATE TABLE users (
    id SERIAL PRIMARY KEY NOT NULL,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE categories (
    id SERIAL PRIMARY KEY NOT NULL,
    user_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

ALTER TABLE categories
ADD CONSTRAINT fk_user
    FOREIGN KEY (user_id) REFERENCES users(id);

CREATE TABLE accounts (
    id SERIAL PRIMARY KEY NOT NULL,
    user_id INT NOT NULL,
    category_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    value INTEGER NOT NULL,
    date DATE NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

ALTER TABLE accounts
ADD CONSTRAINT fk_user_accounts
    FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE accounts
ADD CONSTRAINT fk_category
    FOREIGN KEY (category_id) REFERENCES categories(id);
