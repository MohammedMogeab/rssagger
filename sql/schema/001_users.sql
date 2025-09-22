-- +goose up

create table users(
    id UUID PRIMARY KEY ,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
   
);

-- +goose down
drop table if exists users;