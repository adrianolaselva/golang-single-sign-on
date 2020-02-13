
-- +migrate Up
create table if not exists roles (
  name varchar(50) NOT NULL,
  primary key(name)
);

-- +migrate Down
drop table if exists roles;
