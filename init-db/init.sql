create schema api;

create table api.users (
  id UUID PRIMARY KEY DEFAULT uuidv7()
  email text
  full_name text not null
  /* A typical Argon2id string looks like this:
    $argon2id$v=19$m=65536,t=3,p=4$qH1N...$zX9a... */
  password text not null
);

insert into api.users (email, full_name, password) values
  ('finish tutorial 0'), ('pat self on back');

create role web_anon nologin;

grant usage ON schema api to web_anon;
GRANT SELECT ON ALL TABLES IN SCHEMA api TO web_anon;

create role authenticator noinherit login password 'mysecretpassword';
grant web_anon to authenticator;


