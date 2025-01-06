create table users (
	id serial primary key,
	username varchar(100) not null,
	email varchar(100) not null,
	password varchar(100) not null,
	created_at timestamp with time zone default now(),
	updated_at timestamp with time zone default now(),
	unique (email, username)
);