ALTER TABLE posts
ADD COLUMN category_id INTEGER,
ADD CONSTRAINT posts_category_id_fkey FOREIGN KEY (category_id) REFERENCES categories (id);
