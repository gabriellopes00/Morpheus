CREATE TABLE IF NOT EXISTS "event_tycket_options" (
    id UUID UNIQUE NOT NULL,
    event_id UUID NOT NULL,
    title VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE
);