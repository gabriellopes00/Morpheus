CREATE TABLE IF NOT EXISTS "event_locations" (
    id UUID UNIQUE NOT NULL,
    event_id UUID NOT NULL,
    street VARCHAR NOT NULL,
    district VARCHAR NOT NULL,
    state CHAR(2) NOT NULL,
    city VARCHAR NOT NULL,
    postal_code CHAR(8) NOT NULL,
    description TEXT DEFAULT NULL,
    number SMALLINT NOT NULL CHECK (number >= 0),
    latitude FLOAT DEFAULT NULL,
    longitude FLOAT DEFAULT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE
);