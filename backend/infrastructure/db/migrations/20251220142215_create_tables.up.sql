CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL
    photo_id INTEGER NOT NULL,
);

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    price INTEGER NOT NULL,
    description TEXT,
    category_id INTEGER,
    product_order INTEGER,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

CREATE TABLE IF NOT EXISTS photos (
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL,
    photo_path TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    pic_order INTEGER,
    FOREIGN KEY (product_id) REFERENCES products(id)
);