
-- +migrate Up
create table if not exists  oauth_clients (
  id varchar(255) not null,
  user_id varchar(60) not null,
  name varchar(120) not null,
  secret varchar(120) not null,
  scopes varchar(1024) null,
  redirect varchar(255) null,
  revoked boolean not null default false,
  created_at datetime not null,
  updated_at datetime not null,
  deleted_at datetime null,
  foreign key(user_id) references oauth_users(id),
  primary key(id)
);

-- +migrate Down
drop table if exists oauth_clients;
