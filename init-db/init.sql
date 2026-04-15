CREATE schema auth;




-- 1. Users Table
-- Uses native uuidv7() as the default primary key
CREATE TABLE auth.users (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    full_name TEXT NOT NULL,
    is_email_verified BOOLEAN DEFAULT FALSE NOT NULL, 
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ DEFAULT NULL
);

INSERT INTO auth.users (email, password, full_name) VALUES
('alex.rivera@example.com', 'P@ssw0rd123', 'Alex Rivera'),
('sarah.chen@techmail.org', 'SecureKey!99', 'Sarah Chen'),
('jordan.smith@webmail.net', 'QueryMaster#1', 'Jordan Smith'),
('marta.gomez@pro-dev.io', 'DevOps_Life2026', 'Marta Gomez'),
('liam.wilson@startup.com', 'BlueSky$88', 'Liam Wilson'),
('priya.sharma@data-viz.com', 'ChartLover!22', 'Priya Sharma'),
('kevin.adams@cloud-sync.net', 'Nebula_77!#', 'Kevin Adams');


--  Refresh Tokens Table
CREATE TABLE auth.refresh_tokens (
    id UUID PRIMARY KEY DEFAULT uuidv7(), 
    user_id UUID NOT NULL REFERENCES auth.users(id) ON DELETE CASCADE,
    token_hash TEXT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    expires_at TIMESTAMPTZ NOT NULL,
    revoked_at TIMESTAMPTZ 
);



-- PostgREST setup
create role web_anon nologin;

grant usage ON schema auth to web_anon;
GRANT SELECT ON ALL TABLES IN SCHEMA auth TO web_anon;

create role authenticator noinherit login password 'mysecretpassword';
grant web_anon to authenticator;

-- End PostgREST setup

/* -- 3. Roles Table
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    role_name TEXT UNIQUE NOT NULL
);

-- 4. User Roles (Many-to-Many)
CREATE TABLE user_roles (
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    role_id INTEGER REFERENCES roles(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, role_id)
); */


