CREATE DATABASE readinglist;

CREATE ROLE readinglist WITH LOGIN PASSWORD '';

CREATE TABLE IF NOT EXISTS books (
    id bigserial PRIMARY KEY,  
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title text NOT NULL,
    published integer NOT NULL,
    pages integer NOT NULL,
    genres text[] NOT NULL,
    version integer NOT NULL DEFAULT 1
);

ALTER TABLE books ADD CONSTRAINT books_pages_check CHECK (pages >= 0);
-- TODO update constraint to be a vaild int not date range
ALTER TABLE books ADD CONSTRAINT books_published_check CHECK (published BETWEEN 1450 AND date_part('published', now()));

ALTER TABLE books ADD CONSTRAINT genres_length_check CHECK (array_length(genres, 1) BETWEEN 1 AND 5);


INSERT INTO books (title, published, pages, genres) VALUES ('Reclaim', 2023, 102, ARRAY [ 'productivity','self-help' ]) RETURNING id, created_at, version;


ALTER TABLE books DROP CONSTRAINT IF EXISTS books_published_check;