Create table beers (
   id serial primary key,
   name varchar(50) not null,
   description varchar(500) not null,
   currency varchar(50) not null,
   brewery varchar(50) not null,
   country varchar(50) not null,
   price numeric(10,2) not null
);