Create table beers (
    id serial primary key,
    name varchar(50) not null,
    description varchar(500) not null,
    price numeric(10,2) not null
);