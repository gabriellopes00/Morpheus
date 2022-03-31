CREATE TABLE IF NOT EXISTS "event_ticket_options" (
    id UUID UNIQUE NOT NULL,
    title VARCHAR NOT NULL,
    description TEXT DEFAULT NULL,
    event_id UUID NOT NULL,
    sales_start_datetime TIMESTAMP WITH TIME ZONE NOT NULL,
    sales_end_datetime TIMESTAMP WITH TIME ZONE NOT NULL,
    minimum_buys_quantity SMALLINT NOT NULL,
    maximum_buys_quantity SMALLINT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE
);