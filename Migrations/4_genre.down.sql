begin;
ALTER TABLE genre drop CONSTRAINT genre_name_unique;
commit;