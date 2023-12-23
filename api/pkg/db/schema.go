package db

var schema = `
CREATE TABLE if not exists users(
	id SERIAL PRIMARY KEY,
	firstName TEXT,
	lastName TEXT,
	email TEXT UNIQUE,
	password TEXT,
	active BOOLEAN,
	createdAt DATE,
	updatedAt DATE
);

CREATE TABLE if not exists revenu(
	id SERIAL PRIMARY KEY,
	name TEXT,
	amount DECIMAL(10, 2),
	createdAt DATE,
	updatedAt DATE
);

CREATE TABLE if not exists expense(
	id SERIAL PRIMARY KEY,
	name TEXT,
	amount DECIMAL(10, 2),
	createdAt DATE,
	updatedAt DATE
);
`
