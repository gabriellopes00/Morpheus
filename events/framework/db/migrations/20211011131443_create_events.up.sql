DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'EVENT_STATUS') THEN
        CREATE TYPE EVENT_STATUS AS ENUM ('available', 'finished', 'canceled');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'EVENT_AGE_GROUP') THEN
        CREATE TYPE EVENT_AGE_GROUP AS ENUM ('0', '10', '12', '14', '16', '18');
    END IF;
END$$;


CREATE TABLE IF NOT EXISTS "events" (
    id UUID UNIQUE NOT NULL,
    name VARCHAR NOT NULL,
    description TEXT DEFAULT NULL,
    is_available BOOLEAN NOT NULL,
    organizer_account_id UUID NOT NULL,
    age_group EVENT_AGE_GROUP NOT NULL,
    maximum_capacity INTEGER NOT NULL CHECK (maximum_capacity >= 1),
    status EVENT_STATUS NOT NULL,
    ticket_price REAL NOT NULL CHECK (ticket_price >= 0),
    dates TIMESTAMP [] NOT NULL,
    location_street VARCHAR NOT NULL,
    location_district VARCHAR NOT NULL,
    location_state VARCHAR NOT NULL,
    location_city VARCHAR NOT NULL,
    location_postal_code VARCHAR NOT NULL,
    location_description TEXT DEFAULT NULL,
    location_number SMALLINT NOT NULL CHECK (location_number >= 0),
    location_latitude FLOAT DEFAULT NULL,
    location_longitude FLOAT DEFAULT NULL,
    created_at TIMESTAMP NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (organizer_account_id) REFERENCES accounts(id) ON DELETE CASCADE ON UPDATE CASCADE
);