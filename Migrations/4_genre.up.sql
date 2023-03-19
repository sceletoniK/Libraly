begin;
ALTER TABLE genre ADD CONSTRAINT genre_name_unique UNIQUE (name);
commit;