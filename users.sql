CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255),
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    password VARCHAR(60),
    user_active BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (
    email,
    first_name,
    last_name,
    password,
    user_active,
    created_at,
    updated_at
) VALUES (
    'admin@example.com',
    'Admin',
    'User',
    '$2a$12$1zGLuYDDNvATh4RA4avbKuheAMpb1svexSzrQm7up.bnpwQHs0jNe', 
    true,
    '2026-03-14 00:00:00',
    '2026-03-14 00:00:00'
);