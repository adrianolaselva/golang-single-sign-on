
-- +migrate Up
create table if not exists  oauth_clients (
  client_id varchar(255) primary key,
  client_secret varchar(255),
  user_uuid varchar(60) not null,
  redirect varchar(255) null,
  created_at datetime not null,
  updated_at datetime not null
);

-- +migrate Down
drop table if exists oauth_clients;
