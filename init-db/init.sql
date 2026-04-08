create schema api;

-- 1. Users Table
-- Uses native uuidv7() as the default primary key
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    full_name TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);


--  Refresh Tokens Table
CREATE TABLE refresh_tokens (
    -- You can also use uuidv7() here for the token ID itself
    id UUID PRIMARY KEY DEFAULT uuidv7(), 
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token_hash TEXT NOT NULL UNIQUE,
    expires_at TIMESTAMPTZ NOT NULL,
    revoked_at TIMESTAMPTZ, 
    created_at TIMESTAMPTZ DEFAULT NOW()
);



-- postgREST setup
create role web_anon nologin;

grant usage ON schema api to web_anon;
GRANT SELECT ON ALL TABLES IN SCHEMA api TO web_anon;

create role authenticator noinherit login password 'mysecretpassword';
grant web_anon to authenticator;
-- End postgREST setup

-- 3. Roles Table
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    role_name TEXT UNIQUE NOT NULL
);

-- 4. User Roles (Many-to-Many)
CREATE TABLE user_roles (
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    role_id INTEGER REFERENCES roles(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, role_id)
);

insert into api.users (email, full_name, password) values
  ('finish tutorial 0'), ('pat self on back');
