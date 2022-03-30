create table users (
	user_id serial not null primary key,
	firstname varchar(20) not null,
	lastname varchar(20) not null,
	created_at timestamptz default current_timestamp
);


insert into users(firstname, lastname) values ('Muhammad', 'Sarvar');
insert into users(firstname, lastname) values ('Izzat', 'Ahmedov');
insert into users(firstname, lastname) values ('Nima', 'Gap');


	