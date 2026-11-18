-- +goose Up
CREATE TABLE IF NOT EXISTS products (
                                        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                        name VARCHAR(255) NOT NULL,
                                        description TEXT,
                                        price NUMERIC(10, 2) NOT NULL CHECK (price >= 0),
                                        stock INTEGER NOT NULL DEFAULT 0 CHECK (stock >= 0),
                                        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                                        updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Индекс для быстрого поиска по названию
CREATE INDEX idx_products_name ON products USING gin(to_tsvector('russian', name));

-- +goose Down
DROP INDEX IF EXISTS idx_products_name;
DROP TABLE IF EXISTS products;