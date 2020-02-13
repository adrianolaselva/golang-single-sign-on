
-- +migrate Up
create table if not exists  oauth_personal_access_clients (
  uuid varchar(255) primary key,
  client_id varchar(255),
  created_at datetime not null,
  updated_at datetime not null
);

-- +migrate Down
drop table if exists oauth_personal_access_clients;
