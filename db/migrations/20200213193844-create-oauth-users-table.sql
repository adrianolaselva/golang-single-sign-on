
-- +migrate Up
create table if not exists oauth_users (
  id varchar(36) not null,
  name varchar(120) NOT NULL,
  last_name varchar(120) NOT NULL,
  email varchar(120) NOT NULL,
  username varchar(120) NOT NULL,
  password varchar(500) NOT NULL,
  birthday date null,
  activated boolean default false,
  created_at datetime not null,
  updated_at datetime not null,
  expires_at datetime null,
  deleted_at datetime null,
  unique(username),
  unique(email),
  primary key(id)
);

-- +migrate Down
drop table if exists oauth_users;