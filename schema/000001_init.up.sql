CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE currencies
(
    id serial not null unique,
    name varchar(31) not null,
    ticket varchar(31) not null
);

CREATE TABLE sources
(
    id serial not null unique,
    user_id int references users(id) on delete cascade not null,
    source_type varchar(255),
    balance numeric not null default 0.0,
    currency_id int references currencies(id) on delete set null

);

CREATE TABLE categories
(
    id serial not null unique,
    name varchar(255) not null,
    user_id int references users(id) on delete cascade not null
);

CREATE TYPE activity_type AS ENUM ('income', 'expense');

CREATE TABLE activities
(
    id serial not null unique,
    user_id int references users(id) on delete cascade not null,
    source_id int references sources(id) on delete set null,
    category_id int references categories(id) on delete set null,
    activity_type activity_type not null,
    change numeric not null,
    label varchar(255) not null,
    activity_date timestamptz not null
)