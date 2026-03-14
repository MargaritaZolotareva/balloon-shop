ALTER TABLE categories ADD COLUMN category_order INT;

UPDATE categories SET category_order = 1 WHERE id = 48;
UPDATE categories SET category_order = 2 WHERE id = 49;
UPDATE categories SET category_order = 3 WHERE id = 51;
UPDATE categories SET category_order = 4 WHERE id = 45;
UPDATE categories SET category_order = 5 WHERE id = 46;
UPDATE categories SET category_order = 6 WHERE id = 50;
UPDATE categories SET category_order = 7 WHERE id = 44;
UPDATE categories SET category_order = 8 WHERE id = 47;