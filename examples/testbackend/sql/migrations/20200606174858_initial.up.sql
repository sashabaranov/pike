CREATE TABLE animals (
	id SERIAL PRIMARY KEY,
	name VARCHAR(128),
	age INTEGER,
	photo_url TEXT,
	created timestamp(0) without time zone default (now() at time zone 'utc')
);

