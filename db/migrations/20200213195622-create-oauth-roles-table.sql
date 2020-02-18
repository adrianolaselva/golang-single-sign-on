
-- +migrate Up
create table if not exists oauth_roles (
  id varchar(36) not null,
  name varchar(50) not null,
  unique(name),
  primary key(id)
);

-- +migrate Down
drop table if exists oauth_roles;
