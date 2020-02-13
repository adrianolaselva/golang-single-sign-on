
-- +migrate Up
create table if not exists users (
  uuid varchar(60) not null,
  username varchar(50) NOT NULL,
  email varchar(50),
  password varchar(500),
  activated boolean default false,
  created_at datetime not null,
  updated_at datetime not null,
  unique(username),
  primary key(uuid)
);

-- +migrate Down
drop table if exists users;