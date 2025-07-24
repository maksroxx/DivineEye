CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS alerts (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    coin TEXT NOT NULL,
    price DOUBLE PRECISION NOT NULL,
    direction TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS notification_alerts (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    coin TEXT NOT NULL,
    price DOUBLE PRECISION NOT NULL,
    direction TEXT NOT NULL
);

ALTER TABLE alerts ADD COLUMN triggered BOOLEAN DEFAULT false;
ALTER TABLE notification_alerts ADD COLUMN triggered BOOLEAN DEFAULT false;

CREATE INDEX idx_alerts_coin_price ON notification_alerts(coin, price, direction);
CREATE INDEX idx_alerts_triggered ON alerts(triggered);