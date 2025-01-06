create table products (
	id serial primary key,
	name varchar(50) not null,
	price numeric(10, 2) not null,
	code varchar(50) not null,
    qtd numeric(10, 2) not null,
	unity varchar(50) not null,
	created_at timestamp with time zone default now(),
	updated_at timestamp with time zone default now(),
	unique (name, code)
);