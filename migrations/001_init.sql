-- Initial SQL migration (provided as reference; Gorm AutoMigrate is used by default)

CREATE TABLE IF NOT EXISTS nodes (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    name VARCHAR(128) UNIQUE NOT NULL,
    address VARCHAR(256) NOT NULL,
    is_leader BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS services (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    name VARCHAR(128) UNIQUE NOT NULL,
    version VARCHAR(64),
    owner VARCHAR(128)
);
