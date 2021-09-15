CREATE TABLE notes
(
  id serial not null unique,
  title varchar(255) not null,
  text varchar(255) not null
)