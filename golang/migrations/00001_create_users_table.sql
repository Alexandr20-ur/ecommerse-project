-- +goose Up
create table users (
    -- 1. Идентификатор всегда UUID
    id UUID primary key default gen_random_uuid(),

    -- 2. Базовая информация
    email text not null unique,
    phone text unique,
    password_hash varchar(255) not null,

    -- 3. Профиль пользователя
    full_name text,
    avatar_url text,
    role varchar(50) not null default 'user', -- 'user', 'admin', 'moderator'

    -- 4. Cтатусы и мета-информация
    is_active boolean not null default true,
    is_email_verified boolean not null default false,
    las_login_at timestamp,

    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp
);

create index idx_users_email on users(email);
create index idx_users_role on users(role);
create index idx_users_is_active on users(is_active);
create index idx_users_created_at on users(created_at);

-- +goose Down
drop index idx_users_email;
drop index idx_users_role;
drop index idx_users_is_active;
drop index idx_users_created_at;

delete table users;
