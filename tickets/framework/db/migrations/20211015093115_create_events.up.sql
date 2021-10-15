CREATE TABLE IF NOT EXISTS "events" (
    id UUID UNIQUE NOT NULL,
    date TIMESTAMP WITH TIME ZONE NOT NULL,

    PRIMARY KEY (id)
);