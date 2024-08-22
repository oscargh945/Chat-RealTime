CREATE TABLE IF NOT EXISTS users (
    id uuid primary key default uuid_generate_v4(),
    user_name varchar(255) not null,
    email varchar(200) unique,
    password varchar not null,
    create_at timestamptz not null default now(),

)