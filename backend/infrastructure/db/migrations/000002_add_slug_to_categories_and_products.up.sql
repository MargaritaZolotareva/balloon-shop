ALTER TABLE categories ADD COLUMN slug VARCHAR(255);
ALTER TABLE categories ADD CONSTRAINT categories_slug_unique UNIQUE (slug);

ALTER TABLE products ADD COLUMN slug VARCHAR(255);
ALTER TABLE products ADD CONSTRAINT products_slug_unique UNIQUE (slug);